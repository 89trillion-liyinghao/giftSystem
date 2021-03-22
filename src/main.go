package main

import (
	"giftSystem/src/router"
	"giftSystem/src/dao"
	l "giftSystem/src/log"
	"log"
)

//初始化读取配置文件
func init()  {
	//初始化日志系统
	err := l.InitLog()
	if err != nil{
		//写日志
		log.Fatalf("日志系统初始化失败: %v\n", err)
	}

	//初始化数据库配置
	err = dao.SetupSetting()
	if err != nil{
		//写日志
		l.Error.Fatalf("读取redis配置文件失败: %v\n", err)
	}

	//连接redis数据库
	err = dao.ConnRedis()
	if err != nil{
		//写日志
		l.Error.Fatalf("redis数据库连接失败:  %v\n",err)
	}
}

//启动服务
func main()  {
	router := router.InitRouter()
	router.Run(":8080")
}
