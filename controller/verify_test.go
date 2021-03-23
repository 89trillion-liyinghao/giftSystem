package controller

import (
	"encoding/json"
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
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.String(http.StatusOK, "礼品码不能为空")
		return
	}

	//验证用户是否可以领取礼品码
	var gift gift
	suc := VerifyGift1(code, uid, &gift)
	if suc == "" {
		c.String(http.StatusOK, "礼品码错误或不存在")
		return
	} else if suc == "重复领取" {
		c.String(http.StatusOK, "不能重复领取")
		return
	}

	//调用奖励接口，增加奖励
	result := AddGift1(uid, code, gift)

	fmt.Printf("用户\"" + uid + "\"礼品码：\"" + code + "\"获取成功\n")
	c.String(http.StatusOK, "获取成功，获得礼品为："+result)
}

func VerifyGift1(code string, uid string, gift *gift) string {
	//检查奖品码是否存在
	//ex := rd.RedisStore{}.Exist(code)
	//if !ex{
	//	l.Trace.Printf("礼品码：\""+code+"\"不存在\n")
	//	return ""
	//}

	//获取礼品
	giftStore := rd.RedisStore{}.Get(code, false)
	if giftStore == "" {
		fmt.Printf("礼品码：\"" + code + "\"不存在\n")
		return ""
	}
	err := json.Unmarshal([]byte(giftStore), gift)
	if err != nil {
		fmt.Printf("json绑定失败: %v\n", err)
		return ""
	}

	//查询是否重复领取
	exi := rd.RedisStore{}.Exist(code, uid)
	if exi {
		return "重复领取"
	}

	if gift.Count > 0 {
		gift.Count--
		giftJson, _ := json.Marshal(gift)
		rd.RedisStore{}.Set(code, string(giftJson))
	} else if gift.Count == 0 {
		rd.RedisStore{}.Del(code)
		return ""
	}

	return giftStore

}

func AddGift1(uid string, code string, gift gift) string {

	//奖励逻辑
	//fmt.Println("执行增加奖励逻辑")

	//返回奖励结果
	result, err := json.Marshal(gift)
	if err != nil {
		return ""
	}

	//保存用户领取信息
	suc := rd.RedisStore{}.SetUid(code, uid)
	if !suc {
		//写日志
		fmt.Printf("礼品获取失败\n")
		return ""
	}
	return string(result)
}
