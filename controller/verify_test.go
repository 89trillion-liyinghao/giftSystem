package controller

import (
	"fmt"
	rd "giftSystem/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type gift struct {
	Count   int `json:"count"`   //礼品数量 负数为无限领取
	Gold    int `json:"gold"`    //增加金币数量
	Diamond int `json:"diamond"` //增加钻石数量
	Prop    int `json:"prop"`    //增加道具数量
}

func init() {
	//初始化数据库配置
	err := rd.SetupSetting()
	if err != nil {
		//写日志
		fmt.Println("读取redis配置文件失败")
		return
	}

	//连接redis数据库
	err = rd.ConnRedis()
	if err != nil {
		//写日志
		fmt.Println("redis数据库连接失败")
		return
	}
}

/*
VerifyCode函数测试单元
访问localhost:8083/TestVerifyCode
参数
cookie：userName=121
query：code=8A8S2DZX

返回打印领取结果
*/

func TestVerifyCode(t *testing.T) {
	r := gin.Default()
	r.GET("/TestVerifyCode", VerifyCode1)
	_ = r.Run(":8083")
}

func VerifyCode1(c *gin.Context) {
	//检查用户是否登录
	uid, err := c.Cookie("userName")
	if err != nil || uid == "" {
		c.String(http.StatusOK, "请先登录")
		return
	}

	//礼品码判空
	//code打桩
	code := "XXXXXXXX"
	//code := c.DefaultQuery("code", "")
	//if code == "" {
	//	c.String(http.StatusOK, "礼品码不能为空")
	//	return
	//}

	//验证用户是否可以领取礼品码
	//var gift gift

	suc := "成功" //打桩测试
	//suc := "重复领取"
	//奖品信息
	if suc == "" {
		c.String(http.StatusOK, "礼品码错误或不存在")
		return
	} else if suc == "重复领取" {
		c.String(http.StatusOK, "不能重复领取")
		return
	}

	//调用奖励接口，增加奖励
	//result := AddGift1(uid, code, gift)
	//打桩测试
	result := "礼品内容XXX"

	fmt.Printf("用户\"" + uid + "\"礼品码：\"" + code + "\"获取成功\n")
	c.String(http.StatusOK, "获取成功，获得礼品为："+result)
}
