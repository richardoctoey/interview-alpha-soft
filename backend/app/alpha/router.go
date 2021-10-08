package http

import (
	"backend/app/alpha/handler/music"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/music", music.AddMusic)
	r.PATCH("/music/:id", music.UpdateMusic)
	r.DELETE("/music/:id", music.DeleteMusic)
	r.GET("/music", music.ListMusic)
}
