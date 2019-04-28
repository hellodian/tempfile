package tiger
import (
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/crypto/sha3"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdk"
	"fmt"
)


//ShuffleMainRound This is a sample method
func (t *Tiger) GetRandomNum(reveal []byte) (w bn.Number){
	w=bn.NBytes(sha3.Sum256(reveal, t.sdk.Block().BlockHash(), t.sdk.Block().RandomNumber()))
	return
}



func (t *Tiger) GetTransferData(settings *Settings, tokenName *string, transferAmount *bn.Number) {
	for _, name := range settings.TokenNames {
		transferReceipt := t.sdk.Message().GetTransferToMe(name)
		if transferReceipt != nil {
			*tokenName = name
			*transferAmount = transferReceipt.Value
			break
		}
	}

}


func (t *Tiger) LockedInBetsInit(tokenNameList []types.Address) {
	for _, value := range tokenNameList {
		t._setLockedInBets(value, bn.N(0))
	}
}


func (t *Tiger) createLimitMaps(tokenNameList []types.Address) map[string]Limit {
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




func (t *Tiger) checkSettings(newSettings *Settings) {

	sdk.Require(len(newSettings.TokenNames) > 0,
		types.ErrInvalidParameter, "tokenNames cannot be empty")

	for _, tokenName := range newSettings.TokenNames {
		token := t.sdk.Helper().TokenHelper().TokenOfName(tokenName)
		sdk.Require(token != nil,
			types.ErrInvalidParameter, fmt.Sprintf("tokenName=%s is not exist", tokenName))

		sdk.Require(newSettings.LimitMaps[tokenName].MaxLimit > 0,
			types.ErrInvalidParameter, "MaxBet must be bigger than zero")

		sdk.Require(newSettings.LimitMaps[tokenName].MaxProfit >= 0,
			types.ErrInvalidParameter, "MaxProfit can not be negative")

		sdk.Require(newSettings.LimitMaps[tokenName].MinLimit > 0 && newSettings.LimitMaps[tokenName].MinLimit < newSettings.LimitMaps[tokenName].MaxLimit,
			types.ErrInvalidParameter, "MinBet must be bigger than zero and smaller than MaxBet")
	}


	sdk.Require(newSettings.SendToCltRatio >= 0 && newSettings.SendToCltRatio < PERMILLE,
		types.ErrInvalidParameter,
		fmt.Sprintf("SendToCltRatio must be bigger than zero and smaller than %d", PERMILLE))
	if newSettings.FeeRatio==0{
		sdk.Require(newSettings.SendToCltRatio == 0,
			types.ErrInvalidParameter, "if SendToCltRatio is zero,SendToCltRatio must be zero")
	}

	sdk.Require(newSettings.FeeRatio > 0 && newSettings.FeeRatio < PERMILLE,
		types.ErrInvalidParameter,
		fmt.Sprintf("FeeRatio must be bigger than zero and  smaller than %d", PERMILLE))

	sdk.Require(newSettings.FeeMiniNum > 0,
		types.ErrInvalidParameter, "FeeMinimum must be bigger than zero")

	sdk.Require(newSettings.BetExpirationBlocks > 0,
		types.ErrInvalidParameter, "BetExpirationBlocks must be bigger than zero")
}

func (t *Tiger) checkRecFeeInfo(infos []RecFeeInfo) {
	sdk.Require(len(infos) > 0,
		types.ErrInvalidParameter, "The length of RecvFeeInfos must be larger than zero")

	allRatio := int64(0)
	for _, info := range infos {
		sdk.Require(info.RecFeeRatio > 0,
			types.ErrInvalidParameter, "ratio must be larger than zero")
		sdk.RequireAddress(t.sdk, info.RecFeeAddr)
		sdk.Require(info.RecFeeAddr != t.sdk.Message().Contract().Account(),
			types.ErrInvalidParameter, "address cannot be contract account address")

		allRatio += info.RecFeeRatio
	}

	//The allocation ratio set must add up to 1000
	sdk.Require(allRatio <= 1000,
		types.ErrInvalidParameter, "The sum of ratio must be less or equal 1000")
}


//Transfer to fee's receiving address
func (t *Tiger) transferToRecFeeAddr(tokenName string,fee bn.Number) {

	account := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Account())

	infos := t._recFeeInfo()
	for _, info := range infos {
		account.TransferByName(tokenName, info.RecFeeAddr, fee.MulI(info.RecFeeRatio).DivI(PERMILLE))
	}
}



//ShuffleMainCards 组建 免费游戏的牌面数据  和主游戏有区别 免费游戏最后一个是19个数
func (t *Tiger) ShuffleFeeCards(reveal [][]byte, playinfo *PlayerInfo) {
	w := make([]int64, 3)
	for i := 0; i < len(reveal); i++ {
		if i == len(reveal)-1 { //表示是最后一轮了
			w = t.ShuffleRound(reveal[i], t.pokerFeeSet[i][:19])
		} else {
			w = t.ShuffleRound(reveal[i], t.pokerFeeSet[i][:])
		}
		t.AssemblePoker(int64(i), w, playinfo)
	}

}
//ShuffleMainCards 组建 主游戏的牌面数据  为了在同一个区块内产生不同的随机数  传5个reveal 过来
func (t *Tiger) ShuffleMainCards(reveal [][]byte, playaddress types.Address, playInfo *PlayerInfo) {
	main := t._pokerMainSet()
	for i, v := range reveal {
		w := t.ShuffleRound(v, main[i][:])
		t.AssemblePoker(int64(i), w, playInfo)
	}

}

//ShuffleMainRound
func (t *Tiger) ShuffleRound(reveal []byte, s []int64) (w []int64) {
	w = make([]int64, 3)
	//取第一个数
	bytes := t.GetRandomNum(reveal)
	first := bytes.ModI(int64(len(s)))
	firstIndex := first.V.Int64()
	w[0] = s[firstIndex]
	newSlice := append(s[:firstIndex], s[firstIndex+1:]...)
	second := t.GetRandomNum(reveal).ModI(int64(len(newSlice)))
	secondIndex := second.V.Int64()
	//取第二个数
	w[1] = newSlice[secondIndex]
	newSlice = append(newSlice[:secondIndex], newSlice[secondIndex+1:]...)
	third := t.GetRandomNum(reveal).ModI(int64(len(newSlice)))
	thirdIndex := third.V.Int64()
	//取第三个数
	w[2] = newSlice[thirdIndex]
	return

}

//ShuffleMainRound 组建轮数据
func (t *Tiger) AssemblePoker(f int64, s []int64, playInfo *PlayerInfo) {
	playInfo.Poker[0][f] = Poker{0, f, s[0]}
	playInfo.Poker[1][f] = Poker{1, f, s[1]}
	playInfo.Poker[2][f] = Poker{2, f, s[2]}
	pokerlist:=[]Poker{playInfo.Poker[0][f],playInfo.Poker[1][f],playInfo.Poker[2][f]}

	//fire event
	t.emitAssemblePoker(f,pokerlist)

}

