package DataAccess

import (
	"blog/Tools"
	"blog/model"
	"errors"
	"github.com/go-xorm/xorm"
)

type BlogDao struct {
}

func (blogDao *BlogDao) SetBlogInfo(reqBlog model.Blog) error {
	result, _ := Tools.DbEngine.InsertOne(reqBlog) //result = 0 代表失败， 1 代表成功（存在成功一半的可能）
	if result == 0 {
		return errors.New("blog 写入数据库失败")
	}
	return nil
}

func (blogDao *BlogDao) UpdateBlogInfo(reqBlog model.Blog) {
	_, _ = Tools.DbEngine.ID(reqBlog.BlogTitle).Cols("blog_photo", "blog_content").Update(reqBlog) // update blog set blog_photo =  ， blog_content =  where title =  reqBlog.BlogTitle
}

func (blogDao *BlogDao) DeleteBlog(reqBlog model.Blog) {
	_, _ = Tools.DbEngine.ID(reqBlog.BlogTitle).Delete(reqBlog) //err 都为nil 如果blog中包含有bool,float64或者float32类型，有可能会使删除失败
}

func (blogDao *BlogDao) AllBlog() *xorm.Rows {
	var blog model.Blog
	rowsBlog, _ := Tools.DbEngine.Rows(&blog)
	return rowsBlog

}
