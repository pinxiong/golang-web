package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
			log.FieldKeyFunc:  "caller",
		},
		PrettyPrint: false,
	})

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/hello", func(context *gin.Context) {
		user := context.Query("user")
		context.JSON(200, gin.H{
			"message": "Hello " + user,
		})
		log.Infof("Received a request! user:%v", user)
	})
	router.Run()
}
