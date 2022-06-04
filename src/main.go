package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
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

var keepRunning = true

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/start", start)
	router.GET("/stop", stop)
	router.GET("/health", status)
	router.GET("/check", status)
	router.GET("/status", status)
	router.Run()
}

func start(context *gin.Context) {
	keepRunning = true
	go func() {
		log.Info("It is starting....")
		for keepRunning {
			time.Sleep(time.Duration(1) * time.Nanosecond)
		}
		log.Info("It stopped.")
	}()
	context.JSON(200, gin.H{
		"message": "A CPU-consuming task is started!",
	})
}

func stop(context *gin.Context) {
	keepRunning = false
	context.JSON(200, gin.H{
		"message": "All started tasks have been stopped!",
	})
}

func status(context *gin.Context) {
	context.JSON(200, gin.H{
		"status":      "Healthy",
		"keepRunning": keepRunning,
	})
	log.Info("Receive a request to check service's status!")
}
