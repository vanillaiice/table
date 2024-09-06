package table

import "github.com/tealeg/xlsx/v3"

// CellStyle is the style of the cell.
type CellStyle xlsx.Style

// BorderPosition is the position of the border.
type BorderPosition int

// enum for BorderPosition.
const (
	BorderPositionTop BorderPosition = iota
	BorderPositionBottom
	BorderPositionRight
	BorderPositionLeft
)

// SetBorders sets all border.
func (c *CellStyle) SetBorders(border string) *CellStyle {
	c.Border.Top = border
	c.Border.Bottom = border
	c.Border.Right = border
	c.Border.Left = border
	return c
}

// SetBordersColor sets the all border colors.
func (c *CellStyle) SetBordersColor(borderColor string) *CellStyle {
	c.Border.TopColor = borderColor
	c.Border.BottomColor = borderColor
	c.Border.RightColor = borderColor
	c.Border.LeftColor = borderColor
	return c
}

// SetBorder sets a single border.
func (c *CellStyle) SetBorder(borderPosition BorderPosition, border string) *CellStyle {
	switch borderPosition {
	case BorderPositionTop:
		c.Border.Top = border
	case BorderPositionBottom:
		c.Border.Bottom = border
	case BorderPositionRight:
		c.Border.Right = border
	case BorderPositionLeft:
		c.Border.Left = border
	}

	return c
}

// SetBorderColor sets the border color for a single border.
func (c *CellStyle) SetBorderColor(borderPosition BorderPosition, borderColor string) *CellStyle {
	switch borderPosition {
	case BorderPositionTop:
		c.Border.TopColor = borderColor
	case BorderPositionBottom:
		c.Border.BottomColor = borderColor
	case BorderPositionRight:
		c.Border.RightColor = borderColor
	case BorderPositionLeft:
		c.Border.LeftColor = borderColor
	}

	return c
}

// SetFillPatternType sets the fill pattern type.
func (c *CellStyle) SetFillPatternType(fillPatternType string) *CellStyle {
	c.Fill.PatternType = fillPatternType
	return c
}

// SetFillBgColor sets the fill background color.
func (c *CellStyle) SetFillBgColor(fillBgColor string) *CellStyle {
	c.Fill.BgColor = fillBgColor
	return c
}

// SetFillFgColor sets the fill foreground color.
func (c *CellStyle) SetFillFgColor(fillFgColor string) *CellStyle {
	c.Fill.FgColor = fillFgColor
	return c
}

// SetFontSize sets the font size.
func (c *CellStyle) SetFontSize(fontSize float64) *CellStyle {
	c.Font.Size = fontSize
	return c
}

// SetFontName sets the font name.
func (c *CellStyle) SetFontName(fontName string) *CellStyle {
	c.Font.Name = fontName
	return c
}

// SetFontFamily sets the font family.
func (c *CellStyle) SetFontFamily(fontFamily int) *CellStyle {
	c.Font.Family = fontFamily
	return c
}

// SetFontCharset sets the font charset.
func (c *CellStyle) SetFontCharset(fontCharset int) *CellStyle {
	c.Font.Charset = fontCharset
	return c
}

// SetFontColor sets the font color.
func (c *CellStyle) SetFontColor(fontColor string) *CellStyle {
	c.Font.Color = fontColor
	return c
}

// SetBold sets if the style is bold.
func (c *CellStyle) SetBold(bold bool) *CellStyle {
	c.Font.Bold = bold
	return c
}

// SetItalic sets if the style is italic.
func (c *CellStyle) SetItalic(italic bool) *CellStyle {
	c.Font.Italic = italic
	return c
}

// SetUnderline sets if the style is underlined.
func (c *CellStyle) SetUnderline(underline bool) *CellStyle {
	c.Font.Underline = underline
	return c
}

// SetStrike sets if the style is striked.
func (c *CellStyle) SetStrike(strike bool) *CellStyle {
	c.Font.Strike = strike
	return c
}

// SetApplyBorder sets if the style is applied to the border.
func (c *CellStyle) SetApplyBorder(applyBorder bool) *CellStyle {
	c.ApplyBorder = applyBorder
	return c
}

// SetApplyFill sets if the style is applied to the fill.
func (c *CellStyle) SetApplyFill(applyFill bool) *CellStyle {
	c.ApplyFill = applyFill
	return c
}

// SetApplyFont sets if the style is applied to the font.
func (c *CellStyle) SetApplyFont(applyFont bool) *CellStyle {
	c.ApplyFont = applyFont
	return c
}

// SetApplyAlignment sets if the style is applied to the alignment.
func (c *CellStyle) SetApplyAlignment(applyAlignment bool) *CellStyle {
	c.ApplyAlignment = applyAlignment
	return c
}

// SetAlignementH sets the horizontal alignment.
func (c *CellStyle) SetAlignementH(alignmentH string) *CellStyle {
	c.Alignment.Horizontal = alignmentH
	return c
}

// SetIndent sets the indent.
func (c *CellStyle) SetIndent(indent int) *CellStyle {
	c.Alignment.Indent = indent
	return c
}

// SetShrinkToFit sets the shrink to fit.
func (c *CellStyle) SetShrinkToFit(shrinkToFit bool) *CellStyle {
	c.Alignment.ShrinkToFit = shrinkToFit
	return c
}

// SetTextRotation sets the text rotation.
func (c *CellStyle) SetTextRotation(textRotation int) *CellStyle {
	c.Alignment.TextRotation = textRotation
	return c
}

// SetAlignementV sets the vertical alignment.
func (c *CellStyle) SetAlignementV(alignmentV string) *CellStyle {
	c.Alignment.Vertical = alignmentV
	return c
}

// SetWrapText sets the wrap text.
func (c *CellStyle) SetWrapText(wrapText bool) *CellStyle {
	c.Alignment.WrapText = wrapText
	return c
}
