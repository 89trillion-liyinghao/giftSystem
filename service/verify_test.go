package service

import (
	"encoding/json"
	"fmt"
	"testing"
)

type giftJSON struct {
	Count   int `json:"count"`   //礼品数量 负数为无限领取
	Gold    int `json:"gold"`    //增加金币数量
	Diamond int `json:"diamond"` //增加钻石数量
	Prop    int `json:"prop"`    //增加道具数量
}

var code = "SAS2DZX9"                                                   //设定礼品码
var user_id = "111"                                                     //设定用户id
var exi = false                                                         //是否重复领取
var giftStore = "{\"count\": 5,\"gold\": 5,\"diamond\": 5,\"prop\": 5}" //设定礼品json字符串

/*
Verify函数单元测试
打印返回结果，成功返回礼品内容
*/
func TestVerify(t *testing.T) {
	/*设定礼品内容*/
	var g giftJSON
	g.Count = 5
	g.Gold = 5
	g.Diamond = 5
	g.Prop = 5

	res := VerifyGift1(code, user_id, &g)

	fmt.Println(res)
}

func VerifyGift1(code string, uid string, gift *giftJSON) string {
	//检查奖品码是否存在
	//ex := rd.RedisStore{}.Exist(code)
	//if !ex{
	//	l.Trace.Printf("礼品码：\""+code+"\"不存在\n")
	//	return ""
	//}

	//获取礼品
	//giftStore := rd.RedisStore{}.Get(code, false)
	if giftStore == "" {
		fmt.Printf("礼品码：\"" + code + "\"不存在\n")
		return ""
	}
	err := json.Unmarshal([]byte(giftStore), gift)
	if err != nil {
		fmt.Printf("json绑定失败: %v\n", err)
		return ""
	}

	//查询是否重复领取
	//exi := rd.RedisStore{}.Exist(code, uid)
	if exi {
		return "重复领取"
	}

	if gift.Count > 0 {
		gift.Count--
		//giftJson, _ := json.Marshal(gift)
		//rd.RedisStore{}.Set(code, string(giftJson))
	} else if gift.Count == 0 {
		//rd.RedisStore{}.Del(code)
		return ""
	}

	return giftStore

}
