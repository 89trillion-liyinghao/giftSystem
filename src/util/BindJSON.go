package util

import (
	"giftSystem/src/entity"
	l "giftSystem/src/log"
	"github.com/gin-gonic/gin"
)

/*绑定礼品json数据*/
func Bind(c *gin.Context,json *entity.AddGift) error {
	err := c.ShouldBindJSON(json)
	if err != nil {
		l.Error.Printf("json绑定失败: %v\n", err)
		return err
	}
	return nil
}
