package dao_test

import (
	"fmt"
	"github.com/spf13/viper"
)

/*
dao层：读取指定数据库的配置文件
*/

//redis配置
type RedisSettingS struct {
	Addr string
	Password string
}

type Setting struct {
	vp *viper.Viper
}

//定义全局变量
var RedisSetting *RedisSettingS

//读取配置文件到全局变量
func SetupSetting() error {
	s , err := NewSetting()
	if err != nil{
		return err
	}

	err = s.ReadSection("Redis",&RedisSetting)
	if err != nil{
		return err
	}

	return nil
}

func NewSetting() (*Setting,error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("/Users/liyinghao/go/src/giftSystem/src/config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		//写日志
		fmt.Printf("调用NewSetting()读取配置文件失败: %v\n", err)
		return nil, err
	}

	s := &Setting{vp}

	fmt.Print("NewSetting_Redis.Addr:")
	fmt.Println(vp.Get("Redis.Addr"))
	fmt.Print("NewSetting_Redis.PassWord:")
	fmt.Println(vp.Get("Redis.PassWord"))
	fmt.Println("---------------------")

	return s,nil
}

//读取指定配置文件内容
func (s *Setting) ReadSection(k string,v interface{}) error {
	err := s.vp.UnmarshalKey(k,v)
	if err != nil{
		//写日志
		fmt.Printf("调用ReadSection()读取配置文件指定内容失败: %v\n", err)
		return err
	}
	return nil
}