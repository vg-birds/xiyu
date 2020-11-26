package models

import (
	db "test/database"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

type Blog struct {
	Id int `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	ViewNum int `json:"viewNum" form:"viewNum"`
	PublishTime int `json:"publishTime" form:"publishTime"`
}

type BlogTpl struct {
	Id int `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	ViewNum int `json:"viewNum" form:"viewNum"`
	Date string `json:"PublishDate" form:"PublishDate"`
}

func (b *Blog) GetBlogs (page, pageSize int) (blogs []BlogTpl, err error){
	var blog []Blog
	var blogTpl BlogTpl
	offset := (page - 1) * pageSize
	db.SqlDB.Select("id,title,content,viewNum,publishTime").Offset(offset).Limit(pageSize).Find(&blog)
	for _,v := range blog{
		blogTpl.Id = v.Id
		blogTpl.ViewNum = v.ViewNum
		blogTpl.Content = v.Content
		blogTpl.Title = v.Title

		if time.Now().Unix() - int64(v.PublishTime) <=30 {
			blogTpl.Date = "刚刚"
		}else{
			blogTpl.Date = time.Unix(int64(v.PublishTime), 0).Format(TimeLayout)
		}
		blogs = append(blogs, blogTpl)
	}

	return
}

func (b *Blog) GetBlog (id int) (blogTpl BlogTpl, err error)  {
	var blog Blog
	db.SqlDB.Where("id=?", id).First(&blog)

	blogTpl.Id = blog.Id
	blogTpl.Title = blog.Title
	blogTpl.Content = blog.Content
	blogTpl.ViewNum = blog.ViewNum
	if time.Now().Unix() - int64(blog.PublishTime) <=30 {
		blogTpl.Date = "刚刚"
	}else{
		blogTpl.Date = time.Unix(int64(blog.PublishTime), 0).Format(TimeLayout)
	}

	return
}
