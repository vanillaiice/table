package table

import (
	"github.com/tealeg/xlsx/v3"
)

// Cell represents a cell in a row.
type Cell struct {
	Content any       // content of the cell.
	Opts    *CellOpts // cell options.
}

// CellOpts are the cell's options.
type CellOpts struct {
	Style  *CellStyle // excel style of the cell.
	MergeH int        // number of cells to merge horizontally.
	MergeV int        // number of cells to merge vertically.
	Format string     // excel format of the cell.
}

// NewCellOpts returns a pointer to a new CellOpts.
func NewCellOpts(style ...*CellStyle) *CellOpts {
	var _style *CellStyle
	if len(style) == 0 {
		_style = NewCellStyle()
	} else {
		_style = style[0]
	}
	return &CellOpts{Style: _style}
}

// GetStyle returns the CellOpts's underlying CellStyle pointer.
func (co *CellOpts) GetStyle() *CellStyle {
	return co.Style
}

// SetStyle sets the CellOpts's style.
func (co *CellOpts) SetStyle(style *CellStyle) *CellOpts {
	co.Style = style
	return co
}

// SetFormat sets the excel format of the cell.
func (co *CellOpts) SetFormat(format string) *CellOpts {
	co.Format = format
	return co
}

// SetMergeH sets the horizontal merge value of the cell.
func (co *CellOpts) SetMergeH(hcells int) *CellOpts {
	co.MergeH = hcells
	return co
}

// SetMergeV sets the vertical merge value of the cell.
func (co *CellOpts) SetMergeV(vcells int) *CellOpts {
	co.MergeV = vcells
	return co
}

// NewCell returns a new cell.
func NewCell(content any, opts ...*CellOpts) *Cell {
	var _opts *CellOpts
	if len(opts) == 0 {
		_opts = NewCellOpts()
	} else {
		_opts = opts[0]
	}
	return &Cell{Content: content, Opts: _opts}
}

// NewCells creates and returns new cells.
func NewCells(opts *CellOpts, content ...any) []*Cell {
	var cells []*Cell
	for _, v := range content {
		o := *opts
		cells = append(cells, &Cell{Content: v, Opts: &o})
	}
	return cells
}

// render renders the cell.
func (c *Cell) render(cell *xlsx.Cell) {
	if c.Opts.MergeH > 0 {
		c.Opts.MergeH -= 1
	}
	if c.Opts.MergeV > 0 {
		c.Opts.MergeV -= 1
	}
	if c.Opts.Format != "" {
		cell.SetFormat(c.Opts.Format)
	}

	cell.Merge(c.Opts.MergeH, c.Opts.MergeV)
	cell.SetStyle((*xlsx.Style)(c.Opts.Style))
	cell.SetValue(c.Content)
}
