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
func TestNewSetting(t *testing.T)  {
	type RedisSettingS struct {
		Addr string
		Password string
	}

	var redisS RedisSettingS

	vp, _ := NewSetting()
	vp.ReadSection("Redis",&redisS)

	fmt.Println("ReadSection_Redis_Addr:"+redisS.Addr)
	fmt.Println("ReadSection_Redis_Password:"+redisS.Password)

}


type Setting struct {
	vp *viper.Viper
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