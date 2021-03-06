package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func RouterTodo() {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.GET("/sing-in")
		auth.POST("/sing-up")
	}
	item := router.Group("/item")
	{
		item.GET("/item:id")
		item.POST("")
	}
}
