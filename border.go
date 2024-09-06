package table

// AddOuterBorderRows adds outer borders around rows.
func AddOuterBorderRows(borderStyle string, rows ...*Row) {
	if len(rows) == 1 {
		for i, c := range rows[0].cells {
			if i == 0 {
				c.Style.Border.Left = borderStyle
			} else if i == len(rows[0].cells)-1 {
				c.Style.Border.Right = borderStyle
			}

			c.Style.Border.Top = borderStyle
			c.Style.Border.Bottom = borderStyle
		}
	} else {
		for i, row := range rows {
			if i == 0 {
				for j, cell := range row.cells {
					if j == 0 {
						cell.Style.Border.Left = borderStyle
					} else if j == len(row.cells)-1 {
						cell.Style.Border.Right = borderStyle
					}

					cell.Style.Border.Top = borderStyle
				}
			} else if i == len(rows)-1 {
				for j, cell := range row.cells {
					if j == 0 {
						cell.Style.Border.Left = borderStyle
					} else if j == len(row.cells)-1 {
						cell.Style.Border.Right = borderStyle
					}

					cell.Style.Border.Bottom = borderStyle
				}
			} else {
				for j, cell := range row.cells {
					if j == 0 {
						cell.Style.Border.Left = borderStyle
					} else if j == len(row.cells)-1 {
						cell.Style.Border.Right = borderStyle
					}
				}
			}
		}
	}
}

// AddInnerBorderRows adds inner borders for rows.
func AddInnerBorderRows(borderStyle string, rows ...*Row) {
	if len(rows) == 1 {
		for i, c := range rows[0].cells {
			if i == 0 || i == len(rows[0].cells)-1 {
				continue
			}

			c.Style.Border.Right = borderStyle
			c.Style.Border.Left = borderStyle
		}
	} else {
		for i, row := range rows {
			for j, c := range row.cells {
				if i == len(rows)-1 {
					if j != 0 && j != len(row.cells)-1 {
						c.Style.Border.Left = borderStyle
						c.Style.Border.Right = borderStyle
					}

					c.Style.Border.Top = borderStyle
				} else {
					if j != 0 && j != len(row.cells)-1 {
						c.Style.Border.Left = borderStyle
						c.Style.Border.Right = borderStyle
					}

					c.Style.Border.Bottom = borderStyle
				}
			}
		}
	}
}

// AddOuterAndInnerBorderRow adds outer and inner borders for rows.
func AddOuterAndInnerBorderRow(borderStyle string, rows ...*Row) {
	for _, r := range rows {
		for _, c := range r.cells {
			c.Style.SetBorders(borderStyle)
		}
	}
}
