package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "os"
)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
  	host, _ := os.Hostname()
    c.JSON(http.StatusOK, gin.H{"hostname": host})
  })

  r.GET("/hello", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": "world"})
  })

  r.Run()
}
