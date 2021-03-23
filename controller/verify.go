package controller

import (
	"giftSystem/entity"
	"giftSystem/logs"
	sv "giftSystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
