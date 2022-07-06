package Service

import (
	"blog/DataAccess"
	"blog/model"
	"github.com/wonderivan/logger"
)

type BlogService struct {
}

func (blogService *BlogService) PostPublish(reqBlog model.Blog) error {
	err := new(DataAccess.BlogDao).SetBlogInfo(reqBlog)
	if err != nil {
		return err
	}
	return nil
}

func (blogService *BlogService) PostModify(reqBlog model.Blog) error {
	new(DataAccess.BlogDao).UpdateBlogInfo(reqBlog)
	return nil
}

func (blogService *BlogService) PostDelete(reqBlog model.Blog) error {
	new(DataAccess.BlogDao).DeleteBlog(reqBlog)
	return nil
}

func (blogService *BlogService) BlogAll() []model.Blog {
	var blog []model.Blog
	//var mapBlog = make([]map[string]interface{}, 1)
	rowsBlog := new(DataAccess.BlogDao).AllBlog()
	for rowsBlog.Next() {
		tem := model.Blog{}
		err := rowsBlog.Scan(&tem)
		if err != nil {
			logger.Error(err)
		}
		blog = append(blog, tem)
	}
	return blog
}
