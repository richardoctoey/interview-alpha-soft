package music

import (
	"backend/app/common/http"
	"backend/app/common/logger"
	"github.com/gin-gonic/gin"
)

func AddMusic(c *gin.Context) {
	formRequest := MusicFormRequest{}
	if err := c.BindJSON(&formRequest); err != nil {
		logger.Error(err, nil)
		http.ResponseFailed(c, err.Error(), nil)
		return
	}
}

func ListMusic(c *gin.Context) {

}

func UpdateMusic(c *gin.Context) {

}

func DeleteMusic(c *gin.Context) {

}