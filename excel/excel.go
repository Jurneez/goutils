package excel

import (
	"github.com/tealeg/xlsx"
)

type Excel struct {
	Sheets []*Sheet
}
type Sheet struct {
	titles []string // excel的头部名称
	cells  [][]interface{}
}

func NewExcel() *Excel {
	return &Excel{
		Sheets: make([]*Sheet, 0),
	}
}

func (e *Excel) SetSheet() {

}

func Excel1() {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value = "haha"
	cell = row.AddCell()
	cell.Value = "xixi"

	err := file.Save("./file.xlsx")
	if err != nil {
		panic(err)
	}

}
