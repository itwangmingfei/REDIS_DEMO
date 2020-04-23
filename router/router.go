package router

import (
	"github.com/gin-gonic/gin"
	"go-redis/controller"
)

func Router(r *gin.Engine) *gin.Engine  {
	r.POST("/set",controller.Setredis)
	r.POST("/get",controller.Getredis)

	return r
}
