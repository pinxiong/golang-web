package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()
	router.GET("/hello", func(context *gin.Context) {
		user := context.Query("user")
		context.JSON(200, gin.H{
			"message": "Hello " + user,
		})
	})
	router.Run()
}
