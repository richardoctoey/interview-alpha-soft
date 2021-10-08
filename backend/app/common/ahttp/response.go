package ahttp

import gin "github.com/gin-gonic/gin"

func ResponseFailed(c *gin.Context, message string, obj interface{}) {
	c.JSON(200, gin.H{
		"result": false,
		"message": message,
		"data": obj,
	})
	return
}

func ResponseSuccess(c *gin.Context, message string, obj interface{}) {
	c.JSON(200, gin.H{
		"result": true,
		"message": message,
		"data": obj,
	})
	return
}