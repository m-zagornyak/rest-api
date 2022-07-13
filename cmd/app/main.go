package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.

	r := gin.Default()
	r.Run()
}
