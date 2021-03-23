package router

import (
	. "giftSystem/controller"
	"github.com/gin-gonic/gin"
)

/*
初始化路由
*/
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/createCode", CreatCode)
	router.GET("/giftVerify", VerifyCode)
	return router
}
