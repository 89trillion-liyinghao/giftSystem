package service

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testGift struct {
	Count   int `json:"count"`   //礼品数量 负数为无限领取
	Gold    int `json:"gold"`    //增加金币数量
	Diamond int `json:"diamond"` //增加钻石数量
	Prop    int `json:"prop"`    //增加道具数量
}

var id = "0"       //设定用户id
var c = "3r3r3223" //设定礼品码
var suc = true     //设定增加礼品是否成功

/*测试AddGift函数,返回增加礼品结果*/
func TestAddGift(t *testing.T) {

	var gift testGift //设定礼品内容
	gift.Count = 5
	gift.Gold = 5
	gift.Diamond = 5
	gift.Prop = 5

	res := AddGift1(id, c, gift)

	if res != "" {
		fmt.Println("用户" + id + "获取礼品码：" + c + "，礼品内容：" + res + "成功")
	}
}

func AddGift1(uid string, code string, gift testGift) string {

	//奖励逻辑
	fmt.Println("执行增加奖励逻辑")

	//返回奖励结果
	result, err := json.Marshal(gift)
	if err != nil {
		return ""
	}

	//保存用户领取信息
	//suc := rd.RedisStore{}.SetUid(code, uid)
	if !suc {
		//写日志
		fmt.Printf("礼品获取失败\n")
		return ""
	}
	return string(result)
}
