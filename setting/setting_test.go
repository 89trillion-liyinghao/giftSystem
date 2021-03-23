package setting

import (
	"fmt"
	"testing"
)

/*
1.NewSetting函数测试单元
输出配置文件所有信息

2.ReadSection函数测试单元
输出Reids配置文件所有信息
*/
func TestNewSetting(t *testing.T) {
	type RedisSettingS struct {
		Addr     string
		Password string
	}

	var redisS RedisSettingS

	//测试调用NewSetting()函数
	vp1, _ := NewSetting()

	fmt.Println("-------------")
	//测试调用ReadSection()读取指定内容
	vp1.ReadSection("Redis", &redisS)
	fmt.Println("ReadSection_Redis_Addr:" + redisS.Addr)
	fmt.Println("ReadSection_Redis_Password:" + redisS.Password)

}
