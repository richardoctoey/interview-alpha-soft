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
	result, err := model.ArtistList()
	if err != nil {
		logger.Error(err, nil)
		ahttp.ResponseFailed(c, err.Error(), nil)
		return
	}
	ahttp.ResponseSuccess(c, "success", result)
	return
}

func UpdateMusic(c *gin.Context) {
	formRequest := model.MusicFormRequest{}
	if err := c.ShouldBind(&formRequest); err != nil {
		logger.Error(err, nil)
		ahttp.ResponseFailed(c, err.Error(), nil)
		return
	}
	artist := model.Artist{}
	err := artist.FindArtistById(c.Param("id"))
	if err != nil {
		logger.Error(err, formRequest)
		ahttp.ResponseFailed(c, err.Error(), nil)
		return
	}
	artist.ParseFromRequest(formRequest)
	if err := artist.Update(); err != nil {
		logger.Error(err, artist)
		ahttp.ResponseFailed(c, err.Error(), artist)
		return
	}
	ahttp.ResponseSuccess(c, "success", artist)
	return
}

func DeleteMusic(c *gin.Context) {
	formRequest := model.MusicFormRequest{}
	if err := c.ShouldBindUri(&formRequest); err != nil {
		logger.Error(err, nil)
		ahttp.ResponseFailed(c, err.Error(), nil)
		return
	}
	if err := model.ArtistDeleteById(c.Param("id")); err != nil {
		logger.Error(err, c.Param("id"))
		ahttp.ResponseFailed(c, err.Error(), c.Param("id"))
		return
	}
	ahttp.ResponseSuccess(c, "success", nil)
	return
}