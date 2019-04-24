package sicbo

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/crypto/ed25519"
	"blockchain/smcsdk/sdk/crypto/sha3"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/types"
	"encoding/hex"
)

//@:public:receipt
type receipt interface {
	emitSetSecretSigner(newSecretSigner types.PubKey)
	emitSetOwner(newAddress types.Address)
	emitSetSettings(tokenNames []string, limit map[string]Limit, feeRatio, feeMiniNum, sendToCltRatio, betExpirationBlocks int64)
	emitSetRecFeeInfo(info []RecFeeInfo)
	emitWithdrawFunds(tokenName string, beneficiary types.Address, withdrawAmount bn.Number)
	emitPlaceBet(tokenName string, gambler types.Address, totalMaybeWinAmount bn.Number, betDataList []BetData, commitLastBlock int64, commit, signData []byte, refAddress types.Address)
	emitSettleBet(tokenName []string, commit []byte, gambler []types.Address, sic Sic, winNumber int64, totalWinAmount map[string]bn.Number, finished bool)
	emitRefundBet(commit []byte, tokenName []string, gambler []types.Address, refundedAmount map[string]bn.Number, finished bool)
	emitFeeFlag(flag bool)
	emitBigOrSmall(num int64)
}

//InitChain - InitChain Constructor of this SicBo
//@:constructor
func (sb *SicBo) InitChain() {

	// init dataN
	settings := SBSettings{}

	settings.FeeRatio = 50
	settings.FeeMiniNum = 300000
	settings.SendToCltRatio = 100
	settings.BetExpirationBlocks = 250
	settings.TokenNames = []string{sb.sdk.Helper().GenesisHelper().Token().Name()}
	//settings.TokenNames = []string{"BCB", "DC", "XT", "USDX"}

	limitMaps := sb.createLimitMaps(settings.TokenNames)

	settings.LimitMaps = limitMaps
	sb._setSettings(&settings)
	sb.LockedInBetsInit(settings.TokenNames)
	sb._setBigOrSmall(1)
}


//SetSecretSigner - Set up the public key
//@:public:method:gas[500]
func (sb *SicBo) SetSecretSigner(newSecretSigner types.PubKey) {

	sdk.RequireOwner(sb.sdk)
	sdk.Require(len(newSecretSigner) == 32,
		types.ErrInvalidParameter, "length of newSecretSigner must be 32 bytes")

	//Save to database
	sb._setSecretSigner(newSecretSigner)

	// fire event
	sb.emitSetSecretSigner(newSecretSigner)
}

//SetOwner - Set contract owner
//@:public:method:gas[500]
func (sb *SicBo) SetOwner(newOwnerAddr types.Address) {
	// only contract owner just can set new owner
	sdk.RequireOwner(sb.sdk)

	sdk.Require(newOwnerAddr != sb.sdk.Message().Contract().Account(),
		types.ErrInvalidParameter, "NewOwner address cannot be contract account address")

	sdk.Require(newOwnerAddr != sb.sdk.Message().Contract().Address(),
		types.ErrInvalidParameter, "NewOwner address cannot be contract address")

	sdk.Require(sb.sdk.Helper().BlockChainHelper().CheckAddress(newOwnerAddr)==nil,
		types.ErrInvalidParameter, "NewOwner address cannot be contract account address")

	sb.sdk.Message().Contract().SetOwner(newOwnerAddr)

	//fire event
	sb.emitSetOwner(sb.sdk.Message().Contract().Owner())

}

// SetSettings - Change game settings
//@:public:method:gas[500]
func (sb *SicBo) SetSettings(newSettingsStr string) {

	sdk.RequireOwner(sb.sdk)

	//Check that the Settings are valid
	newSettings := new(SBSettings)
	err := jsoniter.Unmarshal([]byte(newSettingsStr), newSettings)
	sdk.RequireNotError(err, types.ErrInvalidParameter)
	sb.checkSettings(newSettings)

	//Settings can only be set after all settlement is completed and the refund is completed
	settings := sb._settings()
	for _, tokenName := range settings.TokenNames {
		lockedAmount := sb._lockedInBets(tokenName)
		sdk.Require(lockedAmount.CmpI(0) == 0,
			types.ErrUserDefined, "only lockedAmount is zero that can do SetSettings()")
	}

	sb._setSettings(newSettings)

	// fire event
	sb.emitSetSettings(
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
func (sb *SicBo) SetRecFeeInfo(recFeeInfoStr string) {

	sdk.RequireOwner(sb.sdk)

	info := make([]RecFeeInfo, 0)
	err := jsoniter.Unmarshal([]byte(recFeeInfoStr), &info)
	sdk.RequireNotError(err, types.ErrInvalidParameter)
	//Check that the parameters are valid
	sb.checkRecFeeInfo(info)

	sb._setRecFeeInfo(info)
	// fire event
	sb.emitSetRecFeeInfo(info)
}

// WithdrawFunds - Funds withdrawal
//@:public:method:gas[500]
func (sb *SicBo) WithdrawFunds(tokenName string, beneficiary types.Address, withdrawAmount bn.Number) {

	sdk.RequireOwner(sb.sdk)
	sdk.Require(withdrawAmount.CmpI(0) > 0,
		types.ErrInvalidParameter, "withdrawAmount must be larger than zero")

	account := sb.sdk.Helper().AccountHelper().AccountOf(sb.sdk.Message().Contract().Account())
	lockedAmount := sb._lockedInBets(tokenName)
	unlockedAmount := account.BalanceOfName(tokenName).Sub(lockedAmount)
	sdk.Require(unlockedAmount.Cmp(withdrawAmount) >= 0,
		types.ErrInvalidParameter, "Not enough funds")

	// transfer to beneficiary
	account.TransferByName(tokenName, beneficiary, withdrawAmount)

	// fire event
	sb.emitWithdrawFunds(tokenName, beneficiary, withdrawAmount)
}

//PlaceBet - Bet
//@:public:method:gas[500]
func (sb *SicBo) PlaceBet(betInfoJson string, commitLastBlock int64, betIndex string, commit, signData []byte, refAddress types.Address) {

	sdk.Require(sb.sdk.Message().Sender().Address() != sb.sdk.Message().Contract().Owner(),
		types.ErrNoAuthorization, "The contract owner cannot bet")

	//1. Verify whether the signature and current round betting are legal
	data := append(bn.N(commitLastBlock).Bytes(), commit...)
	sdk.Require(ed25519.VerifySign(sb._secretSigner(), data, signData),
		types.ErrInvalidParameter, "Incorrect signature")

	gambler := sb.sdk.Message().Sender().Address()
	hexCommit := hex.EncodeToString(commit)
	//Is late
	sdk.Require(sb.sdk.Block().Height() <= commitLastBlock,
		types.ErrInvalidParameter, "Commit has expired")

	//Verify that the betIndex has made a note
	sdk.Require(!sb._chkBetInfo(hexCommit, betIndex),
		types.ErrInvalidParameter, "Commit and index should be new")

	//sdk.Require(sb._chkRoundInfo(hexCommit), types.ErrInvalidParameter, "The current wheel does not exist")
	settings := sb._settings()
	var roundInfo *RoundInfo
	if !sb._chkRoundInfo(hexCommit) {
		roundInfo = &RoundInfo{
			Commit:           commit,
			State:            NOAWARD,
			FirstBlockHeight: sb.sdk.Block().Height(),
			Settings:         settings,
			TotalBuyAmount:   sb.CreateMapByTokenName(settings.TokenNames),
		}
		sb._setRoundInfo(hexCommit, roundInfo)
	}
	roundInfo = sb._roundInfo(hexCommit)
	//Whether the current wheel state allows betting
	sdk.Require(NOAWARD == roundInfo.State,
		types.ErrInvalidParameter, "No betting on the current wheel")
	sdk.Require(roundInfo.FirstBlockHeight+roundInfo.Settings.BetExpirationBlocks > sb.sdk.Block().Height(),
		types.ErrInvalidParameter, "This round is time out")

	tokenName := ""
	transferAmount := bn.N(0)
	//Obtain transfer records
	for _, name := range settings.TokenNames {
		transferReceipt := sb.sdk.Message().GetTransferToMe(name)
		if transferReceipt != nil {
			tokenName = name
			transferAmount = transferReceipt.Value
			break
		}
	}

	//Verify receipt of bet transfer
	sdk.Require(tokenName != "" && transferAmount.CmpI(0) > 0,
		types.ErrUserDefined, "Must transfer tokens to me before place a bet")

	betDataList := make([]BetData, 0)
	err := jsoniter.Unmarshal([]byte(betInfoJson), &betDataList)
	sdk.RequireNotError(err, types.ErrInvalidParameter)

	totalAmount := bn.N(0)
	for _, betData := range betDataList {
		amount := betData.BetAmount

		//Verify that the betting scheme is legal
		sdk.Require(betData.BetMode >= BIG && betData.BetMode <= ALLSAMESIC,
			types.ErrInvalidParameter, "The betting range is illegal")
		//Verify that the bet amount is legal
		sdk.Require(amount.CmpI(settings.LimitMaps[tokenName].MinLimit) >= 0 && amount.CmpI(settings.LimitMaps[tokenName].MaxLimit) <= 0,
			types.ErrInvalidParameter, "Amount should be within range")

		totalAmount = totalAmount.Add(amount)
	}

	//Check whether the total amount of bet is equal to the transfer amount
	sdk.Require(totalAmount.Cmp(transferAmount) == 0,
		types.ErrUserDefined, "transfer amount not equal place bet amount")

	//Verify that the reimbursable amount is sufficient
	//Calculate the amount of money you can win
	totalMaybeWinAmount, feeAmount := sb.MaybeWinAmountAndFeeByList(totalAmount, betDataList)
	//Is the amount likely to be won less than or equal to the maximum bonus amount
	sdk.Require(totalMaybeWinAmount.CmpI(settings.LimitMaps[tokenName].MaxProfit) <= 0,
		types.ErrInvalidParameter, "MaxProfit limit violation")

	contractAcct := sb.sdk.Helper().AccountHelper().AccountOf(sb.sdk.Message().Contract().Account())
	//Lock in the amount that may need to be paid
	totalLockedAmount := sb._lockedInBets(tokenName).Add(totalMaybeWinAmount).Add(feeAmount)
	//Contract account balance
	totalUnlockedAmount := contractAcct.BalanceOfName(tokenName)
	//Is the contract account balance greater than or equal to the balance that may need to be paid
	sdk.Require(totalUnlockedAmount.Cmp(totalMaybeWinAmount) >= 0,
		types.ErrInvalidParameter, "Cannot afford to lose this bet")
	sb._setLockedInBets(tokenName, totalLockedAmount)

	//Store bet information
	betInfo := &BetInfo{}
	betInfo.TokenName = tokenName
	betInfo.Amount = totalAmount
	betInfo.BetData = betDataList
	betInfo.WinAmount = totalMaybeWinAmount
	betInfo.Settled = false
	betInfo.Gambler = gambler

	sb._setBetInfo(hexCommit, betIndex, betInfo)

	roundInfo = sb._roundInfo(hexCommit)
	//Round information add betting information k2
	roundInfo.BetInfoSerialNumber = append(roundInfo.BetInfoSerialNumber, betIndex)
	roundInfo.TotalBuyAmount[tokenName] = roundInfo.TotalBuyAmount[tokenName].Add(totalAmount)
	roundInfo.TotalBetCount += 1
	//Save the current wheel information
	sb._setRoundInfo(hexCommit, roundInfo)

	sb.emitPlaceBet(tokenName, gambler, totalMaybeWinAmount, betDataList, commitLastBlock, commit, signData, refAddress)
}

//SettleBet - The lottery and settlement
//@:public:method:gas[500]
func (sb *SicBo) SettleBet(reveal []byte, settleCount int64) {

	sdk.Require(len(reveal) > 0,
		types.ErrInvalidParameter, "Commit should be not exist")

	sdk.RequireOwner(sb.sdk)
	hexCommit := hex.EncodeToString(sha3.Sum256(reveal))
	sdk.Require(sb._chkRoundInfo(hexCommit), types.ErrInvalidParameter, "Commit should be not exist")

	roundInfo := sb._roundInfo(hexCommit)
	//Current wheel configuration
	settings := roundInfo.Settings
	//The bet height of the round to be settled should be less than the settlement height
	sdk.Require(roundInfo.FirstBlockHeight < sb.sdk.Block().Height(),
		types.ErrInvalidParameter, "SettleBet block can not be in the same block as placeBet, or before.")

	sdk.Require(NOAWARD == roundInfo.State || OPENINGAPRIZE == roundInfo.State,
		types.ErrInvalidParameter, "This state does not operate for settlement")

	//For the first settlement, the round information cannot expire
	if NOAWARD == roundInfo.State {
		sdk.Require(roundInfo.FirstBlockHeight+settings.BetExpirationBlocks > sb.sdk.Block().Height(),
			types.ErrInvalidParameter, "This round is time out")
		roundInfo.Sic = sb.GetSicRes(reveal)
		//No one bets on the current round
		if roundInfo.TotalBetCount <= 0 {
			roundInfo.State = AWARDED
			sb._setRoundInfo(hexCommit, roundInfo)
			return
		}
		roundInfo.State = OPENINGAPRIZE
		sb._setRoundInfo(hexCommit, roundInfo)
	}
	//Determine whether all bets have been settled
	if roundInfo.TotalBetCount == roundInfo.ProcessCount {
		roundInfo.State = AWARDED
		sb._setRoundInfo(hexCommit, roundInfo)
		sdk.Require(false,
			types.ErrInvalidParameter, "This round is complete")
	}
	//Initial index
	startIndex := roundInfo.ProcessCount
	if startIndex < 0 {
		startIndex = 0
	}
	endIndex := startIndex + settleCount

	if endIndex >= roundInfo.TotalBetCount {
		endIndex = roundInfo.TotalBetCount
		//Set the database state state to lottery
		roundInfo.State = AWARDED
	}
	sb.DealSettle(roundInfo, hexCommit, startIndex, endIndex)
}

//WithdrawWin - Player settlement
//@:public:method:gas[500]
func (sb *SicBo) WithdrawWin (commit []byte)  {
	hexCommit := hex.EncodeToString(commit)
	sdk.Require(sb._chkRoundInfo(hexCommit), types.ErrInvalidParameter, "Commit should be not exist")
	roundInfo := sb._roundInfo(hexCommit)
	//The bet height of the round to be settled should be less than the settlement height
	sdk.Require(roundInfo.FirstBlockHeight < sb.sdk.Block().Height(),
		types.ErrInvalidParameter, "SettleBet block can not be in the same block as placeBet, or before.")

	sdk.Require(AWARDED != roundInfo.State,
		types.ErrInvalidParameter, "The operators have yet to draw a prize")

	totalBetCount := roundInfo.TotalBetCount
	sdk.Require(totalBetCount != roundInfo.ProcessCount,
		types.ErrInvalidParameter, "This round is complete")

	startIndex := int64(0)
	sb.DealSettle(roundInfo, hexCommit, startIndex, totalBetCount)
}


//RefundBets - Refund will be made if the prize is not paid after the time limit
//@:public:method:gas[500]
func (sb *SicBo) RefundBets(commit []byte, refundCount int64) {

	sdk.Require(len(commit) > 0, types.ErrInvalidParameter, "Commit should be not exist")

	sdk.RequireOwner(sb.sdk)

	hexCommit := hex.EncodeToString(commit)

	sdk.Require(sb._chkRoundInfo(hexCommit),
		types.ErrInvalidParameter, "Commit should be not exist")

	//Determine whether the bet can be refunded
	roundInfo := sb._roundInfo(hexCommit)
	//Current wheel configuration
	settings := roundInfo.Settings
	//The bet height of the round to be settled should be less than the settlement height
	sdk.Require(sb.sdk.Block().Height() > roundInfo.FirstBlockHeight+settings.BetExpirationBlocks,
		types.ErrInvalidParameter, "SettleBet block can not be in the same block as placeBet, or before.")
	//Whether the current round status can be refunded
	sdk.Require(NOAWARD == roundInfo.State,
		types.ErrInvalidParameter, "This status does not operate for a refund")
	//Whether the number of bets processed is less than the total number of bets
	sdk.Require(roundInfo.TotalBetCount > roundInfo.ProcessCount,
		types.ErrInvalidParameter, "There are currently no refundable bets")

	betInfoSerialNumberList := roundInfo.BetInfoSerialNumber
	//Whether betting information exists
	sdk.Require(len(betInfoSerialNumberList) > 0,
		types.ErrInvalidParameter, "There are currently no refundable bets")

	sic := roundInfo.Sic
	//Whether the dice result has been drawn
	sdk.Require(sic == nil || sic.One <= 0 || sic.Two <= 0 || sic.Three <= 0,
		types.ErrInvalidParameter, "The current round has been lottery, can not operate refund")

	//Contract account
	contractAcct := sb.sdk.Helper().AccountHelper().AccountOf(sb.sdk.Message().Contract().Account())

	tokenNameList := settings.TokenNames

	//Money that could be won
	totalPossibleWinAmount := sb.CreateMapByTokenName(tokenNameList)
	refundedAmount := sb.CreateMapByTokenName(tokenNameList)

	//Initial index
	startIndex := roundInfo.ProcessCount
	if startIndex < 0 {
		startIndex = 0
	}
	endIndex := startIndex + refundCount

	if endIndex >= roundInfo.TotalBetCount {
		endIndex = roundInfo.TotalBetCount
		roundInfo.State = REFUNDED
	}
	for i := startIndex; i < endIndex; i++ {
		betInfoKey := betInfoSerialNumberList[i]

		//for _, betInfoKey := range betInfoSerialNumberList {
		betInfo := sb._betInfo(hexCommit, betInfoKey)
		if betInfo.Settled {
			continue
		}
		//The currency name of the bet such as BCB
		tokenName := betInfo.TokenName
		amount := betInfo.Amount
		//If the bet amount is greater than 0, the bet will be transferred to the player's address
		if betInfo.Amount.CmpI(0) > 0 {
			contractAcct.TransferByName(tokenName, betInfo.Gambler, amount)
			refundedAmount[tokenName] = refundedAmount[tokenName].Add(amount)
		}
		//Update the current betting status database to be settled
		betInfo.Settled = true
		sb._setBetInfo(hexCommit, betInfoKey, betInfo)

		//Calculate the amount of money you can win
		totalMaybeWinAmount, feeAmount := sb.MaybeWinAmountAndFeeByList(amount, betInfo.BetData)
		totalPossibleWinAmount[tokenName] = totalPossibleWinAmount[tokenName].Sub(totalMaybeWinAmount).Sub(feeAmount)

	}

	//Unlock lock amount
	for _, tokenName := range tokenNameList {
		lockedInBet := sb._lockedInBets(tokenName)
		sb._setLockedInBets(tokenName, lockedInBet.Sub(totalPossibleWinAmount[tokenName]))
	}

	//Set the database state state to refunded, and ProcessCount updates it
	roundInfo.ProcessCount = endIndex
	sb._setRoundInfo(hexCommit, roundInfo)

	//Send the receipt
	sb.emitRefundBet(commit, tokenNameList, betInfoSerialNumberList, refundedAmount, roundInfo.State == REFUNDED)
}

//SetBigOrSmall - set  bigorsmall Odds
//@:public:method:gas[500]
func (sb *SicBo) SetBigOrSmall(num int64) {
	sdk.RequireOwner(sb.sdk)
	sb._setBigOrSmall(num)
	sb.emitBigOrSmall(sb._bigOrSmall())
}