package service

import (
	"encoding/json"
	"fmt"
	rd "giftSystem/testUtil/dao_test"
	"testing"
)

type giftJSON struct {
	Count   int `json:"count"`       //礼品数量 负数为无限领取
	Gold    int `json:"gold"`        //增加金币数量
	Diamond int `json:"diamond"`     //增加钻石数量
	Prop    int `json:"prop"`        //增加道具数量
}
var code = "0"           //设定礼品码
var user_id = "111"      //设定用户id


/*
Verify函数单元测试
打印返回结果，成功返回礼品内容
*/
func TestVerify(t *testing.T)  {
	/*设定礼品内容*/
	var g giftJSON

	res := VerifyGift(code,user_id,&g)

	fmt.Println(res)
}

func VerifyGift(code string,uid string,gift *giftJSON) string {
	//检查奖品码是否存在
	//ex := rd.RedisStore{}.Exist(code)
	//if !ex{
	//	l.Trace.Printf("礼品码：\""+code+"\"不存在\n")
	//	return ""
	//}

	//获取礼品
	giftStore := rd.RedisStore{}.Get(code,false)
	if giftStore == ""{
		fmt.Printf("礼品码：\""+code+"\"不存在\n")
		return ""
	}
	err := json.Unmarshal([]byte(giftStore),gift)
	if err != nil{
		fmt.Printf("json绑定失败: %v\n", err)
		return ""
	}

	//查询是否重复领取
	exi := rd.RedisStore{}.Exist(code,uid)
	if exi{
		return "重复领取"
	}

	if gift.Count > 0{
		gift.Count--
		giftJson,_ := json.Marshal(gift)
		rd.RedisStore{}.Set(code,string(giftJson))
	}else if gift.Count == 0 {
		rd.RedisStore{}.Del(code)
		return ""
	}

	return giftStore

}
