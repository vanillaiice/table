package table

import (
	"github.com/tealeg/xlsx/v3"
)

// Cell represents a cell in a row.
type Cell struct {
	Content        any        // content of the cell.
	MergeH, MergeV int        // number of cells to merge horizontally and vertically, respectively.
	Style          *CellStyle // style of the cell.
}

// NewCellOpts are the options for creating a new cell.
type NewCellOpts struct {
	Style          *CellStyle // style of the cell.
	MergeH, MergeV int        // number of cells to merge horizontally and vertically, respectively.
}

// NewCellStyle returns the default cell style.
func NewCellStyle() *CellStyle {
	return (*CellStyle)(xlsx.NewStyle())
}

// NewCell returns a new cell.
func NewCell(Content any, style ...*CellStyle) *Cell {
	var s *CellStyle
	if len(style) > 0 {
		s = style[0]
	} else {
		s = NewCellStyle()
	}

	return &Cell{Content: Content, Style: s}
}

// NewCells createsand returns new cells.
func NewCells(opt *NewCellOpts, content ...any) []*Cell {
	var cells []*Cell
	for _, v := range content {
		s := *opt.Style
		cells = append(cells, &Cell{Content: v, MergeH: opt.MergeH, MergeV: opt.MergeV, Style: &s})
	}
	return cells
}

// render renders the cell.
func (c *Cell) render(cell *xlsx.Cell) {
	if c.MergeH > 0 {
		c.MergeH = c.MergeH - 1
	}
	if c.MergeV > 0 {
		c.MergeV = c.MergeV - 1
	}

	cell.Merge(c.MergeH, c.MergeV)

	cell.SetStyle((*xlsx.Style)(c.Style))

	cell.SetValue(c.Content)
}
