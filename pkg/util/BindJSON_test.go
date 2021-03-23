package util

import (
	"encoding/json"
	"fmt"
	"giftSystem/entity"
	"github.com/gin-gonic/gin"
	"testing"
)

/*
Bind绑定JSON字符串测试
访问/TestBind:8081地址，传递json，控制台打印Json字符串
*/
func TestBind(t *testing.T) {

	r := gin.Default()
	r.POST("/TestBind", Bind01)
	_ = r.Run(":8081")
}

func Bind01(c *gin.Context) {
	var Json entity.AddGift
	Bind(c, &Json)
	result, _ := json.Marshal(Json)
	fmt.Println(string(result))
}
