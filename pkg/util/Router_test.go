package util

import (
	"fmt"
	"testing"
)

/*
Encode函数测试
测试uid = 1-10代码输出，预期结果为不同的8位字符串
*/
func TestEncode(t *testing.T) {

	for uid := 1; uid <= 10; uid++ {
		res := Encode(uint64(uid))
		fmt.Println(res)
	}
}
