package service

import (
	"encoding/json"
	"giftSystem/entity"
	"giftSystem/logs"
	rd "giftSystem/repository"
	"sync"
)

var mutex sync.Mutex
var giftIsEmpty = false

/*
逻辑层：利用礼品码查询redis，返回查询结果
*/
func VerifyGift(code string, uid string, gift *entity.AddGift) string {
	//检查奖品码是否存在
	//ex := rd.RedisStore{}.Exist(code)
	//if !ex{
	//	l.Trace.Printf("礼品码：\""+code+"\"不存在\n")
	//	return ""
	//}

	//查询是否重复领取
	exi := rd.RedisStore{}.Exist(code, uid)
	if exi {
		return "重复领取"
	}

	mutex.Lock() //并发加锁

	//获取礼品
	giftStore := rd.RedisStore{}.Get(code, false)
	if giftStore == "" {
		go logs.Trace.Printf("礼品码：\"" + code + "\"不存在\n")
		mutex.Unlock()
		return ""
	}
	err := json.Unmarshal([]byte(giftStore), gift)
	if err != nil {
		logs.Error.Printf("json绑定失败: %v\n", err)
		mutex.Unlock()
		return ""
	}
	if gift.Count <= 0 {
		mutex.Unlock()
	}

	//内存标记判空
	if giftIsEmpty {
		return ""
	}

	if gift.Count == 0 {
		rd.RedisStore{}.Del(code)
		giftIsEmpty = true
		return ""
	}

	if gift.Count > 0 {
		gift.Count--
		giftJson, _ := json.Marshal(gift)
		rd.RedisStore{}.Set(code, string(giftJson))
		defer mutex.Unlock()
	}

	return giftStore

}
