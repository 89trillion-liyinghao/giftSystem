package entity

type AddGift struct {
	Count   int `json:"count"`       //礼品数量 负数为无限领取
	Gold    int `json:"gold"`        //增加金币数量
	Diamond int `json:"diamond"`     //增加钻石数量
	Prop    int `json:"prop"`        //增加道具数量
}

