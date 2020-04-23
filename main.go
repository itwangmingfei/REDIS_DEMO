package main

import (
	"github.com/gin-gonic/gin"
	"go-redis/router"
	"go-redis/tools"
)

func main(){
	tools.InitRedis()
	r:= gin.Default()
	r.Static("/", "./public")

	router.Router(r)

	r.Run()
}
