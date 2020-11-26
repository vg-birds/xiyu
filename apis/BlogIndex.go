package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"test/models"
)

const pageSize = 10

func GetBlogs(c *gin.Context)  {
	page := c.Request.FormValue("page")
	pages, err := strconv.Atoi(page)
	if err != nil{
		log.Fatalln(err)
	}
	var b models.Blog
	blogs, err := b.GetBlogs(pages, pageSize)

	if err != nil {
		log.Fatalln(err)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"blogs":blogs,
	})
}

func GetBlog(c *gin.Context)  {
	data := c.Request.FormValue("id")
	id, _ := strconv.Atoi(data)

	var b models.Blog
	blog,_ := b.GetBlog(id)

	c.HTML(http.StatusOK, "detail.html", gin.H{
		"blog":blog,
	})
}
