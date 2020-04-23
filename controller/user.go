package controller

import (
	"github.com/gin-gonic/gin"
	"go-redis/model"
	"go-redis/tools"
	"log"
)

func Setredis(ctx *gin.Context){
	var user model.User
	//获取数据
	ctx.ShouldBind(&user)
	//判断是否存在历史中
	name :=user.Name
	//验证历史库中是否存在
	res :=tools.Redis.ToIsset(name)
	log.Println(res)
	//存入redis false 可以存入
	if !res {
		tools.Redis.DoLpush(name)
		//获取需要执行的队列数量
		nums := tools.Redis.DoLen()
		ctx.JSON(200,gin.H{"status":true,"data":name,"nums":nums,"key":tools.DO_QUEUE})
		return
	}
	nums := tools.Redis.DoLen()
	ctx.JSON(200,gin.H{"status":true,"msg":"历史记录存在","nums":nums,"key":tools.DO_QUEUE})

	return

}
func Getredis(ctx *gin.Context){
	//取出数据
	str,_ :=tools.Redis.DoRpop()
	//当前队列还有几个数据
	nums := tools.Redis.DoLen()
	//取出的数据存入历史记录
	tools.Redis.ToSadd(str)

	ctx.JSON(200,gin.H{"status":true,"data":str,"nums":nums,"key":tools.TO_QUEUE})
}


