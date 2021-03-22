package service

import (
	"encoding/json"
	rd "giftSystem/src/dao"
	"giftSystem/src/entity"
	l "giftSystem/src/log"
)

func AddGift(uid string,code string,gift entity.AddGift) string {

	//奖励逻辑



	//返回奖励结果
	result,err := json.Marshal(gift)
	if err != nil{
		return ""
	}

	//保存用户领取信息
	suc := rd.RedisStore{}.SetUid(code,uid)
	if !suc {
		//写日志
		l.Trace.Printf("礼品获取失败\n")
		return ""
	}
	return string(result)
}
