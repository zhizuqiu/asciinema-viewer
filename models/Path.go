package models

// Model Struct
type Path struct {
	Id   int
	Path string `orm:"size(100)"`
}
