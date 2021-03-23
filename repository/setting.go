package repository

import (
	"giftSystem/setting"
)

/*
dao层：读取指定数据库的配置文件
*/

//redis配置
type RedisSettingS struct {
	Addr     string
	Password string
}

//定义全局变量
var RedisSetting *RedisSettingS

//读取配置文件到全局变量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}

	return nil
}
