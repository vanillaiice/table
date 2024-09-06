# table [![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/github.com/vanillaiice/table) [![Go Report Card](https://goreportcard.com/badge/github.com/vanillaiice/table)](https://goreportcard.com/report/github.com/vanillaiice/table)

Package table provides functions to create and render xlsx tables.
It is built atop of the [tealeg/xlsx](https://github.com/tealeg/xlsx) package.

# Example

```go
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

	title := table.RowTitle("Some Cars", 4, "medium")
	sheet.AddRows(title, table.NewRow())

	const (
		redHex    = "ffff0000"
		blueHex   = "ff0000ff"
		orangeHex = "ffff8c00"
	)

	rows := []*table.Row{
		table.NewRow().SetHeight(30).AddCells(
			table.NewCells(
				&table.NewCellOpts{Style: table.NewCellStyle().SetBold(true).SetItalic(true).SetAlignementH("center")},
				"Make", "Model", "Color", "Year",
			)...,
		),
		table.NewRow().AddCells(
			table.NewCells(
				&table.NewCellOpts{Style: table.NewCellStyle().SetWrapText(true).SetFontColor(redHex)},
				"Ford", "Mustang", "Red", 2014,
			)...,
		),
		table.NewRow().AddCells(
			table.NewCells(
				&table.NewCellOpts{Style: table.NewCellStyle().SetWrapText(true).SetFontColor(blueHex)},
				"Subaru", "WRX STI", "Blue", 1998,
			)...,
		),
		table.NewRow().AddCells(
			table.NewCells(
				&table.NewCellOpts{Style: table.NewCellStyle().SetWrapText(true).SetFontColor(orangeHex)},
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
```

# License

GPLv3

# Author

[vanillaiice](https://github.com/vanillaiice)

# Credits

[tealeg/xlsx](https://github.com/tealeg/xlsx)
