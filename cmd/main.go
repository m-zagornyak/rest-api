package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()



	r := gin.Default()
	r.Run(":8080")
}
