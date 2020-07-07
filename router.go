package main

import (
	"github.com/gin-gonic/gin"
	."test/apis"
)

func initRouter() *gin.Engine{
	router := gin.Default()

	router.GET("/", IndexApi)

	router.GET("/person/:id", GetPersonApi)

	router.POST("/person", AddPersonApi)

	router.GET("/persons", GetPersonsApi)

	router.PUT("/person/:id", ModPersonApi)

	router.DELETE("/person/:id", DelPersonApi)
	return router
}
