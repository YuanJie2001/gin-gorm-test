package model

type User struct {
	Name   string // 注意这里的字段首字母必须大写，否则无法被json包访问
	Age    int
	Gender string
}
