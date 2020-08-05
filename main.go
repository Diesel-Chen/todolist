package main

import (
	"todolist/controller"
	"todolist/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化router
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	// 使用跨域中间件
	router.Use(middleware.Cors())

	// router.Static("/static", "./build/static")
	// router.StaticFile("/index.html", "./build/index.html")

	v1 := router.Group("/api/v1/todolist")
	{
		v1.POST("/add", controller.AddTodo)
		v1.GET("/selectAll", controller.SelectAllTodo)
		v1.GET("/selectCompleted", controller.SelectCompleted)
		v1.GET("/selectSingle/:id", controller.SelectSingleTodo)
		v1.PUT("/updateState/:id", controller.UpdateStateTodo)
		v1.PUT("/updateTitle/:id", controller.UpdateTitleTodo)
		v1.DELETE("/delete/:id", controller.DeleteTodo)
	}

	router.Run(":80")
}
