package table

import "slices"

// RowGroup is a struct holding a group of rows.
// It is used to simplify the process of adding rows in
// a sheet. And also to easily add borders around a group
// of rows.
type RowGroup struct {
	rows  [][]*Row
	index int
}

// NewRowGroup initializes and returns a pointer to a RowGroup.
func NewRowGroup() *RowGroup {
	rows := [][]*Row{{}}
	return &RowGroup{rows: rows}
}

// NewGroup creates a new group in the RowGroup struct.
func (rg *RowGroup) NewGroup() *RowGroup {
	rg.index++
	rg.rows = append(rg.rows, []*Row{})
	return rg
}

// NewRow creates a new row, adds it to the current row group,
// and returns it.
func (rg *RowGroup) NewRow() *Row {
	row := NewRow()
	rg.rows[rg.index] = append(rg.rows[rg.index], row)
	return row
}

// AddRow adds a single row to the current group of rows.
func (rg *RowGroup) AddRow(row *Row) *Row {
	rg.rows[rg.index] = append(rg.rows[rg.index], row)
	return row
}

// AddRow adds rows to the current group of rows.
func (rg *RowGroup) AddRows(rows ...*Row) []*Row {
	rg.rows[rg.index] = append(rg.rows[rg.index], rows...)
	return rg.rows[rg.index]
}

// AddBorder adds borders around a group of rows.
func (rg *RowGroup) AddBorder(borderStyle string, borderFunc func(borderStyle string, rows ...*Row)) {
	borderFunc(borderStyle, rg.rows[rg.index]...)
}

// Rows returns the rows in the current group of rows.
func (rg *RowGroup) Rows() []*Row {
	return rg.rows[rg.index]
}

// All returns all the rows in different groups as a single slice.
func (rg *RowGroup) All() []*Row {
	return slices.Concat(rg.rows...)
}
