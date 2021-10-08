package http

import (
	"backend/app/http/handler/music"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/music/add", music.AddMusic)
	r.PATCH("/music/update/:id", music.UpdateMusic)
	r.DELETE("/music/delete/:id", music.DeleteMusic)
	r.GET("/music", music.ListMusic)
}
