package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
)


/*
dao层：利用配置文件建立数据库连接
*/
//定义redis变量
var Redis *redis.Client

//建立redis链接
func ConnRedis() error {
	var ctx = context.Background()
	Redis = redis.NewClient(&redis.Options{
		Addr: RedisSetting.Addr,
		Password: RedisSetting.Password,
		DB: 0,
	})

	_,err := Redis.Ping(ctx).Result()
	if err != nil{
		return err
	}
	return nil
}
