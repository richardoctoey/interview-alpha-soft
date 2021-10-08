package music

import (
	"backend/app/common/ahttp"
	"backend/app/common/logger"
	"backend/app/model"
	"github.com/gin-gonic/gin"
)

func AddMusic(c *gin.Context) {
	formRequest := model.MusicFormRequest{}
	if err := c.BindJSON(&formRequest); err != nil {
		logger.Error(err, nil)
		ahttp.ResponseFailed(c, err.Error(), nil)
		return
	}
	objMusic := new(model.Artist)
	objMusic.ParseFromRequest(formRequest)
	if err := objMusic.Create(); err != nil {
		ahttp.ResponseFailed(c, err.Error(), objMusic)
		return
	}
	ahttp.ResponseSuccess(c, "success", objMusic.ArtistId)
	return
}

func ListMusic(c *gin.Context) {

}

func UpdateMusic(c *gin.Context) {

}

func DeleteMusic(c *gin.Context) {
	formRequest := model.MusicFormRequest{}
	if err := c.ShouldBindUri(&formRequest); err != nil {
		logger.Error(err, nil)
		ahttp.ResponseFailed(c, err.Error(), nil)
		return
	}
	if err := model.ArtistDeleteById(formRequest.Id); err != nil {
		logger.Error(err, formRequest.Id)
		ahttp.ResponseFailed(c, err.Error(), formRequest.Id)
		return
	}
	ahttp.ResponseSuccess(c, "success", nil)
	return
}