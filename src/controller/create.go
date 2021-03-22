package controller

import (
	"encoding/json"
	"giftSystem/src/entity"
	l "giftSystem/src/log"
	. "giftSystem/src/service"
	"giftSystem/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)


/*
应用层：定义利用礼品码以及礼品创建缓存接口，返回创建信息
*/
func CreatCode(c *gin.Context)  {
	//gift := c.DefaultQuery("gift","")
	//if gift == ""{
	//	c.String(http.StatusOK,"礼品内容不能为空")
	//	return
	//}

	var giftJSON entity.AddGift
	err := util.Bind(c,&giftJSON)
	if err != nil{
		c.String(http.StatusOK,"服务器错误，请重试")
		return
	}

	//获取json字符串
	gift,err := json.Marshal(giftJSON)
	if err != nil{
		l.Error.Printf("json字符串生成失败: %v\n", err)
		c.String(http.StatusOK,"服务器错误，请重试")
		return
	}

	res := CreateGift(string(gift))

	if res != ""{
		c.String(http.StatusOK,"礼品码创建成功,礼品码为："+res)
	} else{
		c.String(http.StatusOK,"奖品码创建失败，请重试")
	}
	return
}
