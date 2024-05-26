package model

type FieldType string

const (
	FieldTypeInt    FieldType = "int"
	FieldTypeString FieldType = "string"
)

type Database struct {
	//数据库名
	Name   string  `json:"name"`
	Tables []Table `json:"tables"`
}

type Table struct {
	//表名
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	//时间戳
	Timestamp int64 `json:"timestamp"`
	//字段名
	Name string `json:"name"`
	//字段值,照理说不该存在字符串类型。对于float类型，要求使用者提供规则转化为整形进行存储
	Value int64 `json:"value"`
}
