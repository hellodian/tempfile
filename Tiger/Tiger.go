package tiger

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/crypto/ed25519"
	_ "blockchain/smcsdk/sdk/crypto/sha3"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/types"
)

//@:public:receipt
type receipt interface {
	emitSetSecretSigner(newSecretSigner types.PubKey)
	emitSetOwner(newAddress types.Address)
	emitSetSettings(tokenNames []string, limit map[string]Limit, feeRatio, feeMiniNum, sendToCltRatio, betExpirationBlocks int64)
	emitSetRecFeeInfo(info []RecFeeInfo)
	emitAssemblePoker(num int64,listPoker []Poker)
	emitDigtalCurrency(token string,num int64)
	emitWithdrawFunds(tokenName string, beneficiary types.Address, withdrawAmount bn.Number)
	emitPlaceBet(tokenName string, pokerFix [][]Poker, mul, times, totaltimes int64, sum int64, gambler types.Address, WinAmount bn.Number, commitLastBlock int64, commit, signData []byte, refAddress types.Address)
}

//Tiger This is struct of contract
//@:contract:Tiger
//@:version:1.0
//@:organization:orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111
//@:author:cc196aa70a386ca672aed1062003d3c07cd6fc3f6b3b0ca71e7095e33a5c73b0
type Tiger struct {
	sdk sdk.ISmartContract

	//@:public:store:cache
	secretSigner types.PubKey // Check to sign the public key

	//@:public:store:cache
	lockedInBets map[string]bn.Number // Lock amount (unit cong) key: currency name

	//@:public:store:cache
	settings *Settings

	//@:public:store:cache
	recFeeInfo []RecFeeInfo

	//@:public:store
	pokerMainSet [5][20]int64

	//@:public:store
	pokerFeeSet [5][20]int64

	//@:public:store
	betInfo map[types.Address]*PlayerInfo
}


//InitChain - InitChain Constructor of this Tiger
//@:constructor
func (t *Tiger) InitChain() {
	// init data
	settings := Settings{}

	settings.FeeRatio = 50
	settings.FeeMiniNum = 300000
	settings.SendToCltRatio = 100
	settings.BetExpirationBlocks = 250
	settings.TokenNames = []string{t.sdk.Helper().GenesisHelper().Token().Name()}
	//settings.TokenNames = []string{"BCB", "DC", "XT", "USDX"}

	limitMaps := t.createLimitMaps(settings.TokenNames)

	settings.LimitMaps = limitMaps
	t._setSettings(&settings)
	t.LockedInBetsInit(settings.TokenNames)

	//	//主牌
	main := [5][20]int64{}
	fee := [5][20]int64{}
	main[0] = [20]int64{4, 4, 3, 3, 3, 2, 2, 2, 2, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 1}
	main[1] = [20]int64{0, 0, 0, 4, 4, 3, 3, 2, 2, 2, 2, 2, 10, 10, 10, 9, 9, 9, 9, 1}
	main[2] = [20]int64{0, 0, 0, 4, 4, 4, 3, 3, 2, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 1}
	main[3] = [20]int64{0, 0, 0, 0, 4, 4, 4, 3, 3, 3, 2, 10, 10, 10, 10, 9, 9, 9, 9, 1}
	main[4] = [20]int64{0, 0, 0, 0, 4, 4, 3, 3, 3, 3, 3, 2, 2, 10, 9, 9, 9, 9, 9, 1}
	//免费的牌的设置
	fee[0] = [20]int64{4, 3, 3, 2, 2, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 9, 9, 1}
	fee[1] = [20]int64{0, 4, 4, 4, 3, 3, 2, 2, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 9, 1}
	fee[2] = [20]int64{4, 4, 4, 3, 3, 3, 3, 2, 2, 10, 10, 10, 10, 10, 9, 9, 9, 9, 9, 1}
	fee[3] = [20]int64{4, 4, 3, 3, 3, 2, 2, 2, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 9, 1}
	fee[4] = [20]int64{4, 4, 3, 3, 3, 2, 2, 2, 10, 10, 10, 10, 9, 9, 9, 9, 9, 9, 1}

	t._setPokerMainSet(main)
	t._setPokerFeeSet(fee)
}

//SetPoker - set Tiger poker
//@:public:method:gas[500]
func (t *Tiger) SetPoker(main, fee string) {
	newmain := new([5][20]int64)
	newfee := new([5][20]int64)
	err := jsoniter.Unmarshal([]byte(main), newmain)
	sdk.RequireNotError(err, types.ErrInvalidParameter)
	err = jsoniter.Unmarshal([]byte(fee), newfee)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	sdk.RequireOwner(t.sdk)
	sdk.Require(len(newmain) == ARRAYSIZE && len(newfee) == ARRAYSIZE,
		types.ErrInvalidParameter, "length of array must be five")
	for _, v := range newmain {
		for _, value := range v {
			sdk.Require((value >= 0 && value <= 4) || value == 9 || value == 10,
				types.ErrInvalidParameter, "poker num is wrong")
		}

	}
	for _, v := range newfee {
		for _, value := range v {
			sdk.Require((value >= 0 && value <= 4) || value == 9 || value == 10,
				types.ErrInvalidParameter, "poker num is wrong")
		}

	}
	t._setPokerMainSet(*newmain)
	t._setPokerMainSet(*newfee)

}

//SetSecretSigner - Set up the public key
//@:public:method:gas[500]
func (t *Tiger) SetSecretSigner(newSecretSigner types.PubKey) {

	sdk.RequireOwner(t.sdk)
	sdk.Require(len(newSecretSigner) == 32,
		types.ErrInvalidParameter, "length of newSecretSigner must be 32 bytes")

	//Save to database
	t._setSecretSigner(newSecretSigner)

	// fire event
	t.emitSetSecretSigner(newSecretSigner)
}

//SetOwner - Set contract owner
//@:public:method:gas[500]
func (t *Tiger) SetOwner(newOwnerAddr types.Address) {
	// only contract owner just can set new owner
	sdk.RequireOwner(t.sdk)

	sdk.Require(newOwnerAddr != t.sdk.Message().Contract().Account(),
		types.ErrInvalidParameter, "NewOwner address cannot be contract account address")

	sdk.Require(newOwnerAddr != t.sdk.Message().Contract().Address(),
		types.ErrInvalidParameter, "NewOwner address cannot be contract address")

	sdk.Require(t.sdk.Helper().BlockChainHelper().CheckAddress(newOwnerAddr) == nil,
		types.ErrInvalidParameter, "NewOwner address cannot be contract account address")

	t.sdk.Message().Contract().SetOwner(newOwnerAddr)

	//fire event
	t.emitSetOwner(t.sdk.Message().Contract().Owner())

}

// SetSettings - Change game settings
//@:public:method:gas[500]
func (t *Tiger) SetSettings(newSettingsStr string) {

	sdk.RequireOwner(t.sdk)

	//Check that the Settings are valid
	newSettings := new(Settings)
	err := jsoniter.Unmarshal([]byte(newSettingsStr), newSettings)
	sdk.RequireNotError(err, types.ErrInvalidParameter)
	t.checkSettings(newSettings)

	//Settings can only be set after all settlement is completed and the refund is completed
	settings := t._settings()
	for _, tokenName := range settings.TokenNames {
		lockedAmount := t._lockedInBets(tokenName)
		sdk.Require(lockedAmount.CmpI(0) == 0,
			types.ErrUserDefined, "only lockedAmount is zero that can do SetSettings()")
	}

	t._setSettings(newSettings)

	// fire event
	t.emitSetSettings(
		newSettings.TokenNames,
		newSettings.LimitMaps,
		newSettings.FeeRatio,
		newSettings.FeeMiniNum,
		newSettings.SendToCltRatio,
		newSettings.BetExpirationBlocks,
	)
}

// SetRecFeeInfo - Set ratio of fee and receiver's account address
//@:public:method:gas[500]
func (t *Tiger) SetRecFeeInfo(recFeeInfoStr string) {

	sdk.RequireOwner(t.sdk)

	info := make([]RecFeeInfo, 0)
	err := jsoniter.Unmarshal([]byte(recFeeInfoStr), &info)
	sdk.RequireNotError(err, types.ErrInvalidParameter)
	//Check that the parameters are valid
	t.checkRecFeeInfo(info)

	t._setRecFeeInfo(info)
	// fire event
	t.emitSetRecFeeInfo(info)
}

//SampleMethod 用户充钱进来转化积分
//@:public:method:gas[500]
func (t *Tiger) DigtalCurrency(tk string, num int64) {
	//获取用户转账并和传入参数做对比
	gambler := t.sdk.Message().Sender().Address()
	//将用户的信息写入
	tokenName := ""
	tranAmount := bn.N(0)
	t.GetTransferData(t._settings(), &tokenName, &tranAmount)
	sdk.Require(tk == tokenName && tranAmount.CmpI(num) == 0,
		types.ErrInvalidParameter, "Incorrect parameters")
	contractAcct := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Account())
	settings := t._settings()
	//转到平台
	fee := tranAmount.MulI(settings.FeeRatio).DivI(PERMILLE)
	if settings.SendToCltRatio > 0 {
		clt := fee.MulI(settings.SendToCltRatio).DivI(PERMILLE)
		contractAcct.TransferByName(tokenName, t.sdk.Helper().BlockChainHelper().CalcAccountFromName("clt", ""), clt)
		fee = fee.Sub(clt)
	}
	//Transfer to other handling address
	t.transferToRecFeeAddr(tokenName, fee)

	//充值积分
	num=num*(1-settings.FeeRatio/PERMILLE)
	//判断该用户是不是存在
	if t._chkBetInfo(gambler) {
		//存在  取出来相加 让后保存
		info := t._betInfo(gambler)
		info.Currency[tokenName] += num
		t._setBetInfo(gambler, info)
	} else {
		tem := make(map[string]int64)
		//把所有支持的币种设置为0
		tem[tokenName] = num
		fee := &FeeInfo{
			FeeCount:  0,
			BetAmout:  0,
			TokenName: "",
		}
		player := &PlayerInfo{
			Address:  gambler,
			Currency: tem,
			Fee:      fee,
		}
		t._setBetInfo(gambler, player)

	}

	//fire event
    t.emitDigtalCurrency(tokenName,num)


}

// WithdrawFunds - Funds withdrawal
//@:public:method:gas[500]
func (t *Tiger) WithdrawFunds(tokenName string, beneficiary types.Address, withdrawAmount bn.Number) {

	sdk.RequireOwner(t.sdk)
	sdk.Require(withdrawAmount.CmpI(0) > 0,
		types.ErrInvalidParameter, "withdrawAmount must be larger than zero")

	account := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Account())
	lockedAmount := t._lockedInBets(tokenName)
	unlockedAmount := account.BalanceOfName(tokenName).Sub(lockedAmount)
	sdk.Require(unlockedAmount.Cmp(withdrawAmount) >= 0,
		types.ErrInvalidParameter, "Not enough funds")

	// transfer to beneficiary
	account.TransferByName(tokenName, beneficiary, withdrawAmount)

	// fire event
	t.emitWithdrawFunds(tokenName, beneficiary, withdrawAmount)
}


// PlaceBet - place bet
//@:public:method:gas[500]
func (t *Tiger) PlaceBet(reveals []byte, tokenName string, betNum, commitLastBlock int64, commit, signData []byte, refAddress types.Address) {

	playerAddress := t.sdk.Message().Sender().Address()
	sdk.Require(playerAddress != t.sdk.Message().Contract().Owner(),
		types.ErrNoAuthorization, "The contract owner cannot bet")
	//1. Verify whether the signature and current round betting are legal
	data := append(bn.N(commitLastBlock).Bytes(), commit...)
	sdk.Require(ed25519.VerifySign(t._secretSigner(), data, signData),
		types.ErrInvalidParameter, "Incorrect signature")
	reveal:=[][]byte{reveals[5:], reveals[10:], reveals[15:], reveals[20:], reveals[25:]}

	//hexCommit := hex.EncodeToString(commit)
	//Is late
	sdk.Require(t.sdk.Block().Height() <= commitLastBlock,
		types.ErrInvalidParameter, "Commit has expired")
	//查出玩家信息
	playInfo := t._betInfo(playerAddress)
	sdk.Require(playInfo != nil,
		types.ErrInvalidParameter, "Process errors, please digtalCurrency first")

	if v, ok := playInfo.Currency[tokenName]; ok {
		//存在  判断投注金额和余额的比较
		sdk.Require(v >= betNum,
			types.ErrInvalidParameter, "bet count is larger than balance")
	} else {
		//panic
		sdk.Require(false,
			types.ErrInvalidParameter, "This currency has not been recharged yet")
	}
	setting := t._settings()
	sdk.Require(bn.N(betNum).CmpI(setting.LimitMaps[tokenName].MinLimit) > 0 && bn.N(betNum).CmpI(setting.LimitMaps[tokenName].MaxLimit) <= 0,
		types.ErrInvalidParameter, "This currency has not been recharged yet")
	//按最高倍率锁币
	contractAcct := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Account())
	//Lock in the amount that may need to be paid
	totalMaybeWinAmount := bn.N(betNum * 100)

	//Contract account balance
	totalUnlockedAmount := contractAcct.BalanceOfName(tokenName)
	//Is the contract account balance greater than or equal to the balance that may need to be paid
	sdk.Require(totalUnlockedAmount.Cmp(totalMaybeWinAmount) >= 0,
		types.ErrInvalidParameter, "Cannot afford to lose this bet")

	//扣分
	playInfo.Currency[tokenName] -= betNum
	//按最大赔率来算
	totalLockedAmount := t._lockedInBets(tokenName).Add(totalMaybeWinAmount)

	sdk.Require(playInfo.Fee.FeeCount == 0,
		types.ErrInvalidParameter, "feecount  must be zero")
	t._setLockedInBets(tokenName, totalLockedAmount)
	//洗牌  连线 开奖
	t.ShuffleMainCards(reveal, playerAddress, playInfo)
	mul, times, pokerFix := t.GetNineLine(playerAddress)
	//test
	if times > 0 {
		playInfo.Fee.BetAmout = betNum
		playInfo.Fee.TokenName = tokenName
		playInfo.Fee.FeeCount += times
	}
	totaltimes := playInfo.Fee.FeeCount
	if mul > 0 { //转币给玩家
		winAmount := (betNum / 9) * mul
		sdk.Require((bn.N(int64(winAmount * PERMILLE)).DivI(PERMILLE)).CmpI(setting.LimitMaps[tokenName].MaxProfit) < 0,
			types.ErrInvalidParameter, "winAmount must be smaller than maxprofit")

		//玩家赢的不能超过最大利润
		contractAcct.TransferByName(tokenName, playerAddress, bn.N(int64(winAmount * PERMILLE)).DivI(PERMILLE))

	}
	//解锁币
	lockedInBet := t._lockedInBets(tokenName)
	t._setLockedInBets(tokenName, lockedInBet.Sub(totalMaybeWinAmount))
	t._setBetInfo(playerAddress, playInfo)

	//发送emit
	t.emitPlaceBet(
		tokenName,
		pokerFix,
		times,
		totaltimes,
		mul,
		(betNum/9)*mul,
		playerAddress,
		totalMaybeWinAmount,
		commitLastBlock,
		commit,
		signData,
		refAddress)

}

// PlaceBet - place bet
//@:public:method:gas[500]
func (t *Tiger) PlaceFeeBet(reveals []byte, tokenName string, betNum, commitLastBlock int64, commit, signData []byte, refAddress types.Address) {
	reveal:=[][]byte{reveals[5:], reveals[10:], reveals[15:], reveals[20:], reveals[25:]}
	playerAddress := t.sdk.Message().Sender().Address()
	//1. Verify whether the signature and current round betting are legal
	data := append(bn.N(commitLastBlock).Bytes(), commit...)
	sdk.Require(ed25519.VerifySign(t._secretSigner(), data, signData),
		types.ErrInvalidParameter, "Incorrect signature")
	//hexCommit := hex.EncodeToString(commit)
	//Is late
	sdk.Require(t.sdk.Block().Height() <= commitLastBlock,
		types.ErrInvalidParameter, "Commit has expired")
	//查出玩家信息
	playInfo := t._betInfo(playerAddress)
	sdk.Require(playInfo != nil,
		types.ErrInvalidParameter, "Process errors, please digtalCurrency first")
	//做测试用的写死 正式的要改回来
	//playInfo.Fee.FeeCount=1000
	//playInfo.Fee.BetAmout=100000001
	//playInfo.Fee.TokenName="LOC"
	sdk.Require(playInfo.Fee.FeeCount > 0 && playInfo.Fee.BetAmout == betNum && playInfo.Fee.TokenName == tokenName,
		types.ErrInvalidParameter, "You can't play free games")

	setting := t._settings()
	sdk.Require(bn.N(betNum).CmpI(setting.LimitMaps[tokenName].MinLimit) > 0 && bn.N(betNum).CmpI(setting.LimitMaps[tokenName].MaxLimit) <= 0,
		types.ErrInvalidParameter, "This currency has not been recharged yet")

	//按最高倍率锁币
	contractAcct := t.sdk.Helper().AccountHelper().AccountOf(t.sdk.Message().Contract().Account())
	//Lock in the amount that may need to be paid
	totalMaybeWinAmount := bn.N(betNum * 100) //按最大赔率来算
	totalLockedAmount := t._lockedInBets(tokenName).Add(totalMaybeWinAmount)
	//Contract account balance
	totalUnlockedAmount := contractAcct.BalanceOfName(tokenName)
	//Is the contract account balance greater than or equal to the balance that may need to be paid
	sdk.Require(totalUnlockedAmount.Cmp(totalMaybeWinAmount) >= 0,
		types.ErrInvalidParameter, "Cannot afford to lose this bet")
	t._setLockedInBets(tokenName, totalLockedAmount)
	//洗牌  连线 开奖
	t.ShuffleFeeCards(reveal, playInfo)
	mul, times, pokerFix := t.GetNineLine(playerAddress)
	//完了一局要减少
	playInfo.Fee.FeeCount--
	if times > 0 {
		playInfo.Fee.FeeCount += times
	}
	totaltimes := playInfo.Fee.FeeCount
	if mul > 0 {
		winAmount := (betNum / 9) * mul
		sdk.Require((bn.N(int64(winAmount * PERMILLE)).DivI(PERMILLE)).CmpI(setting.LimitMaps[tokenName].MaxProfit) < 0,
			types.ErrInvalidParameter, "winAmount must be smaller than maxprofit")
		//转币给玩家
		contractAcct.TransferByName(tokenName, playerAddress, bn.N(int64(winAmount * PERMILLE)).DivI(PERMILLE))
	}
	//解锁币  转账到平台
	lockedInBet := t._lockedInBets(tokenName)
	t._setLockedInBets(tokenName, lockedInBet.Sub(totalMaybeWinAmount))
	t._setBetInfo(playerAddress, playInfo)

	//发送emit
	t.emitPlaceBet(
		tokenName,
		pokerFix,
		times,
		totaltimes,
		mul,
		(betNum/9)*mul,
		playerAddress,
		totalMaybeWinAmount,
		commitLastBlock,
		commit,
		signData,
		refAddress)
}
