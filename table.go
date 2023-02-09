package cmdutils

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetAutoWrapText(true)
	table.SetColWidth(120)
	table.SetRowLine(true)
	if len(header) > 0 {
		table.SetHeader(header)
	}
	table.AppendBulk(data)
	table.Render()
}
