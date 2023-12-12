package router

import "github.com/gin-gonic/gin"

func Routeer() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)

	return r
}
