package sicbo

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
)

//SicBo This is struct of contract
//@:contract:sicbo
//@:version:1.0
//@:organization:orgBtjfCSPCAJ84uQWcpNr74NLMWYm5SXzer
//@:author:01bd6c29d63f5f32aa33955f26a28459988edea4de517f77372e77db33958e6e
type SicBo struct {
	sdk sdk.ISmartContract

	//@:public:store:cache
	secretSigner types.PubKey // Check to sign the public key

	//@:public:store:cache
	lockedInBets map[string]bn.Number // Lock amount (unit cong) key: currency name

	//@:public:store:cache
	settings *SBSettings

	//@:public:store:cache
	recFeeInfo []RecFeeInfo

	//@:public:store
	roundInfo map[string]*RoundInfo
	//@:public:store
	betInfo map[string]map[string]*BetInfo

	//@:public:store
	bigOrSmall int64
}

func (sb *SicBo) LockedInBetsInit(tokenNameList []types.Address) {
	for _, value := range tokenNameList {
		sb._setLockedInBets(value, bn.N(0))
	}
}

func (sb *SicBo) createLimitMaps(tokenNameList []types.Address) map[string]Limit {
	limitMaps := make(map[string]Limit, len(tokenNameList))
	for _, value := range tokenNameList {
		limit := Limit{
			MaxProfit :2E12,
			MaxLimit : 2E10,
			MinLimit : 1E8,
		}
		limitMaps[value] = limit
	}
	return limitMaps
}



type SBSettings struct {
	LimitMaps           map[string]Limit `json:"limitMaps"`           // Maximum winning amount、Maximum bet、Minimum bet limit (cong)
	FeeRatio            int64            `json:"feeRatio"`            // Percentage of handling fee after winning the lottery (thousand-point ratio)
	FeeMiniNum          int64            `json:"feeMiniNum"`          // Minimum handling charge
	SendToCltRatio      int64            `json:"sendToCltRatio"`      // Part of the handling fee sent to CLT (thousandths)
	BetExpirationBlocks int64            `json:"betExpirationBlocks"` // Timeout block interval
	TokenNames          []string         `json:"tokenNames"`          // List of supported token names
}

type RecFeeInfo struct {
	RecFeeRatio int64         `json:"recFeeRatio"` // Commission allocation ratio
	RecFeeAddr  types.Address `json:"recFeeAddr"`  // List of addresses to receive commissions
}

//type RoundInfo struct {
//	Commit              []byte               `json:"commit"`              // Current game random number hash value
//	TotalBuyAmount      map[string]bn.Number `json:"totalBuyAmount"`      // Current total bet amount map key：tokenName(The name of the currency)
//	TotalBetCount       int64                `json:"totalBetCount"`       // Current total number of bets
//	State               int64                `json:"state"`               // Current wheel state 0 Not the lottery 1Has the lottery 2 refunded 3In the lottery
//	ProcessCount        int64                `json:"processCount"`        // Current status processing bet quantity (settlement, refund subscript)
//	FirstBlockHeight    int64                `json:"firstBlockHeight"`    // Block height when the current wheel initializes to determine whether to timeout
//	Settings            *SBSettings          `json:"settings"`            // Configuration information for the current wheel
//	BetInfoSerialNumber []types.Address      `json:"betInfoSerialNumber"` // The current wheel betInfo is associated with the serial number
//	Sic                 *Sic                 `json:"sic"`                 // The current round of lottery results
//}
type RoundInfo struct {
	Commit              []byte               `json:"commit"`              // Current game random number hash value
	TotalBuyAmount      map[string]bn.Number `json:"totalBuyAmount"`      // Current total bet amount map key：tokenName(The name of the currency)
	TotalBetCount       int64                `json:"totalBetCount"`       // Current total number of bets
	State               int64                `json:"state"`               // Current wheel state 0 Not the lottery 1Has the lottery 2 refunded 3In the lottery
	ProcessCount        int64                `json:"processCount"`        // Current status processing bet quantity (settlement, refund subscript)
	FirstBlockHeight    int64                `json:"firstBlockHeight"`    // Block height when the current wheel initializes to determine whether to timeout
	Settings            *SBSettings          `json:"settings"`            // Configuration information for the current wheel
	BetInfoSerialNumber []string      		 `json:"betInfoSerialNumber"` // The current wheel betInfo is associated with the serial number
	Sic                 *Sic                 `json:"sic"`                 // The current round of lottery results
}


type BetInfo struct {
	TokenName string        `json:"tokenName"` // Players bet on currency names
	Gambler   types.Address `json:"gambler"`   // Player betting address
	Amount    bn.Number     `json:"amount"`    // Players bet the total amount
	BetData   []BetData     `json:"betData"`   // Player betting details
	WinAmount bn.Number     `json:"winAmount"` // Players this bet the largest bonus
	Settled   bool          `json:"settled"`   // Whether the current bet has been settled
}

type BetData struct {
	BetMode   int64     `json:"betMode"`   // Betting plan
	BetAmount bn.Number `json:"betAmount"` // Betting amount
}

type Limit struct {
	MaxProfit int64 `json:"maxProfit"` // Maximum winning amount (cong)
	MaxLimit  int64 `json:"maxLimit"`  // Maximum bet limit (cong)
	MinLimit  int64 `json:"minLimit"`  // Minimum bet limit unit (cong)
}
