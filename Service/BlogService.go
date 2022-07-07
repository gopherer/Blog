package Service

import (
	"blog/DataAccess"
	"blog/Model"
	"github.com/wonderivan/logger"
)

type BlogService struct {
}

func (blogService *BlogService) PostPublish(reqBlog Model.Blog) error {
	err := new(DataAccess.BlogDao).SetBlogInfo(reqBlog)
	if err != nil {
		return err
	}
	return nil
}

func (blogService *BlogService) PostModify(reqBlog Model.Blog) error {
	new(DataAccess.BlogDao).UpdateBlogInfo(reqBlog)
	return nil
}

func (blogService *BlogService) PostDelete(reqBlog Model.Blog) error {
	new(DataAccess.BlogDao).DeleteBlog(reqBlog)
	return nil
}

func (blogService *BlogService) BlogAll() []Model.Blog {
	var blog []Model.Blog
	//var mapBlog = make([]map[string]interface{}, 1)
	rowsBlog := new(DataAccess.BlogDao).AllBlog()
	for rowsBlog.Next() {
		tem := Model.Blog{}
		err := rowsBlog.Scan(&tem)
		if err != nil {
			logger.Error(err)
		}
		blog = append(blog, tem)
	}
	return blog
}
