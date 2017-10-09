package main

import (
	"fmt"
	"os"

	"github.com/therox/rtf-doc"
)

func main() {

	// Создаем новый чистый, незамутнённый документ

	d := rtfdoc.NewDocument()

	// Настроить хедер
	d.SetOrientation(rtfdoc.OrientationLandscape)
	d.SetFormat("A4")

	p := d.AddParagraph()
	p.AddText("Green first string (Times New Roman)", 48, rtfdoc.FontTimesNewRoman, rtfdoc.ColorGreen)
	d.AddParagraph().AddText("Blue second string (Arial)", 48, rtfdoc.FontArial, rtfdoc.ColorBlue)
	d.AddParagraph().AddText("Red Third string (Comic Sans)", 48, rtfdoc.FontComicSansMS, rtfdoc.ColorRed)

	// Таблица
	t := d.AddTable().SetWidth(10000)
	//t.SetLeftMargin(50).SetRightMargin(50).SetTopMargin(50).SetBottomMargin(50)
	t.SetMarginLeft(50).SetMarginRight(50).SetMarginTop(50).SetMarginBottom(50)
	t.SetBorderColor(rtfdoc.ColorGreen)

	// // строка таблицы
	tr := t.AddTableRow()

	// // Рассчет ячеек таблицы. Первый ряд
	cWidth := t.GetTableCellWidthByRatio(1, 3)

	// ячейка таблицы
	dc := tr.AddDataCell(cWidth[0])
	dc.SetVerticalMergedFirst()
	p = dc.AddParagraph()
	// текст
	p.AddText("Blue text with cyrillic support with multiline", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorBlue)
	p.AddNewLine()
	p.AddText("Голубой кириллический текст с переносом строки внутри параграфа", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorBlue)
	p.SetAlignt(rtfdoc.AlignJustify)
	p = dc.AddParagraph().SetIndent(40, 0, 0).SetAlignt(rtfdoc.AlignCenter)
	p.AddText("Another paragraph in vertical cell", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorBlue)

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph().SetAlignt(rtfdoc.AlignCenter)
	p.AddText("Green text In top right cell with center align", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorGreen)
	tr = t.AddTableRow()

	cWidth = t.GetTableCellWidthByRatio(1, 1.5, 1.5)
	// // Это соединенная с верхней ячейка. Текст в ней возьмется из первой ячейки.
	dc = tr.AddDataCell(cWidth[0])
	dc.SetVerticalMergedNext()

	dc = tr.AddDataCell(cWidth[1])
	p = dc.AddParagraph()
	p.SetAlignt(rtfdoc.AlignRight)
	p.AddText("Red text In bottom central cell with right align", 16, rtfdoc.FontArial, rtfdoc.ColorRed).SetBold()

	dc = tr.AddDataCell(cWidth[2])
	p = dc.AddParagraph()
	p.SetAlignt(rtfdoc.AlignLeft)
	p.AddText("Black text in bottom right cell with left align", 16, rtfdoc.FontComicSansMS, rtfdoc.ColorBlack).SetItalic()

	p = dc.AddParagraph()

	f, err := os.Open("pic.jpg")
	if err != nil {
		fmt.Println(err)
	}
	pPic1 := p.AddPicture(f, rtfdoc.ImageFormatJpeg)
	pPic1.SetWidth(100).SetHeight(100)
	p.SetAlignt(rtfdoc.AlignCenter)

	f, err = os.Open("pic.jpg")
	if err != nil {
		fmt.Println(err)
	}
	pPic := d.AddParagraph()
	pPic.AddPicture(f, rtfdoc.ImageFormatJpeg)
	pPic.SetAlignt(rtfdoc.AlignCenter)
	// pic.SetWidth(200).SetHeight(150)

	fmt.Println(string(d.Export()))

}