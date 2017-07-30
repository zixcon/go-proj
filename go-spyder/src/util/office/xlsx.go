package office

//import (
//	"github.com/tealeg/xlsx"
//)
//
//func NewFile() *xlsx.File {
//	file := xlsx.NewFile()
//	return file
//}


//func (file *xlsx.File) NewSheet(name string) *xlsx.Sheet {
//	sheet, err := file.AddSheet("sheet1")
//	if err != nil {
//		panic(err)
//	}
//	return sheet
//}
//
//func (sheet *xlsx.Sheet) NewRow() *xlsx.Row {
//	row := sheet.AddRow()
//	return row
//}
//
//func (row *xlsx.Row) NewCell(value string) *xlsx.Cell {
//	cell := row..AddCell()
//	cell.Value = value
//	return cell
//}
//
//func (file *xlsx.File) Save(name string) {
//	file.Save(name)
//}

//func XlsxDemo() {
//	file := xlsx.NewFile()
//	sheet, err := file.AddSheet("sheet1")
//	if err != nil {
//		panic(err)
//	}
//	row := sheet.AddRow()
//	row.SetHeightCM(1) //设置每行的高度
//	cell := row.AddCell()
//	cell.Value = "haha"
//	cell = row.AddCell()
//	cell.Value = "xixi"
//
//	file.Save("file.xlsx")
//}
