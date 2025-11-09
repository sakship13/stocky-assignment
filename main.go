package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"stocky/db"
	"stocky/routes"
	"stocky/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment vars")
	}

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	db.Init()
	services.StartPriceUpdater()

	router := gin.Default()
	routes.RegisterRoutes(router)

	logrus.Info("🚀 Server running on port 8080")
	router.Run(":8080")
}
