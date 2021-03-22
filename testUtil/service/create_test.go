package service

import (
	"encoding/json"
	"fmt"
	u "giftSystem/src/util"
	rd "giftSystem/testUtil/dao_test"
	"testing"
)

var giftId uint64 = 1



/*
Create函数测试类
设定礼品内容，打印8位礼品码字符串
*/
func TestCreate(t *testing.T)  {

	type giftJSON struct {
		Count   int `json:"count"`       //礼品数量 负数为无限领取
		Gold    int `json:"gold"`        //增加金币数量
		Diamond int `json:"diamond"`     //增加钻石数量
		Prop    int `json:"prop"`        //增加道具数量
	}


	/*设定礼品内容*/
	var g giftJSON
	g.Count = 5
	g.Gold = 5
	g.Diamond = 5
	g.Prop = 5

	gi,_ := json.Marshal(g)

	res := CreateGift(string(gi))
	fmt.Println(res)
}


func CreateGift(gift string) string {
	code := u.Encode(giftId)
	giftId++

	/*
		//检查奖品码是否存在
		giftStore := rd.RedisStore{}.Get(code,false)
		if giftStore != ""{
			//写日志
			return ""
		}
	*/

	suc := rd.RedisStore{}.Set(code,gift)
	if !suc {
		//写日志
		fmt.Printf("礼品创建失败\n")
		return ""
	}

	fmt.Printf("礼品创建成功，礼品码："+code+"\n")
	return code
}
