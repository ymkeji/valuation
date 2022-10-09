package excelx

type Excelize interface {
	Search(value string, reg ...bool) (res []string, err error)
	ReadDate(params []map[string]string) (data []map[string]interface{}, err error)
	WriteDate(params []map[string]string, data []map[string]interface{}) (err error)
}
