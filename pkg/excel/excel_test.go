package excel

import (
	"github.com/xuri/excelize/v2"
	"reflect"
	"testing"
)

func TestLetter(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Letter(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Letter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMyExcel(t *testing.T) {
	tests := []struct {
		name string
		want *lzExcelExport
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMyExcel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMyExcel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createFile(t *testing.T) {
	tests := []struct {
		name string
		want *excelize.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createFileName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFileName(); got != tt.want {
				t.Errorf("createFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ExportToPath(t *testing.T) {
	type args struct {
		params []map[string]string
		data   []map[string]interface{}
		path   string
	}
	var tests = []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				params: []map[string]string{
					{
						"key":    "id",
						"title":  "索引",
						"width":  "20",
						"is_num": "0",
					},
				},
				data: []map[string]interface{}{
					{
						"id": 1,
					},
				},
				path: "./",
			},
			want:    "./test.xlsx",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewMyExcel()
			_, err := l.ExportToPath(tt.args.params, tt.args.data, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExportToPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_lzExcelExport_export(t *testing.T) {
	type fields struct {
		file      *excelize.File
		sheetName string
	}
	type args struct {
		params []map[string]string
		data   []map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lzExcelExport{
				file:      tt.fields.file,
				sheetName: tt.fields.sheetName,
			}
			l.export(tt.args.params, tt.args.data)
		})
	}
}

func Test_lzExcelExport_writeData(t *testing.T) {
	type fields struct {
		file      *excelize.File
		sheetName string
	}
	type args struct {
		params []map[string]string
		data   []map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lzExcelExport{
				file:      tt.fields.file,
				sheetName: tt.fields.sheetName,
			}
			l.writeData(tt.args.params, tt.args.data)
		})
	}
}

func Test_lzExcelExport_writeTop(t *testing.T) {
	type fields struct {
		file      *excelize.File
		sheetName string
	}
	type args struct {
		params []map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &lzExcelExport{
				file:      tt.fields.file,
				sheetName: tt.fields.sheetName,
			}
			l.writeTop(tt.args.params)
		})
	}
}
