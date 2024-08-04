package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/mondoo", func(c *gin.Context) {
      message := "Hello from Mondoo Engineer!"
      c.String(200, message)
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
