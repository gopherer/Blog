package Model

type Sort struct {
	SortID   int    `form:"sort_id" xorm:"int pk autoincr comment('分类ID')"`
	SortName string `form:"sort_name" xorm:"varchar(20) unique notnull comment('分类名称')"`
}
