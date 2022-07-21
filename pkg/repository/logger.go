package repository

import (
	"github.com/gin-gonic/gin"
	"os"
	
)

func logger() {
	gin.DisableConsoleColor()

    // Logging to a file.
  f, _ := os.Create("gin.log")
  gin.DefaultWriter = io.MultiWriter(f)
}