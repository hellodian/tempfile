package sicbo

import (
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdk/bn"
)

var _ receipt = (*SicBo)(nil)

//emitSetSecretSigner This is a method of SicBo
func (sb *SicBo) emitSetSecretSigner(newSecretSigner types.PubKey) {
	type setSecretSigner struct {
		NewSecretSigner types.PubKey `json:"newSecretSigner"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		setSecretSigner{
			NewSecretSigner: newSecretSigner,
		},
	)
}

//emitSetOwner This is a method of SicBo
func (sb *SicBo) emitSetOwner(newAddress types.Address) {
	type setOwner struct {
		NewAddress types.Address `json:"newAddress"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		setOwner{
			NewAddress: newAddress,
		},
	)
}

//emitSetSettings This is a method of SicBo
func (sb *SicBo) emitSetSettings(tokenNames []string, limit map[string]Limit, feeRatio, feeMiniNum, sendToCltRatio, betExpirationBlocks int64) {
	type setSettings struct {
		TokenNames          []string         `json:"tokenNames"`
		Limit               map[string]Limit `json:"limit"`
		FeeRatio            int64            `json:"feeRatio"`
		FeeMiniNum          int64            `json:"feeMiniNum"`
		SendToCltRatio      int64            `json:"sendToCltRatio"`
		BetExpirationBlocks int64            `json:"betExpirationBlocks"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		setSettings{
			TokenNames:          tokenNames,
			Limit:               limit,
			FeeRatio:            feeRatio,
			FeeMiniNum:          feeMiniNum,
			SendToCltRatio:      sendToCltRatio,
			BetExpirationBlocks: betExpirationBlocks,
		},
	)
}

//emitSetRecFeeInfo This is a method of SicBo
func (sb *SicBo) emitSetRecFeeInfo(info []RecFeeInfo) {
	type setRecFeeInfo struct {
		Info []RecFeeInfo `json:"info"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		setRecFeeInfo{
			Info: info,
		},
	)
}

//emitWithdrawFunds This is a method of SicBo
func (sb *SicBo) emitWithdrawFunds(tokenName string, beneficiary types.Address, withdrawAmount bn.Number) {
	type withdrawFunds struct {
		TokenName      string        `json:"tokenName"`
		Beneficiary    types.Address `json:"beneficiary"`
		WithdrawAmount bn.Number     `json:"withdrawAmount"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		withdrawFunds{
			TokenName:      tokenName,
			Beneficiary:    beneficiary,
			WithdrawAmount: withdrawAmount,
		},
	)
}

//emitPlaceBet This is a method of SicBo
func (sb *SicBo) emitPlaceBet(tokenName string, gambler types.Address, totalMaybeWinAmount bn.Number, betDataList []BetData, commitLastBlock int64, commit, signData []byte, refAddress types.Address) {
	type placeBet struct {
		TokenName           string        `json:"tokenName"`
		Gambler             types.Address `json:"gambler"`
		TotalMaybeWinAmount bn.Number     `json:"totalMaybeWinAmount"`
		BetDataList         []BetData     `json:"betDataList"`
		CommitLastBlock     int64         `json:"commitLastBlock"`
		Commit              []byte        `json:"commit"`
		SignData            []byte        `json:"signData"`
		RefAddress          types.Address `json:"refAddress"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		placeBet{
			TokenName:           tokenName,
			Gambler:             gambler,
			TotalMaybeWinAmount: totalMaybeWinAmount,
			BetDataList:         betDataList,
			CommitLastBlock:     commitLastBlock,
			Commit:              commit,
			SignData:            signData,
			RefAddress:          refAddress,
		},
	)
}

//emitSettleBet This is a method of SicBo
func (sb *SicBo) emitSettleBet(tokenName []string, commit []byte, gambler []types.Address, sic Sic, winNumber int64, totalWinAmount map[string]bn.Number, finished bool) {
	type settleBet struct {
		TokenName      []string             `json:"tokenName"`
		Commit         []byte               `json:"commit"`
		Gambler        []types.Address      `json:"gambler"`
		Sic            Sic                  `json:"sic"`
		WinNumber      int64                `json:"winNumber"`
		TotalWinAmount map[string]bn.Number `json:"totalWinAmount"`
		Finished       bool                 `json:"finished"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		settleBet{
			TokenName:      tokenName,
			Commit:         commit,
			Gambler:        gambler,
			Sic:            sic,
			WinNumber:      winNumber,
			TotalWinAmount: totalWinAmount,
			Finished:       finished,
		},
	)
}

//emitRefundBet This is a method of SicBo
func (sb *SicBo) emitRefundBet(commit []byte, tokenName []string, gambler []types.Address, refundedAmount map[string]bn.Number, finished bool) {
	type refundBet struct {
		Commit         []byte               `json:"commit"`
		TokenName      []string             `json:"tokenName"`
		Gambler        []types.Address      `json:"gambler"`
		RefundedAmount map[string]bn.Number `json:"refundedAmount"`
		Finished       bool                 `json:"finished"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		refundBet{
			Commit:         commit,
			TokenName:      tokenName,
			Gambler:        gambler,
			RefundedAmount: refundedAmount,
			Finished:       finished,
		},
	)
}

//emitFeeFlag This is a method of SicBo
func (sb *SicBo) emitFeeFlag(flag bool) {
	type feeFlag struct {
		Flag bool `json:"flag"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		feeFlag{
			Flag: flag,
		},
	)
}

//emitBigOrSmall This is a method of SicBo
func (sb *SicBo) emitBigOrSmall(num int64) {
	type bigOrSmall struct {
		Num int64 `json:"num"`
	}

	sb.sdk.Helper().ReceiptHelper().Emit(
		bigOrSmall{
			Num: num,
		},
	)
}
        
                         