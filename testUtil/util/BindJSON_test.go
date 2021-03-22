package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)
type gift struct {
	Count   int `json:"count"`       //礼品数量 负数为无限领取
	Gold    int `json:"gold"`        //增加金币数量
	Diamond int `json:"diamond"`     //增加钻石数量
	Prop    int `json:"prop"`        //增加道具数量
}


/*
Bind绑定JSON字符串测试
访问/TestBind:8080地址，传递json，控制台打印Json字符串
*/

func TestBind(t *testing.T)  {

	r := gin.Default()
	r.POST("/TestBind",Bind01)
	_ = r.Run(":8080")
}

func Bind01(c *gin.Context)  {
	var Json gift
	Bind(c,&Json)
	result,_ := json.Marshal(Json)
	fmt.Println(string(result))
}

func Bind(c *gin.Context,json *gift) error {
	err := c.ShouldBindJSON(json)
	if err != nil {
		fmt.Println("json绑定失败")
		return err
	}
	return nil
}
