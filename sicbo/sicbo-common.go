package sicbo

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"fmt"
)

func (sb *SicBo) checkSettings(newSettings *SBSettings) {

	sdk.Require(len(newSettings.TokenNames) > 0,
		types.ErrInvalidParameter, "tokenNames cannot be empty")

	for _, tokenName := range newSettings.TokenNames {
		token := sb.sdk.Helper().TokenHelper().TokenOfName(tokenName)
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

func (sb *SicBo) checkRecFeeInfo(infos []RecFeeInfo) {
	sdk.Require(len(infos) > 0,
		types.ErrInvalidParameter, "The length of RecvFeeInfos must be larger than zero")

	allRatio := int64(0)
	for _, info := range infos {
		sdk.Require(info.RecFeeRatio > 0,
			types.ErrInvalidParameter, "ratio must be larger than zero")
		sdk.RequireAddress(sb.sdk, info.RecFeeAddr)
		sdk.Require(info.RecFeeAddr != sb.sdk.Message().Contract().Account(),
			types.ErrInvalidParameter, "address cannot be contract account address")

		allRatio += info.RecFeeRatio
	}

	//The allocation ratio set must add up to 1000
	sdk.Require(allRatio <= 1000,
		types.ErrInvalidParameter, "The sum of ratio must be less or equal 1000")
}

//Initializes the map according to tokenNameList
func (sb *SicBo) CreateMapByTokenName(tokenNameList []string) (maps map[string]bn.Number) {
	maps = make(map[string]bn.Number, len(tokenNameList))
	for _, value := range tokenNameList {
		maps[value] = bn.N(0)
	}
	return
}

//Transfer to fee's receiving address
func (sb *SicBo) transferToRecFeeAddr(tokenNameList []types.Address, recFeeMap map[string]bn.Number) {

	account := sb.sdk.Helper().AccountHelper().AccountOf(sb.sdk.Message().Contract().Account())
	for _, tokenName := range tokenNameList {
		recFee := recFeeMap[tokenName]
		if recFee.CmpI(0) <= 0 {
			continue
		}
		infos := sb._recFeeInfo()
		for _, info := range infos {
			account.TransferByName(tokenName, info.RecFeeAddr, recFee.MulI(info.RecFeeRatio).DivI(PERMILLE))
		}
	}
}

//Settlement public method
func (sb *SicBo) DealSettle(roundInfo *RoundInfo, hexCommit string, startIndex int64, endIndex int64)  {
	settings := roundInfo.Settings
	tokenNameList := settings.TokenNames
	//Money that could be won
	totalPossibleWinAmount := sb.CreateMapByTokenName(tokenNameList)
	//The actual winning money
	totalWinAmount := sb.CreateMapByTokenName(tokenNameList)
	//Total handling charge
	totalFeeAmount := sb.CreateMapByTokenName(tokenNameList)
	//Key of betting information
	betInfoSerialNumberList := roundInfo.BetInfoSerialNumber
	//The lottery information
	sic := roundInfo.Sic
	//Contract account
	contractAcct := sb.sdk.Helper().AccountHelper().AccountOf(sb.sdk.Message().Contract().Account())
	var winCount int64 = 0

	for i := startIndex; i < endIndex; i++ {
		betInfoKey := betInfoSerialNumberList[i]
		//for _, betInfoKey := range betInfoSerialNumberList {
		betInfo := sb._betInfo(hexCommit, betInfoKey)
		if betInfo.Settled {
			continue
		}
		//The currency name of the bet such as BCB
		tokenName := betInfo.TokenName
		//The actual winning amount, the actual winning amount
		//winAmount := sic.GetBetWinAmount(betInfo.BetData)
		winAmount := sb.GetBetWinAmount(sic,betInfo.BetData)
		totalMaybeWinAmount, feeAmount := sb.MaybeWinAmountAndFeeByList(betInfo.Amount, betInfo.BetData)
		//If the winning amount is greater than 0, transfer the prize money to the player's address
		if winAmount.CmpI(0) > 0 {
			//if feeFlag true,Deduction of handling fees
				winAmount = winAmount.Sub(feeAmount)
			totalWinAmount[tokenName] = totalWinAmount[tokenName].Add(winAmount)
			contractAcct.TransferByName(tokenName, betInfo.Gambler, winAmount)
			winCount++
		}
		//Update the current betting status database to be settled
		betInfo.Settled = true
		sb._setBetInfo(hexCommit, betInfoKey, betInfo)
		//Calculate the amount of money you can win
		totalPossibleWinAmount[tokenName] = totalPossibleWinAmount[tokenName].Add(totalMaybeWinAmount).Add(feeAmount)
		//fee
		totalFeeAmount[tokenName] = totalFeeAmount[tokenName].Add(feeAmount)
	}

	//Unlock lock amount
	for _, tokenName := range tokenNameList {
		lockedInBet := sb._lockedInBets(tokenName)
		sb._setLockedInBets(tokenName, lockedInBet.Sub(totalPossibleWinAmount[tokenName]))
	}

	//participation in profit
	if settings.SendToCltRatio > 0 {
		for _, tokenName := range tokenNameList {
			amount := totalFeeAmount[tokenName].MulI(roundInfo.Settings.SendToCltRatio).DivI(PERMILLE)
			contractAcct.TransferByName(tokenName, sb.sdk.Helper().BlockChainHelper().CalcAccountFromName("clt",""), amount)
			totalFeeAmount[tokenName] = totalFeeAmount[tokenName].Sub(amount)
		}
	}

	//Transfer to other handling address
	sb.transferToRecFeeAddr(tokenNameList, totalFeeAmount)

	roundInfo.ProcessCount = endIndex
	sb._setRoundInfo(hexCommit, roundInfo)

	//Send the receipt
	sb.emitSettleBet(tokenNameList, roundInfo.Commit, roundInfo.BetInfoSerialNumber, *roundInfo.Sic, winCount, totalWinAmount, roundInfo.State == AWARDED)
}