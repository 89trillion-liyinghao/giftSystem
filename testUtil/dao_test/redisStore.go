package dao_test

import (
	"context"
	"fmt"
	"time"
)

/*
dao层：封装redis方法
*/

var ctx = context.Background()

//定义缓存头
const GIFTCODE = "test_giftcode:"

type RedisStore struct {

}

/*封装redis写入方法，缓存时间5min*/
func (r RedisStore) Set(code string,gift string) bool {
	key := GIFTCODE + code
	err := TestRedis.Set(ctx,key,gift,time.Minute*1200).Err()
	if err != nil{
		return false
	}
	return true
}

/*封装redis获取缓存，判断获取后是否删除*/
func (r RedisStore) Get(code string, clear bool) string {
	key := GIFTCODE + code
	val,err := TestRedis.Get(ctx,key).Result()
	if err != nil{
		return ""
	}

	if clear{
		err := TestRedis.Del(ctx,key).Err()
		if err != nil{
			//写日
			fmt.Println("礼品码："+code+"删除失败，" +"%v\n", err)
			return ""
		}
		fmt.Println("礼品领取完毕，删除礼品码："+code+"\n")
		return ""
	}
	return val
}

/*查询key是否存在*/
func (r RedisStore) Exist(code string,uid string) bool {
	key := code + uid
	_,err := TestRedis.Get(ctx,key).Result()
	if err != nil{
		return false
	}
	return true
}

func (r RedisStore) Del(code string)  {
	key := GIFTCODE + code
	err := TestRedis.Del(ctx,key).Err()
	if err != nil{
		//写日
		fmt.Println("礼品码："+code+"删除失败，" +"%v\n", err)
		return
	}
	fmt.Println("礼品领取完毕，删除礼品码："+code+"\n")
}

func (r RedisStore) SetUid(code string,uid string) bool {
	key := code + uid
	err := TestRedis.Set(ctx,key,"Verified",time.Minute*1200).Err()
	if err != nil{
		return false
	}
	return true
}
