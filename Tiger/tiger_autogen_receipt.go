package tiger

import (
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdk/bn"
)

var _ receipt = (*Tiger)(nil)

//emitSetSecretSigner This is a method of Tiger
func (t *Tiger) emitSetSecretSigner(newSecretSigner types.PubKey) {
	type setSecretSigner struct {
		NewSecretSigner types.PubKey `json:"newSecretSigner"`
	}

	t.sdk.Helper().ReceiptHelper().Emit(
		setSecretSigner{
			NewSecretSigner: newSecretSigner,
		},
	)
}

//emitSetOwner This is a method of Tiger
func (t *Tiger) emitSetOwner(newAddress types.Address) {
	type setOwner struct {
		NewAddress types.Address `json:"newAddress"`
	}

	t.sdk.Helper().ReceiptHelper().Emit(
		setOwner{
			NewAddress: newAddress,
		},
	)
}

//emitSetSettings This is a method of Tiger
func (t *Tiger) emitSetSettings(tokenNames []string, limit map[string]Limit, feeRatio, feeMiniNum, sendToCltRatio, betExpirationBlocks int64) {
	type setSettings struct {
		TokenNames          []string         `json:"tokenNames"`
		Limit               map[string]Limit `json:"limit"`
		FeeRatio            int64            `json:"feeRatio"`
		FeeMiniNum          int64            `json:"feeMiniNum"`
		SendToCltRatio      int64            `json:"sendToCltRatio"`
		BetExpirationBlocks int64            `json:"betExpirationBlocks"`
	}

	t.sdk.Helper().ReceiptHelper().Emit(
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

//emitSetRecFeeInfo This is a method of Tiger
func (t *Tiger) emitSetRecFeeInfo(info []RecFeeInfo) {
	type setRecFeeInfo struct {
		Info []RecFeeInfo `json:"info"`
	}

	t.sdk.Helper().ReceiptHelper().Emit(
		setRecFeeInfo{
			Info: info,
		},
	)
}

//emitAssemblePoker This is a method of Tiger
func (t *Tiger) emitAssemblePoker(num int64, listPoker []Poker) {
	type assemblePoker struct {
		Num       int64   `json:"num"`
		ListPoker []Poker `json:"listPoker"`
	}

	t.sdk.Helper().ReceiptHelper().Emit(
		assemblePoker{
			Num:       num,
			ListPoker: listPoker,
		},
	)
}

//emitDigtalCurrency This is a method of Tiger
func (t *Tiger) emitDigtalCurrency(token string, num int64) {
	type digtalCurrency struct {
		Token string `json:"token"`
		Num   int64  `json:"num"`
	}

	t.sdk.Helper().ReceiptHelper().Emit(
		digtalCurrency{
			Token: token,
			Num:   num,
		},
	)
}

//emitWithdrawFunds This is a method of Tiger
func (t *Tiger) emitWithdrawFunds(tokenName string, beneficiary types.Address, withdrawAmount bn.Number) {
	type withdrawFunds struct {
		TokenName      string        `json:"tokenName"`
		Beneficiary    types.Address `json:"beneficiary"`
		WithdrawAmount bn.Number     `json:"withdrawAmount"`
	}

	t.sdk.Helper().ReceiptHelper().Emit(
		withdrawFunds{
			TokenName:      tokenName,
			Beneficiary:    beneficiary,
			WithdrawAmount: withdrawAmount,
		},
	)
}

//emitPlaceBet This is a method of Tiger
func (t *Tiger) emitPlaceBet(tokenName string, pokerFix [][]Poker, mul, times, totaltimes int64, sum int64, gambler types.Address, WinAmount bn.Number, commitLastBlock int64, commit, signData []byte, refAddress types.Address) {
	type placeBet struct {
		TokenName       string        `json:"tokenName"`
		PokerFix        [][]Poker     `json:"pokerFix"`
		Mul             int64         `json:"mul"`
		Times           int64         `json:"times"`
		Totaltimes      int64         `json:"totaltimes"`
		Sum             int64         `json:"sum"`
		Gambler         types.Address `json:"gambler"`
		WinAmount       bn.Number     `json:"WinAmount"`
		CommitLastBlock int64         `json:"commitLastBlock"`
		Commit          []byte        `json:"commit"`
		SignData        []byte        `json:"signData"`
		RefAddress      types.Address `json:"refAddress"`
	}

	t.sdk.Helper().ReceiptHelper().Emit(
		placeBet{
			TokenName:       tokenName,
			PokerFix:        pokerFix,
			Mul:             mul,
			Times:           times,
			Totaltimes:      totaltimes,
			Sum:             sum,
			Gambler:         gambler,
			WinAmount:       WinAmount,
			CommitLastBlock: commitLastBlock,
			Commit:          commit,
			SignData:        signData,
			RefAddress:      refAddress,
		},
	)
}
        
                         