package model

type Blog struct {
	BlogTitle      string `form:"blog_title" xorm:"pk varchar(45)"`
	BlogPhoto      string `form:"blog_photo" xorm:"varchar(255)"`
	BlogContent    string `form:"blog_content" xorm:"varchar(3000)"`
	BlogCreateTime string `form:"blog_create_time" xorm:"varchar(30) notnull"`
	BlogModifyTime string `form:"blog_modify_time" xorm:"varchar(30)"`
}
