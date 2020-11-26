package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "test/apis"
)

func initRouter() *gin.Engine{
	router := gin.Default()

	router.LoadHTMLGlob("tpl/html/*")

	router.StaticFS("/staticfile", http.Dir("./static/res/static"))
	router.StaticFS("/layui", http.Dir("./static/res/layui"))
	router.StaticFS("/statichtml", http.Dir("./tpl/html"))

	router.GET("/blogs", GetBlogs)

	//router.GET("/", IndexApi)
	//
	//router.GET("/person/:id", GetPersonApi)
	//
	//router.POST ("/person", AddPersonApi)
	//
	//router.GET("/persons", GetPersonsApi)
	//
	//router.PUT("/person/:id", ModPersonApi)
	//
	//router.DELETE("/person/:id", DelPersonApi)

	return router
}
