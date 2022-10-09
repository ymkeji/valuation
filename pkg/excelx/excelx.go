package excelx

import (
	"fmt"
	"io"
	"strconv"

	"github.com/xuri/excelize/v2"
)

var (
	defaultSheet   = "Sheet1" //默认Sheet名称
	defaultHeight  = 25.0     //默认行高度
	defaultHeadRow = 1
)

type ExcelX struct {
	f       *excelize.File
	Sheet   string
	HeadRow int
}

type ExcelErr struct {
	Row int    `json:"row"`
	Msg string `json:"msg"`
}

type Option func(excel *ExcelX)

func NewExcel(opts ...Option) Excelize {
	excelX := &ExcelX{
		f:       createFile(),
		Sheet:   defaultSheet,
		HeadRow: defaultHeadRow,
	}

	for _, opt := range opts {
		opt(excelX)
	}

	return excelX
}

func SetFile(r io.Reader, opt ...excelize.Options) Option {
	return func(excel *ExcelX) {
		excel.f, _ = excelize.OpenReader(r, opt...)
	}
}

func SetHeadRow(headRow int) Option {
	return func(excel *ExcelX) {
		excel.HeadRow = headRow
	}
}

func (e *ExcelX) Search(value string, reg ...bool) (res []string, err error) {
	res, err = e.f.SearchSheet(e.Sheet, value, reg...)
	if err != nil {
		return
	}
	return
}

func (e *ExcelX) ReadDate(params []map[string]string) (data []map[string]interface{}, err error) {
	rows, err := e.f.Rows(e.Sheet)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for r, rNo := 0, 1; rows.Next(); rNo++ {
		//数据开始行数
		if e.HeadRow >= rNo {
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

func (e *ExcelX) WriteDate(params []map[string]string, data []map[string]interface{}) (err error) {
	lineStyle, err := e.f.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	if err != nil {
		return err
	}
	//数据写入
	var j = 2 //数据开始行数
	for i, val := range data {
		//设置行高
		if err = e.f.SetRowHeight(e.Sheet, i+1, defaultHeight); err != nil {
			return err
		}
		//逐列写入
		var word = 'A'
		for _, conf := range params {
			valKey := conf["key"]
			line := fmt.Sprintf("%c%v", word, j)
			isNum := conf["is_num"]
			//设置值
			if isNum != "0" {
				valNum := fmt.Sprintf("'%v", val[valKey])
				if err = e.f.SetCellValue(e.Sheet, line, valNum); err != nil {
					return err
				}
			} else {
				if err = e.f.SetCellValue(e.Sheet, line, val[valKey]); err != nil {
					return err
				}
			}
			//设置样式
			if err = e.f.SetCellStyle(e.Sheet, line, line, lineStyle); err != nil {
				return err
			}
			word++
		}
		j++
	}
	//设置行高 尾行
	err = e.f.SetRowHeight(e.Sheet, len(data)+1, defaultHeight)
	return
}

func createFile() *excelize.File {
	f := excelize.NewFile()
	index := f.NewSheet(defaultSheet)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	return f
}
