package table_test

import (
	"os"
	"testing"
	"time"

	"github.com/vanillaiice/table"
)

var tbl *table.Table

func TestMain(m *testing.M) {
	tbl = table.NewTable()

	os.Exit(m.Run())
}

func TestNewTable(t *testing.T) {
	sh, err := tbl.AddSheet("demo")
	if err != nil {
		t.Fatal(err)
	}

	sh.AddRows(table.RowTitle("a normal row", 3, "thin"))

	row := sh.NewRow().AddCells(table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle()}, "a", "b", "c")...)
	table.AddOuterBorderRows("medium", row)

	sh.NewRow()

	sh.AddRows(table.RowTitle("row with horizontally merged cells (with outer border)", 9, "thin"))

	row = table.RowWithMergedCellsH(
		&table.Cell{Content: "foo", MergeH: 3, Style: table.NewCellStyle().SetUnderline(true)},
		&table.Cell{Content: "bar", MergeH: 2, Style: table.NewCellStyle().SetBold(true)},
		&table.Cell{Content: "baz", MergeH: 4, Style: table.NewCellStyle().SetItalic(true)},
	)
	table.AddOuterBorderRows("thin", row)
	sh.AddRows(row)

	sh.NewRow()

	sh.AddRows(table.RowTitle("row with horizontally merged cells (with inner & outer borders)", 15, "thin"))

	row = table.RowWithMergedCellsH(table.NewCells(&table.NewCellOpts{MergeH: 5, Style: table.NewCellStyle().SetFontName("Luxi Mono")}, "foo", "bar", "baz")...)
	table.AddOuterAndInnerBorderRow("thin", row)
	sh.AddRows(row)

	sh.NewRow()

	sh.AddRows(table.RowTitle("row with vertically merged cells", 3, "thin").SetHeight(40))

	rows := table.RowWithMergedCellsV(table.NewCells(&table.NewCellOpts{MergeV: 5, Style: table.NewCellStyle().SetFontSize(16).SetFontName("URW Gothic")}, "foo", "bar", "baz")...)
	table.AddOuterBorderRows("thin", rows...)
	sh.AddRows(rows...)

	sh.NewRow()

	sh.AddRows(table.RowTitle("row with vertically merged cells (different merge values)", 2, "thin", table.NewCellStyle().SetWrapText(true).SetFontSize(12).SetBold(true)))

	rows = table.RowWithMergedCellsV(&table.Cell{Content: "fa", MergeV: 2, Style: table.NewCellStyle()}, &table.Cell{Content: "fb", MergeV: 3, Style: table.NewCellStyle()})

	// manually set the content of the first cell of the last row
	rows[len(rows)-1].IterCells(func(i int, c *table.Cell) error {
		if c.Content == nil {
			c.Content = "fc"
		}

		return nil
	})
	table.AddOuterBorderRows("thin", rows...)

	sh.AddRows(rows...)

	sh.NewRow()

	sh.AddRows(table.RowTitle("row with horizontally and vertically merged cells", 9, "thin"))

	rows = table.RowWithMergedCells(table.NewCells(&table.NewCellOpts{MergeH: 3, MergeV: 2, Style: table.NewCellStyle().SetFontColor("ff660066")}, "fff", "bbb", "zzz")...)
	table.AddOuterBorderRows("thick", rows...)
	sh.AddRows(rows...)

	sh.NewRow()

	rows = table.GenericTable(
		"a generic table",
		table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle()}, "f", "o", "o"),
		[]*table.Row{
			table.NewRow().AddCells(table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle()}, "a", "b", "c")...),
			table.NewRow().AddCells(table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle()}, "d", "e", "f")...),
			table.NewRow().AddCells(table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle()}, "g", "h", "i")...),
		},
		"thin",
	)
	table.AddOuterAndInnerBorderRow("thin", rows...)
	sh.AddRows(rows...)

	sh.NewRow()

	row = sh.NewRow().AddCells(table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle().SetFillFgColor("ff005200").SetFillPatternType("solid").SetBordersColor("ffff69b4 ")}, "b", "a", "r")...)
	table.AddInnerBorderRows("thick", row)

	sh.NewRow()

	sh.AddRows(table.RowTitle("a generic table with inner borders", 5, "thin"))

	sh.NewRow()

	rows = []*table.Row{
		table.NewRow().AddCells(table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle().SetWrapText(true)}, "j", "k", "l", 1, time.Now())...),
		table.NewRow().AddCells(table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle().SetWrapText(true)}, "m", "n", "o", 2, time.Now().Add(time.Hour))...),
		table.NewRow().AddCells(table.NewCells(&table.NewCellOpts{Style: table.NewCellStyle().SetWrapText(true)}, "p", "q", "r", 3, time.Now().Add(2*time.Hour))...),
	}
	table.AddInnerBorderRows("medium", rows...)
	sh.AddRows(rows...)

	if err := tbl.Render(); err != nil {
		t.Fatal(err)
	}

	if err := tbl.Save("test.xlsx"); err != nil {
		t.Fatal(err)
	}
}
