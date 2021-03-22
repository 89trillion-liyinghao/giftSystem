package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

/*
NewSetting()初始化viper，定位配置文件
ReadSection()读取配置文件指定位置内容
*/

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting,error) {
	vp := viper.New()
	vp.SetConfigName("config")         //读取文件名
	vp.AddConfigPath("src/config")     //配置文件路径
	vp.SetConfigType("yaml")           //配置文件后缀
	err := vp.ReadInConfig()
	if err != nil {
		//写日志
		fmt.Printf("调用NewSetting()读取配置文件失败: %v\n", err)
		return nil, err
	}

	s := &Setting{vp}
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