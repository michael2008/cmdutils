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

type ColorFunc func(row []string) []tablewriter.Colors

func PrintTableRich(header []string, data [][]string, mergeCols []int, colorFunc ColorFunc) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetAutoWrapText(true)
	table.SetReflowDuringAutoWrap(false)
	table.SetColWidth(120)
	table.SetBorder(true)
	table.SetRowLine(true)
	if mergeCols != nil {
		table.SetAutoMergeCellsByColumnIndex(mergeCols)
	}
	if len(header) > 0 {
		table.SetHeader(header)
	}
	if colorFunc == nil {
		table.AppendBulk(data)
	} else {
		for _, row := range data {
			table.Rich(row, colorFunc(row))
		}
	}
	table.Render()
}
