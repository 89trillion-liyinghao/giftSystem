package controller_test

import (
	"encoding/json"
	"fmt"
	u "giftSystem/pkg/util"
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

/*
Create函数测试单元
访问localhost:8080/TestCreatCode
发送post请求
json：
{
	"count": 5,
	"gold": 5,
	"diamond": 5,
	"prop": 5
}
打印返回创建的礼品码
*/
func TestCreate(t *testing.T) {
	r := gin.Default()
	r.POST("/TestCreatCode", CreatCode)
	_ = r.Run(":8082")
}

func CreatCode(c *gin.Context) {
	//gift := c.DefaultQuery("gift","")
	//if gift == ""{
	//	c.String(http.StatusOK,"礼品内容不能为空")
	//	return
	//}

	var giftJSON gift
	err := Bind(c, &giftJSON)
	if err != nil {
		c.String(http.StatusOK, "服务器错误，请重试")
		return
	}

	//获取json字符串
	gift, err := json.Marshal(giftJSON)
	if err != nil {
		fmt.Printf("json字符串生成失败: %v\n", err)
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

func Bind(c *gin.Context, json *gift) error {
	err := c.ShouldBindJSON(json)
	if err != nil {
		fmt.Println("json绑定失败")
		return err
	}
	return nil
}

var giftId uint64 = 1

func CreateGift(gift string) string {
	code := u.Encode(giftId)
	giftId++

	/*
		//检查奖品码是否存在
		giftStore := rd.RedisStore{}.Get(code,false)
		if giftStore != ""{
			//写日志
			return ""
		}
	*/

	//suc := rd.RedisStore{}.Set(code,gift)
	//if !suc {
	//	//写日志
	//	fmt.Printf("礼品创建失败\n")
	//	return ""
	//}

	fmt.Printf("礼品创建成功，礼品码：" + code + "\n")
	return code
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
