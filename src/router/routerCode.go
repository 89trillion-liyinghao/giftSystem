package router

import (
	"github.com/gin-gonic/gin"
	. "giftSystem/src/controller"
)


/*
初始化路由
*/
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/createCode",CreatCode)
	router.GET("/giftVerify",VerifyCode)
	return router
}
