package service

import (
	rd "giftSystem/src/dao"
	l "giftSystem/src/log"
	u "giftSystem/src/util"
)

//逻辑层：创建礼品码，成功返回礼品码，失败返回""
var giftId uint64 = 1

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

	suc := rd.RedisStore{}.Set(code, gift)
	if !suc {
		//写日志
		l.Trace.Printf("礼品创建失败\n")
		return ""
	}

	go l.Trace.Printf("礼品创建成功，礼品码：" + code + "\n")
	return code
}
