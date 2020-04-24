package router

import (
	"github.com/gin-gonic/gin"
	"go-redis/controller"
)

func Router(r *gin.Engine) *gin.Engine  {
	r.POST("/set",controller.Setredis)
	r.POST("/get",controller.Getredis)
	r.GET("/qimao",controller.Qimao)
	r.GET("/qimaopage",controller.QimaoPage)
	r.GET("/Qimaoget",controller.Qimaoget)
	return r
}
