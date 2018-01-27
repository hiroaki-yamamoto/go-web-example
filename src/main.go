package main

import "github.com/gin-gonic/gin"

func main() {
  svr := gin.Default()
  svr.GET("/hello", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "Hello World",
    })
  })
  svr.Run()
}
