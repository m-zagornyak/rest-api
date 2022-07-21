package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {

}

func InitRouter() {
	
}


func RouterTodo() {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.GET(SingIn, h.SingIn)
		auth.POST("/sing-up")
	}
	item := router.Group("/item")
	{
		item.GET("/:id")
		item.POST("/:id")

	}
}

