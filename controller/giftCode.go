package controller

import (
	"encoding/json"
	"giftSystem/entity"
	"giftSystem/logs"
	"giftSystem/pkg/util"
	. "giftSystem/service"
	sv "giftSystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
应用层：定义利用礼品码以及礼品创建缓存接口，返回创建信息
*/
func CreatCode(c *gin.Context) {
	//gift := c.DefaultQuery("gift","")
	//if gift == ""{
	//	c.String(http.StatusOK,"礼品内容不能为空")
	//	return
	//}

	var giftJSON entity.AddGift
	err := util.Bind(c, &giftJSON)
	if err != nil {
		c.String(http.StatusOK, "服务器错误，请重试")
		return
	}

	//获取json字符串
	gift, err := json.Marshal(giftJSON)
	if err != nil {
		logs.Error.Printf("json字符串生成失败: %v\n", err)
		c.String(http.StatusOK, "服务器错误，请重试")
		return
	}

	res := CreateGift(string(gift))

	if res != "" {
		c.String(http.StatusOK, "礼品码创建成功,礼品码为："+res)
	} else {
		c.String(http.StatusOK, "奖品码创建失败，请重试")
	}
	return
}

/*
应用层：定义利用礼品码获取礼品接口，成功返回礼品
*/
func VerifyCode(c *gin.Context) {
	//检查用户是否登录
	uid, err := c.Cookie("userName")
	if err != nil || uid == "" {
		c.String(http.StatusOK, "请先登录")
		return
	}

	//测试用id设置
	//uid := "000000"

	//礼品码判空
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.String(http.StatusOK, "礼品码不能为空")
		return
	}

	//验证用户是否可以领取礼品码
	var gift entity.AddGift
	suc := sv.VerifyGift(code, uid, &gift)
	if suc == "" {
		c.String(http.StatusOK, "礼品码错误或不存在")
		return
	} else if suc == "重复领取" {
		c.String(http.StatusOK, "不能重复领取")
		return
	}

	//调用奖励接口，增加奖励
	result := sv.AddGift(uid, code, gift)

	go logs.Trace.Printf("用户\"" + uid + "\"礼品码：\"" + code + "\"获取成功\n")
	c.String(http.StatusOK, "获取成功，获得礼品为："+result)
}
