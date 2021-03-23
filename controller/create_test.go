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
