package table

// GenericTable creates a generic table, with a title, headings and rows.
func GenericTable(title string, headings []*Cell, rows []*Row, borderStyle string) []*Row {
	tblRows := []*Row{}

	if title != "" {
		tblRows = append(tblRows, RowTitle(title, len(headings), borderStyle))
	}

	if len(headings) > 0 {
		headingsRow := RowWithMergedCells(headings...)
		for _, hr := range headingsRow {
			for _, c := range hr.cells {
				c.Style.SetBold(true).
					SetAlignementH("center").
					SetAlignementV("center")
			}
		}
		AddOuterBorderRows(borderStyle, headingsRow...)
		tblRows = append(tblRows, headingsRow...)
	}

	return append(tblRows, rows...)
}

// RowTitle creates a title row.
func RowTitle(title string, cols int, borderStyle string, cellStyle ...*CellStyle) *Row {
	var style *CellStyle

	if len(cellStyle) > 0 {
		style = cellStyle[0]
	} else {
		style = NewCellStyle()
		style.
			SetFontSize(style.Font.Size + 2).
			SetBold(true).
			SetAlignementH("center").
			SetAlignementV("center").
			SetWrapText(true)
	}

	row := RowWithMergedCellsH(&Cell{Content: title, MergeH: cols, Style: style})

	AddOuterBorderRows(borderStyle, row)

	return row
}

// RowWithMergedCells creates a row with merged cells.
func RowWithMergedCells(cells ...*Cell) []*Row {
	row := RowWithMergedCellsH(cells...)
	return append([]*Row{}, RowWithMergedCellsV(row.cells...)...)
}

// RowWithMergedCellsH creates a row with horizontally merged cells.
func RowWithMergedCellsH(cells ...*Cell) *Row {
	row := NewRow()

	for _, c := range cells {
		row.AddCells(c)
		for i := 1; i < c.MergeH; i++ {
			row.AddCells(&Cell{Style: NewCellStyle()})
		}
	}

	return row
}

// RowWithMergedCellsV creates a row with vertically merged cells.
func RowWithMergedCellsV(cells ...*Cell) []*Row {
	rows := []*Row{{cells: cells}}

	var nMaxRows int
	for _, c := range cells {
		var nRows int
		for i := 1; i < c.MergeV; i++ {
			if nMaxRows != 0 && nRows < nMaxRows {
				nRows++
			} else {
				nRows++
				nMaxRows++
				row := NewRow()
				for range cells {
					row.AddCells(&Cell{Style: NewCellStyle()})
				}
				rows = append(rows, row)
			}
		}
	}

	return rows
}
