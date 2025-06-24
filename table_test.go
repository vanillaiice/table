package table_test

import (
	"os"
	"testing"
	"time"

	"github.com/vanillaiice/table"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestAll(t *testing.T) {
	tbl := table.NewTable()

	sh, err := tbl.AddSheet("demo")
	if err != nil {
		t.Fatal(err)
	}

	sh.AddRows(table.RowTitle("a normal row", 3, "thin"))
	row := sh.NewRow().AddCells(table.NewCells(table.NewCellOpts(), "a", "b", "c")...)
	table.AddOuterBorderRows("medium", row)

	sh.NewRow()
	sh.AddRows(table.RowTitle("row with horizontally merged cells (with outer border)", 9, "thin"))
	row = table.RowWithMergedCellsH(
		table.NewCell("foo", table.NewCellOpts(table.NewCellStyle().SetUnderline(true)).SetMergeH(3)),
		table.NewCell("bar", table.NewCellOpts(table.NewCellStyle().SetBold(true)).SetMergeH(2)),
		table.NewCell("baz", table.NewCellOpts(table.NewCellStyle().SetItalic(true)).SetMergeH(4)),
	)
	table.AddOuterBorderRows("thin", row)
	sh.AddRows(row)

	sh.NewRow()
	sh.AddRows(table.RowTitle("row with horizontally merged cells (with inner & outer borders)", 15, "thin"))
	row = table.RowWithMergedCellsH(table.NewCells(table.NewCellOpts(table.NewCellStyle().SetFontName("Luxi Mono")).SetMergeH(5), "foo", "bar", "baz")...)
	table.AddOuterAndInnerBorderRow("thin", row)
	sh.AddRows(row)

	sh.NewRow()
	sh.AddRows(table.RowTitle("row with vertically merged cells", 3, "thin").SetHeight(40))
	rows := table.RowWithMergedCellsV(table.NewCells(table.NewCellOpts(table.NewCellStyle().SetFontSize(16).SetFontName("URW Gothic")).SetMergeV(5), "foo", "bar", "baz")...)
	table.AddOuterBorderRows("thin", rows...)
	sh.AddRows(rows...)

	sh.NewRow()
	sh.AddRows(table.RowTitle("row with vertically merged cells (different merge values)", 2, "thin", table.NewCellStyle().SetWrapText(true).SetFontSize(12).SetBold(true)))
	rows = table.RowWithMergedCellsV(
		table.NewCell("fa", table.NewCellOpts().SetMergeV(2)),
		table.NewCell("fb", table.NewCellOpts().SetMergeV(3)),
	)
	// we manually set the content of the first cell of the last row.
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
	rows = table.RowWithMergedCells(table.NewCells(table.NewCellOpts(table.NewCellStyle().SetFontColor("ff660066")).SetMergeV(2).SetMergeH(3), "fff", "bbb", "zzz")...)
	table.AddOuterBorderRows("thick", rows...)
	sh.AddRows(rows...)

	sh.NewRow()
	rows = table.GenericTable(
		"a generic table",
		table.NewCells(table.NewCellOpts(), "f", "o", "o"),
		[]*table.Row{
			table.NewRow().AddCells(table.NewCells(table.NewCellOpts(), "a", "b", "c")...),
			table.NewRow().AddCells(table.NewCells(table.NewCellOpts(), "d", "e", "f")...),
			table.NewRow().AddCells(table.NewCells(table.NewCellOpts(), "g", "h", "i")...),
		},
		"thin",
	)
	table.AddOuterAndInnerBorderRow("thin", rows...)
	sh.AddRows(rows...)

	sh.NewRow()
	row = sh.NewRow().AddCells(table.NewCells(table.NewCellOpts(table.NewCellStyle().SetFillFgColor("ff005200").SetFillPatternType("solid").SetBordersColor("ffff69b4")), "b", "a", "r")...)
	table.AddInnerBorderRows("thick", row)

	sh.NewRow()
	sh.AddRows(table.RowTitle("a generic table with inner borders", 5, "thin"))
	sh.NewRow()
	rows = []*table.Row{
		table.NewRow().AddCells(table.NewCells(table.NewCellOpts(table.NewCellStyle().SetWrapText(true)), "j", "k", "l", 1, time.Now())...),
		table.NewRow().AddCells(table.NewCells(table.NewCellOpts(table.NewCellStyle().SetWrapText(true)), " m", "n", "o", 2, time.Now().Add(time.Hour))...),
		table.NewRow().AddCells(table.NewCells(table.NewCellOpts(table.NewCellStyle().SetWrapText(true)), "p", "q", "r", 3, time.Now().Add(2*time.Hour))...),
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
