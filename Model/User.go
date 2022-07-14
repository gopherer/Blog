package Model

// form tag 可接收html form表单 application/x-www-form-urlencoded的内容 也可以接受json文本的内容
//json tag 只可以接收json的文本内容
//可以使用xorm框架中 xorm tag标签将 结构体映射成数据库中的表
//binding tag 可用 但不好 最好专门写个函数处理前端输入的数据
//例子 Name string  `xorm:"varchar(25) notnull unique 'usr_name' comment('姓名')"`
type User struct {
	UserAccount string `form:"user_account" xorm:"pk varchar(20)"`
	UserPwd     string `form:"user_pwd" xorm:"varchar(20) notnull " `
	UserNick    string `form:"user_nick" xorm:"varchar(20) comment('昵称')"`
	UserIcon    string `form:"user_icon" xorm:"varchar(100) comment('头像')"`
	UserProfile string `form:"user_profile" xorm:"varchar(200) comment('用户简介（签名）')"`
	UserContact string `form:"user_contact" xorm:"varchar(50) comment('联系方式')"`
}
