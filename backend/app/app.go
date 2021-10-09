package app

import (
	"backend/app/alpha"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func StartApp() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowAllOrigins: false,
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	http.SetupRouter(router)
	router.Run(":8080")
}
