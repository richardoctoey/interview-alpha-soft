package app

import (
	"backend/app/alpha"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()
	http.SetupRouter(router)
	router.Run(":8080")
}
