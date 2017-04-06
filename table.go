package rtfdoc

import "fmt"

func getDefaultTableProperties() TableProperties {
	tp := TableProperties{
		align: "c",
	}
	tp.SetMargins(100, 100, 100, 100)
	return tp
}

func NewTable() Table {
	return Table{TableProperties: getDefaultTableProperties()}
}

func (t *Table) AddRow(row TableRow) {
	t.Data = append(t.Data, row)
}

func (t *TableProperties) SetMargins(left, top, right, bottom int) {
	margins := ""
	if left != 0 {
		margins += fmt.Sprintf(" \\trpaddl%d", left)
	}
	if top != 0 {
		margins += fmt.Sprintf(" \\trpaddt%d", top)
	}
	if right != 0 {
		margins += fmt.Sprintf(" \\trpaddr%d", right)
	}
	if bottom != 0 {
		margins += fmt.Sprintf(" \\trpaddb%d", bottom)
	}
	margins += " "
	t.margins = margins
}

func (t *TableProperties) getMargins() string {
	return t.margins
}

func (t Table) Compose() string {
	res := ""
	var align = ""
	if t.align != "" {
		align = fmt.Sprintf("\\trq%s", t.align)
	}
	for _, tr := range t.Data {
		res += fmt.Sprintf("\n{\\trowd %s", align)
		res += t.getMargins()
		res += tr.Compose()
		res += "\n\\row}"
	}
	return res
}

func NewTableRow() TableRow {
	return TableRow{}
}
func (tr *TableRow) AddCell(cell TableCell) {
	*tr = append(*tr, cell)
}

func (tr TableRow) Compose() string {
	res := ""
	if len(tr) != 0 {
		cBegin := 0
		for _, dc := range tr {
			cBegin += dc.getCellWidth()
			res += fmt.Sprintf("\n%s %s %s \\cellx%v", dc.getVerticalMergedProperty(), dc.getCellMargins(), dc.getBorders(), cBegin)

		}
		for _, dc := range tr {
			res += dc.cellCompose()
		}
	}
	return res
}

func NewDataCell(width int) DataCell {
	cp := CellProperties{}
	cp.CellWidth = width
	dc := DataCell{
		Cell{
			content:        Paragraph{},
			CellProperties: cp,
		},
	}
	dc.SetBorders(true, true, true, true)
	return dc
}
func NewDataCellWithProperties(cp CellProperties) DataCell {
	return DataCell{Cell{
		content:        Paragraph{},
		CellProperties: cp,
	}}
}

func (cp *CellProperties) SetProperties(cellWidth int, borders string) {
	cp.CellWidth = cellWidth
	cp.borders = borders
	return
}

func (dc *DataCell) SetContent(c Paragraph) {
	dc.content = c
}

func (dc DataCell) cellCompose() string {
	res := fmt.Sprintf("\n\\pard\\intbl %s \\cell", dc.Cell.content.CellCompose())

	return res
}

func (dc DataCell) getCellWidth() int {
	return dc.CellWidth
}

func (dc *DataCell) SetBorders(left, top, right, bottom bool) {
	b := ""
	bTemplStr := "\\clbrdr%s\\brdrw15\\brdrs"
	if left {
		b += fmt.Sprintf(bTemplStr, "l")
	}
	if top {
		b += fmt.Sprintf(bTemplStr, "t")
	}
	if right {
		b += fmt.Sprintf(bTemplStr, "r")
	}
	if bottom {
		b += fmt.Sprintf(bTemplStr, "b")
	}
	dc.borders = b
}

func (dc DataCell) getBorders() string {
	return dc.borders
}

func (tp *TableProperties) SetAlign(align string) {
	switch align {
	case "c", "center":
		tp.align = "c"
	case "l", "left":
		tp.align = "l"
	case "r", "right":
		tp.align = "r"
	default:
		tp.align = ""
	}
}

func (tp *TableProperties) GetAlign() string {
	return tp.align
}

func GetTableCellWidthByRatio(tableWidth int, ratio ...float64) []int {

	cellRatioSum := 0.0
	for _, cellRatio := range ratio {
		cellRatioSum += cellRatio
	}
	var cellWidth = make([]int, len(ratio))
	for i := range ratio {
		cellWidth[i] = int(ratio[i] * (float64(tableWidth) / cellRatioSum))
	}
	return cellWidth
}

func (dc *DataCell) SetVerticalMerged(isFirst, isNext bool) {
	if isFirst {
		dc.VerticalMerged.code = "\\clvmgf"
	}
	if isNext {
		dc.VerticalMerged.code = "\\clvmrg"
	}
}

func (dc DataCell) getVerticalMergedProperty() string {
	return dc.VerticalMerged.code
}

func (dc *DataCell) SetCellMargins(left, top, right, bottom int) {
	m := ""
	if left != 0 {
		m += fmt.Sprintf("\\clpadl%d", left)
	}
	if top != 0 {
		m += fmt.Sprintf("\\clpadt%d", top)
	}
	if right != 0 {
		m += fmt.Sprintf("\\clpadr%d", right)
	}
	if bottom != 0 {
		m += fmt.Sprintf("\\clpadb%d", bottom)
	}
	dc.margins = m
}

func (dc DataCell) getCellMargins() string {
	return dc.margins
}