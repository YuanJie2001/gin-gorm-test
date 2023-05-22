package model

type Book struct {
	ID     int64
	Name   string // 注意这里的字段首字母必须大写，否则无法被json包访问
	Age    int8
	Gender string
}
