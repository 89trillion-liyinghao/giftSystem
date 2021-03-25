package setting

import (
	"fmt"
	"github.com/spf13/viper"
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
	vp1, _ := NewSetting_mock()

	fmt.Println("-------------")
	//测试调用ReadSection()读取指定内容
	vp1.ReadSection_mock("Redis", &redisS)
	fmt.Println("ReadSection_Redis_Addr:" + redisS.Addr)
	fmt.Println("ReadSection_Redis_Password:" + redisS.Password)

}
func NewSetting_mock() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")                                    //读取文件名
	vp.AddConfigPath("/Users/liyinghao/go/src/giftSystem/config") //配置文件路径
	vp.SetConfigType("yaml")                                      //配置文件后缀
	err := vp.ReadInConfig()
	if err != nil {
		//写日志
		fmt.Printf("调用NewSetting()读取配置文件失败: %v\n", err)
		return nil, err
	}

	s := &Setting{vp}
	return s, nil
}

//读取指定配置文件内容
func (s *Setting) ReadSection_mock(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		//写日志
		fmt.Printf("调用ReadSection()读取配置文件指定内容失败: %v\n", err)
		return err
	}

	return nil
}
