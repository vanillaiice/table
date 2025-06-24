package main

import (
	"github.com/vanillaiice/table"
)

func main() {
	tbl := table.NewTable()
	sheet, err := tbl.AddSheet("cars")
	if err != nil {
		panic(err)
	}

	const (
		redHex    = "ffff0000"
		blueHex   = "ff0000ff"
		orangeHex = "ffff8c00"
	)

	title := table.RowTitle("Some Cars", 4, "medium")
	sheet.AddRows(title, table.NewRow())

	rows := []*table.Row{
		table.NewRow().SetHeight(30).AddCells(
			table.NewCells(
				table.NewCellOpts(table.NewCellStyle().SetBold(true).SetItalic(true).SetAlignementH("center")),
				"Make", "Model", "Color", "Year",
			)...,
		),
		table.NewRow().SetHeight(30).AddCells(
			table.NewCells(
				table.NewCellOpts(table.NewCellStyle().SetWrapText(true).SetFontColor(redHex)),
				"Ford", "Mustang", "Red", 2014,
			)...,
		),
		table.NewRow().SetHeight(30).AddCells(
			table.NewCells(
				table.NewCellOpts(table.NewCellStyle().SetWrapText(true).SetFontColor(blueHex)),
				"Subaru", "WRX STI", "Blue", 1998,
			)...,
		),
		table.NewRow().SetHeight(30).AddCells(
			table.NewCells(
				table.NewCellOpts(table.NewCellStyle().SetWrapText(true).SetFontColor(orangeHex)),
				"Honda", "Civic EK9", "Orange", 1999,
			)...,
		),
	}
	table.AddOuterAndInnerBorderRow("thin", rows...)
	sheet.AddRows(rows...)

	if err := tbl.Render(); err != nil {
		panic(err)
	}

	if err := tbl.Save("cars.xlsx"); err != nil {
		panic(err)
	}
}
