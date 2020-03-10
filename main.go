package main

import (
	"data/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoueter() *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("view/*")

	router.POST("/post",controllers.RegisterPost)

	router.GET("/",controllers.RegisterGet)

	router.GET("/get",controllers.Get)

	return router

}

func  main(){
	router:=InitRoueter()

	router.Static("/static","static")

	router.Run(":8030")

}
