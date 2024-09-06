package table

import (
	"github.com/tealeg/xlsx/v3"
)

// Row represents a row in a table.
type Row struct {
	cells  []*Cell // cells in the row.
	height float64 // height of the row.
}

// SetHeight sets the height of the row.
func (r *Row) SetHeight(h float64) *Row {
	r.height = h
	return r
}

// NewRow returns a new row.
func NewRow() *Row {
	return &Row{cells: []*Cell{}}
}

// NewCell returns a new cell from a row.
func (r *Row) NewCell() *Cell {
	cell := &Cell{}
	r.cells = append(r.cells, cell)
	return cell
}

// AddCells adds cells to a row.
func (r *Row) AddCells(cells ...*Cell) *Row {
	r.cells = append(r.cells, cells...)
	return r
}

// IterCells iterates over the cells in the row and calls the function f.
func (r *Row) IterCells(f func(int, *Cell) error) {
	for i, c := range r.cells {
		f(i, c)
	}
}

// MergeRows merges multiple rows.
func MergeRows(rows ...*Row) *Row {
	row := NewRow()

	for _, r := range rows {
		row.AddCells(r.cells...)
	}

	return row
}

// render renders the row.
func (r *Row) render(row *xlsx.Row) {
	for _, cell := range r.cells {
		cell.render(row.AddCell())
	}
}
