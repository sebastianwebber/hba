package ui

import (
	"bufio"
	"fmt"
	"io"

	"github.com/sebastianwebber/hba"

	"github.com/olekukonko/tablewriter"
)

var writer *bufio.Writer

// DisplayRules display rules as a tabled
func DisplayRules(in []hba.Rule, out io.Writer) {
	writer = bufio.NewWriter(out)
	defer writer.Flush()

	if len(in) == 0 {
		writer.WriteString("No rules found\n")
		return
	}
	renderTable(in, writer)
}

func renderTable(in []hba.Rule, out io.Writer) {

	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Line", "Type", "Database", "User", "addresses", "method", "comments"})

	for _, v := range in {
		table.Append(
			[]string{
				fmt.Sprintf("%d", v.LineNumber),
				v.Type,
				v.DatabaseName,
				v.UserName,
				prettyAddress(v),
				v.AuthMethod,
				v.Comments,
			})
	}

	table.SetBorder(false)
	table.SetCaption(true, fmt.Sprintf("(%d rows)", len(in)))

	table.Render()
}

func prettyAddress(r hba.Rule) string {
	if r.DNSAddress != "" {
		return r.DNSAddress
	}

	if r.NetworkMask == nil {
		return ""
	}

	octMask, _ := r.NetworkMask.Size()
	return fmt.Sprintf("%s/%d", r.IPAddress.String(), octMask)
}
