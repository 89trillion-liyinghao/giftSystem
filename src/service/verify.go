package service

import (
	"encoding/json"
	rd "giftSystem/src/dao"
	"giftSystem/src/entity"
	l "giftSystem/src/log"
)


/*
逻辑层：利用礼品码查询redis，返回查询结果
*/
func VerifyGift(code string,uid string,gift *entity.AddGift) string {
	//检查奖品码是否存在
	//ex := rd.RedisStore{}.Exist(code)
	//if !ex{
	//	l.Trace.Printf("礼品码：\""+code+"\"不存在\n")
	//	return ""
	//}

	//获取礼品
	giftStore := rd.RedisStore{}.Get(code,false)
	if giftStore == ""{
		l.Trace.Printf("礼品码：\""+code+"\"不存在\n")
		return ""
	}
	err := json.Unmarshal([]byte(giftStore),gift)
	if err != nil{
		l.Error.Printf("json绑定失败: %v\n", err)
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
