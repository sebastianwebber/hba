package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sebastianwebber/hba/cmd/ui"

	prompt "github.com/c-bata/go-prompt"
	"github.com/sebastianwebber/hba"
)

var (
	version = "devel"
	hbaFile string

	sugestionList = []prompt.Suggest{
		{Text: "show", Description: "to display all rules from pg_hba.conf file"},
		{Text: "quit", Description: "to quit"},
		{Text: "help", Description: "for help with hba commands"},
	}

	histCommands []string
	fileRules    *[]hba.Rule
)

func init() {
	flag.StringVar(&hbaFile, "hba-file", "", "full path of the pg_hba.conf file")
	flag.Parse()

	if hbaFile == "" {
		fmt.Println("Missing pg_hba.conf config file")
		os.Exit(1)
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(sugestionList, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Printf("hba version %s\n", version)
	loadFile()
	fmt.Println(`Type "help" for help.`)
	for {
		t := prompt.Input(fmt.Sprintf("%s#= ", hbaFile), completer,
			prompt.OptionTitle("sql-prompt"),
			prompt.OptionHistory(histCommands),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSuggestionBGColor(prompt.DarkGray))

		routeCmd(t)
	}
}

func routeCmd(cmd string) {
	defer func() {
		histCommands = append(histCommands, cmd)
	}()

	if strings.HasPrefix(cmd, "show") {
		parts := strings.Fields(cmd)

		if len(parts) == 1 {
			showRules(*fileRules)
			return
		}

		showRules(
			ui.Filter(
				*fileRules,
				parts[1]))
		return
	}

	if cmd == "help" {
		showHelp()
		return
	}

	os.Exit(0)
}

func loadFile() {
	var err error

	file, err := os.Open(hbaFile)
	defer file.Close()

	if err != nil {
		log.Fatalf("could not open file: %v", file)
	}

	fileRules, err = hba.ParseReader(file)

	if err != nil {
		log.Fatalf("could not parse file: %v", file)
	}

	fmt.Printf("Loaded rules from '%s' file.\n", hbaFile)

}

func showRules(rules []hba.Rule) {
	ui.DisplayRules(rules, os.Stdout)
}

func showHelp() {
	fmt.Println("You are using hba, an experimental command-line interface to pg_hba.conf file.")
	fmt.Println("type:")
	for i := 0; i < len(sugestionList); i++ {
		fmt.Printf("\t%s %s\n", sugestionList[i].Text, sugestionList[i].Description)
	}
}
