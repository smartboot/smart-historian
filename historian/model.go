package historian

type Request struct {
	//数据库名
	Database string `json:"database"`
	//表名
	Table string `json:"table"`
	//字段名
	Fields []Field `json:"fields"`
}

type Field struct {
	//时间戳
	Timestamp int64 `json:"timestamp"`
	//字段名
	Name string `json:"name"`
	//字段值
	Value interface{} `json:"value"`
}

type IntField struct {
	Field
	Value int64 `json:"value"`
}

type StringField struct {
	Field
	Value string `json:"value"`
}

type FloatField struct {
	Field
	Value float64 `json:"value"`
}
