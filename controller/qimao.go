package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hoisie/httplib"
	"go-redis/model"
)

func Qimao(ctx *gin.Context){
	surl := "https://www.qimao.com/"
	rsp,_ :=httplib.Get(surl).AsString()
	var qimao model.Qimao
	str :=qimao.QmUrls(rsp)

	ctx.JSON(200,gin.H{"title":str})
}
func QimaoPage(ctx *gin.Context)  {
	url :=`https://www.qimao.com/shuku/143170/`
	rsp,_ :=httplib.Get(url).AsString()

	var qimao model.Qimao
	title :=qimao.Qmtitle(rsp)
	author := ""//qimao.Qmauthor(rsp)
	master := qimao.Qmmast(rsp)
	ctx.JSON(200,gin.H{"title":title,"Author":author,"主角":master})
}

func Qimaoget(ctx *gin.Context){
	url :=`https://www.qimao.com/shuku/143170/`
	rsp,_ :=httplib.Get(url).AsString()
	var qimao model.Qimao
	res := qimao.Qmmast(rsp)
	ctx.JSON(200,gin.H{"status":res})
}
