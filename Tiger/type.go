package tiger

import "blockchain/smcsdk/sdk/types"

type Poker struct {
	X int64   //坐标
	Y int64  //坐标
	Value int64   //牌的值
}
//不同牌对应的值
//W=0
//K=4
//Q=3
//J=2
//S=1


type PlayerInfo struct {
	Address types.Address  //地址
	Poker    [3][5]Poker
	Fee      *FeeInfo   //免费游戏信息
	Currency  map[string]int64 //当前账户 币种的余额
}
type FeeInfo struct {
	FeeCount   int64
	BetAmout   int64
	TokenName  string
}

// Settings - contract settings
type Settings struct {
	LimitMaps           map[string]Limit `json:"limitMaps"`           // Maximum winning amount、Maximum bet、Minimum bet limit (cong)
	FeeRatio            int64    `json:"feeRatio"`            // Percentage of handling fee after winning the lottery (thousand-point ratio)
	FeeMiniNum          int64    `json:"feeMiniNum"`          // Minimum handling charge
	SendToCltRatio      int64    `json:"sendToCltRatio"`      // Part of the handling fee sent to CLT (thousandths)
	BetExpirationBlocks int64    `json:"betExpirationBlocks"` // Timeout block interval
	TokenNames          []string `json:"tokenNames"`          // List of supported token names
}



type RecFeeInfo struct {
	RecFeeRatio int64         `json:"recFeeRatio"` // Commission allocation ratio
	RecFeeAddr  types.Address `json:"recFeeAddr"`  // List of addresses to receive commissions
}

type Limit struct {
	MaxProfit int64 `json:"maxProfit"` // Maximum winning amount (cong)
	MaxLimit  int64 `json:"maxLimit"`  // Maximum bet limit (cong)
	MinLimit  int64 `json:"minLimit"`  // Minimum bet limit unit (cong)
}



