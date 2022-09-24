package excel

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

var (
	defaultSheetName = "Sheet1" //默认Sheet名称
	defaultHeight    = 25.0     //默认行高度
)

type LzExcelExport struct {
	file      *excelize.File
	sheetName string //可定义默认sheet名称
}

type ExcelErr struct {
	Row int
	Msg string
}

func NewMyExcel() *LzExcelExport {
	return &LzExcelExport{file: createFile(), sheetName: defaultSheetName}
}

func ReadMyExcel(r io.Reader, opt ...excelize.Options) (*LzExcelExport, error) {
	f, err := excelize.OpenReader(r, opt...)
	if err != nil {
		return nil, err
	}
	return &LzExcelExport{
		file:      f,
		sheetName: defaultSheetName,
	}, nil
}

func (l *LzExcelExport) Search(value string, reg ...bool) (res []string, err error) {
	res, err = l.file.SearchSheet(l.sheetName, value, reg...)
	if err != nil {
		return
	}
	return
}

// 导出基本的表格
func (l *LzExcelExport) ExportToPath(params []map[string]string, data []map[string]interface{}, path string) (string, error) {
	l.export(params, data)
	name := CreateFileName()
	filePath := path + "/" + name
	err := l.file.SaveAs(filePath)
	return filePath, err
}

// 设置首行
func (l *LzExcelExport) writeTop(params []map[string]string) {
	topStyle, _ := l.file.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center","vertical":"center"}}`)
	var word = 'A'
	//首行写入
	for _, conf := range params {
		title := conf["title"]
		width, _ := strconv.ParseFloat(conf["width"], 64)
		line := fmt.Sprintf("%c1", word)
		//设置标题
		_ = l.file.SetCellValue(l.sheetName, line, title)
		//列宽
		_ = l.file.SetColWidth(l.sheetName, fmt.Sprintf("%c", word), fmt.Sprintf("%c", word), width)
		//设置样式
		_ = l.file.SetCellStyle(l.sheetName, line, line, topStyle)
		word++
	}
}

// 写入数据
func (l *LzExcelExport) writeData(params []map[string]string, data []map[string]interface{}) {
	lineStyle, _ := l.file.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	//数据写入
	var j = 2 //数据开始行数
	for i, val := range data {
		//设置行高
		_ = l.file.SetRowHeight(l.sheetName, i+1, defaultHeight)
		//逐列写入
		var word = 'A'
		for _, conf := range params {
			valKey := conf["key"]
			line := fmt.Sprintf("%c%v", word, j)
			isNum := conf["is_num"]

			//设置值
			if isNum != "0" {
				valNum := fmt.Sprintf("'%v", val[valKey])
				_ = l.file.SetCellValue(l.sheetName, line, valNum)
			} else {
				_ = l.file.SetCellValue(l.sheetName, line, val[valKey])
			}

			//设置样式
			_ = l.file.SetCellStyle(l.sheetName, line, line, lineStyle)
			word++
		}
		j++
	}
	//设置行高 尾行
	_ = l.file.SetRowHeight(l.sheetName, len(data)+1, defaultHeight)
}

// ReadData 读取数据
func (l *LzExcelExport) ReadData(params []map[string]string, headRow int) (data []map[string]interface{}, err error) {
	rows, err := l.file.Rows(l.sheetName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for r, rNo := 0, 1; rows.Next(); rNo++ {
		//数据开始行数
		if headRow >= rNo {
			continue
		}
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if len(row) == 0 || len(row) < len(params) {
			break
		}
		data = append(data, map[string]interface{}{})
		for _, conf := range params {
			i, _ := strconv.Atoi(conf["id"])
			data[r][conf["key"]] = row[i-1]
		}
		r++
	}
	return

}

func (l *LzExcelExport) export(params []map[string]string, data []map[string]interface{}) {
	l.writeTop(params)
	l.writeData(params, data)
}

func createFile() *excelize.File {
	f := excelize.NewFile()
	// 创建一个默认工作表
	sheetName := defaultSheetName
	index := f.NewSheet(sheetName)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	return f
}

func CreateFileName() string {
	name := time.Now().Format("2006-01-02-15-04-05")
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("excle-%v-%v.xlsx", name, rand.Int63n(time.Now().Unix()))
}

// Letter 遍历a-z
func Letter(length int) []string {
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, string(rune('A'+i)))
	}
	return str
}
