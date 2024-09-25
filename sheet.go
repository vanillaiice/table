package table

import "github.com/tealeg/xlsx/v3"

// Sheet represents a sheet in a xlsx file.
type Sheet struct {
	name string // name of the sheet.
	rows []*Row // rows in the sheet.
}

// NewRow returns a new row from a sheet.
func (s *Sheet) NewRow() *Row {
	row := &Row{cells: []*Cell{}}
	s.rows = append(s.rows, row)
	return row
}

// AddRows adds rows to the sheet.
func (s *Sheet) AddRows(rows ...*Row) {
	s.rows = append(s.rows, rows...)
}

// render renders the sheet.
func (s *Sheet) render() (*xlsx.Sheet, error) {
	sheet, err := xlsx.NewSheet(s.name)
	if err != nil {
		return nil, err
	}

	for _, row := range s.rows {
		r := sheet.AddRow()
		if row.height != 0 {
			r.SetHeight(row.height)
		}

		row.render(r)
	}

	return sheet, nil
}
