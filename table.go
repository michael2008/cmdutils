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
		table.SetHeaderColor(newColorHeaderBold(len(header))...)
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
		table.SetHeaderColor(newColorHeaderBold(len(header))...)
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

// ColorMapping maps headerValue to func generating color by cellValue.
type ColorMapping map[string]func(cellValue string) tablewriter.Colors

func PrintTableRichMapping(header []string, data [][]string, mergeCols []int, colorMapping ColorMapping) {
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
		table.SetHeaderColor(newColorHeaderBold(len(header))...)
	}
	if colorMapping == nil {
		table.AppendBulk(data)
	} else {
		var funcList = make([]func(cellValue string) tablewriter.Colors, len(header))
		for idx, h := range header {
			if f, ok := colorMapping[h]; ok {
				funcList[idx] = f
			}
		}
		for _, row := range data {
			if len(row) > len(header) {
				row = row[:len(header)]
			}
			colorRow := newColorRow(len(header))
			for idx, cellValue := range row {
				if f := funcList[idx]; f != nil {
					color := f(cellValue)
					if len(color) == 0 {
						color = tablewriter.Colors{}
					}
					colorRow[idx] = color
				}
			}
			table.Rich(row, colorRow)
		}
	}
	table.Render()
}

func newColorRow(rowSize int) []tablewriter.Colors {
	list := make([]tablewriter.Colors, rowSize)
	for i := range list {
		list[i] = tablewriter.Colors{}
	}
	return list
}

func newColorHeaderBold(headerSize int) []tablewriter.Colors {
	list := make([]tablewriter.Colors, headerSize)
	for i := range list {
		list[i] = tablewriter.Colors{tablewriter.Bold}
	}
	return list
}
