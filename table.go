/*
Package table provides functions to create and render xlsx tables.
It is built atop of the tealeg/xlsx (https://github.com/tealeg/xlsx) package.
*/
package table

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// Table represents a xlsx table.
type Table struct {
	file      *xlsx.File        // file is the xlsx file that contains the table.
	sheets    []*Sheet          // sheets is the list of sheets in the table.
	sheetsMap map[string]*Sheet // sheetsMap is the map of sheets in the table.
}

// TableRenderOpts are the options for rendering a table.
type TableRenderOpts struct {
	SetAutoColWidth bool // SetAutoColWidth sets the automatic column width for the specified columns during rendering.
}

// NewTable creates a new table.
func NewTable() *Table {
	return &Table{file: xlsx.NewFile(), sheets: []*Sheet{}, sheetsMap: map[string]*Sheet{}}
}

// Save writes the table to the specified path.
func (t *Table) Save(path string) error {
	return t.file.Save(path)
}

// AddSheet adds a sheet to the table.
func (t *Table) AddSheet(name string) (*Sheet, error) {
	if _, ok := t.sheetsMap[name]; ok {

		return nil, fmt.Errorf("sheet %s already exists", name)
	}

	sheet := &Sheet{name: name, rows: []*Row{}}
	t.sheetsMap[name] = sheet
	t.sheets = append(t.sheets, sheet)

	return sheet, nil
}

// GetSheetFromName returns a sheet with the specified name from the table.
func (t *Table) GetSheetFromName(name string) (*Sheet, error) {
	sh, ok := t.sheetsMap[name]
	if !ok {
		return nil, fmt.Errorf("sheet %s not found", name)
	}

	return sh, nil
}

// GetSheetFromIndex returns a sheet with the specified index from the table.
func (t *Table) GetSheetFromIndex(index int) (*Sheet, error) {
	if index > len(t.sheets) || index < 0 {
		return nil, fmt.Errorf("sheet #%d out of range", index)
	}

	return t.sheets[index], nil
}

// Render renders the table.
func (t *Table) Render(tableRenderOpts ...*TableRenderOpts) error {
	var opt *TableRenderOpts
	if len(tableRenderOpts) > 1 {
		opt = tableRenderOpts[0]
	} else {
		opt = &TableRenderOpts{SetAutoColWidth: false}
	}

	for _, sh := range t.sheets {
		sheet, err := sh.render()
		if err != nil {
			return err
		}

		if opt.SetAutoColWidth {
			for i := 1; i <= sheet.MaxCol; i++ {
				if err := sheet.SetColAutoWidth(i, func(s string) float64 {
					return float64(len(s))
				}); err != nil {
					return err
				}
			}
		}

		if _, err := t.file.AppendSheet(*sheet, sh.name); err != nil {
			return err
		}
	}

	return nil
}
