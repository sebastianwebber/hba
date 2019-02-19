package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	fmt.Println(`Type "help" for help.`)
	for {
		t := prompt.Input(fmt.Sprintf("%s#= ", hbaFile), completer,
			prompt.OptionTitle("sql-prompt"),
			prompt.OptionHistory(histCommands),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSuggestionBGColor(prompt.DarkGray))

		switch t {
		case "show":
			showRules()
		case "help":
			showHelp()
		case "quit":
			os.Exit(0)
		}

		histCommands = append(histCommands, t)
	}
}

func showRules() {

	file, err := os.Open(hbaFile)
	defer file.Close()

	if err != nil {
		log.Fatalf("could not open file: %v", file)
	}

	all, err := hba.ParseReader(file)

	if err != nil {
		log.Fatalf("could not parse file: %v", file)
	}

	ui.DisplayRules(*all, os.Stdout)

}

func showHelp() {
	fmt.Println("You are using hba, an experimental command-line interface to pg_hba.conf file.")
	fmt.Println("type:")
	for i := 0; i < len(sugestionList); i++ {
		fmt.Printf("\t%s %s\n", sugestionList[i].Text, sugestionList[i].Description)
	}
}
