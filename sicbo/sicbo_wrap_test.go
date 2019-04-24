package sicbo

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/smcsdk/sdkimpl/sdkhelper"
	"blockchain/smcsdk/utest"
	"fmt"
)

var (

	contractName       = "sicbo" //contract name
	contractMethods    = []string{"SetSecretSigner(types.PubKey)", "SetSettings(string)", "SetRecFeeInfo(string)", "WithdrawFunds(string,types.Address,bn.Number)", "PlaceBet(string,int64,string,[]byte,[]byte,types.Address)", "SettleBet([]byte,int64)", "WithdrawWin([]byte)", "RefundBets([]byte,int64)"}
	contractInterfaces = []string{}
	orgID              = "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111"

)

//TestObject This is a struct for test
type TestObject struct {
	obj *SicBo
}

//FuncRecover recover panic by Assert
func FuncRecover(err *types.Error) {
	if rerr := recover(); rerr != nil {
		if _, ok := rerr.(types.Error); ok {
			err.ErrorCode = rerr.(types.Error).ErrorCode
			err.ErrorDesc = rerr.(types.Error).ErrorDesc
			fmt.Println(err)
		} else {
			panic(rerr)
		}
	}
}

//NewTestObject This is a function
func NewTestObject(sender sdk.IAccount) *TestObject {
	return &TestObject{&SicBo{sdk: utest.UTP.ISmartContract}}
}

//transfer This is a method of TestObject
func (t *TestObject) transfer(balance bn.Number) *TestObject {
	contract := t.obj.sdk.Message().Contract()
	utest.Transfer(t.obj.sdk.Message().Sender(), t.obj.sdk.Helper().GenesisHelper().Token().Name(), contract.Account(), balance)
	t.obj.sdk = sdkhelper.OriginNewMessage(t.obj.sdk, contract, t.obj.sdk.Message().MethodID(), t.obj.sdk.Message().(*object.Message).OutputReceipts())
	return t
}

//setSender This is a method of TestObject
func (t *TestObject) setSender(sender sdk.IAccount) *TestObject {
	t.obj.sdk = utest.SetSender(sender.Address())
	return t
}

//run This is a method of TestObject
func (t *TestObject) run() *TestObject {
	t.obj.sdk = utest.ResetMsg()
	return t
}

//InitChain This is a method of TestObject
func (t *TestObject) InitChain() {
	utest.NextBlock(1)
	t.obj.InitChain()
	utest.Commit()
	return
}

//SetSecretSigner This is a method of TestObject
func (t *TestObject) SetSecretSigner(newSecretSigner types.PubKey) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.SetSecretSigner(newSecretSigner)
	utest.Commit()
	return
}

//SetOwner This is a method of TestObject
func (t *TestObject) SetOwner(newOwnerAddr types.Address) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.SetOwner(newOwnerAddr)
	utest.Commit()
	return
}

//SetSettings This is a method of TestObject
func (t *TestObject) SetSettings(newSettingsStr string) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.SetSettings(newSettingsStr)
	utest.Commit()
	return
}

//SetRecFeeInfo This is a method of TestObject
func (t *TestObject) SetRecFeeInfo(recFeeInfoStr string) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.SetRecFeeInfo(recFeeInfoStr)
	utest.Commit()
	return
}

//WithdrawFunds This is a method of TestObject
func (t *TestObject) WithdrawFunds(tokenName string, beneficiary types.Address, withdrawAmount bn.Number) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.WithdrawFunds(tokenName, beneficiary, withdrawAmount)
	utest.Commit()
	return
}

//PlaceBet This is a method of TestObject
func (t *TestObject) PlaceBet(betInfoJson string, commitLastBlock int64, betIndex string, commit, signData []byte, refAddress types.Address) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.PlaceBet(betInfoJson, commitLastBlock, betIndex, commit, signData, refAddress)
	utest.Commit()
	return
}

//SettleBet This is a method of TestObject
func (t *TestObject) SettleBet(reveal []byte, settleCount int64) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.SettleBet(reveal, settleCount)
	utest.Commit()
	return
}

//WithdrawWin This is a method of TestObject
func (t *TestObject) WithdrawWin(commit []byte) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.WithdrawWin(commit)
	utest.Commit()
	return
}

//RefundBets This is a method of TestObject
func (t *TestObject) RefundBets(commit []byte, refundCount int64) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.RefundBets(commit, refundCount)
	utest.Commit()
	return
}

func (t *TestObject) PaiJuCommon(description string, One, Two, Three int64) {
	sicData := &Sic{}
	sicData.One = One
	sicData.Two = Two
	sicData.Three = Three
	//fmt.Println("----------", description, "--------------")
	//	//fmt.Println("牌九式12=", sicData.PaiJuToIsWin(PAIGOW12))
	//	//fmt.Println("牌九式13=", sicData.PaiJuToIsWin(PAIGOW13))
	//	//fmt.Println("牌九式14=", sicData.PaiJuToIsWin(PAIGOW14))
	//	//fmt.Println("牌九式15=", sicData.PaiJuToIsWin(PAIGOW15))
	//	//fmt.Println("牌九式16=", sicData.PaiJuToIsWin(PAIGOW16))
	//	//fmt.Println("牌九式23=", sicData.PaiJuToIsWin(PAIGOW23))
	//	//fmt.Println("牌九式24=", sicData.PaiJuToIsWin(PAIGOW24))
	//	//fmt.Println("牌九式25=", sicData.PaiJuToIsWin(PAIGOW25))
	//	//fmt.Println("牌九式26=", sicData.PaiJuToIsWin(PAIGOW26))
	//	//fmt.Println("牌九式34=", sicData.PaiJuToIsWin(PAIGOW34))
	//	//fmt.Println("牌九式35=", sicData.PaiJuToIsWin(PAIGOW35))
	//	//fmt.Println("牌九式36=", sicData.PaiJuToIsWin(PAIGOW36))
	//	//fmt.Println("牌九式45=", sicData.PaiJuToIsWin(PAIGOW45))
	//	//fmt.Println("牌九式46=", sicData.PaiJuToIsWin(PAIGOW46))
	//	//fmt.Println("牌九式56=", sicData.PaiJuToIsWin(PAIGOW56))
	fmt.Println("----------", description, "--------------")
	fmt.Println("牌九式12=", t.obj.PaiJuToIsWin(sicData,PAIGOW12))
	fmt.Println("牌九式13=", t.obj.PaiJuToIsWin(sicData,PAIGOW13))
	fmt.Println("牌九式14=", t.obj.PaiJuToIsWin(sicData,PAIGOW14))
	fmt.Println("牌九式15=", t.obj.PaiJuToIsWin(sicData,PAIGOW15))
	fmt.Println("牌九式16=", t.obj.PaiJuToIsWin(sicData,PAIGOW16))
	fmt.Println("牌九式23=", t.obj.PaiJuToIsWin(sicData,PAIGOW23))
	fmt.Println("牌九式24=", t.obj.PaiJuToIsWin(sicData,PAIGOW24))
	fmt.Println("牌九式25=", t.obj.PaiJuToIsWin(sicData,PAIGOW25))
	fmt.Println("牌九式26=", t.obj.PaiJuToIsWin(sicData,PAIGOW26))
	fmt.Println("牌九式34=", t.obj.PaiJuToIsWin(sicData,PAIGOW34))
	fmt.Println("牌九式35=", t.obj.PaiJuToIsWin(sicData,PAIGOW35))
	fmt.Println("牌九式36=", t.obj.PaiJuToIsWin(sicData,PAIGOW36))
	fmt.Println("牌九式45=", t.obj.PaiJuToIsWin(sicData,PAIGOW45))
	fmt.Println("牌九式46=", t.obj.PaiJuToIsWin(sicData,PAIGOW46))
	fmt.Println("牌九式56=", t.obj.PaiJuToIsWin(sicData,PAIGOW56))
}

func (t *TestObject) PointIsWin(description string, One, Two, Three int64) {
	sicData := &Sic{}
	sicData.One = One
	sicData.Two = Two
	sicData.Three = Three
	fmt.Println("开奖结果", sicData)

	isBool1, number1 := t.obj.PointIsWin(sicData,POINT1)
	t.PointIsWinPrint(isBool1, number1, 1)

	isBool2, number2 := t.obj.PointIsWin(sicData,POINT2)
	t.PointIsWinPrint(isBool2, number2, 2)

	isBool3, number3 := t.obj.PointIsWin(sicData,POINT3)
	t.PointIsWinPrint(isBool3, number3, 3)

	isBool4, number4 := t.obj.PointIsWin(sicData,POINT4)
	t.PointIsWinPrint(isBool4, number4, 4)

	isBool5, number5 := t.obj.PointIsWin(sicData,POINT5)
	t.PointIsWinPrint(isBool5, number5, 5)

	isBool6, number6 := t.obj.PointIsWin(sicData,POINT6)
	t.PointIsWinPrint(isBool6, number6, 6)
}

func (t *TestObject) PointIsWinPrint(isBool bool, number int64, point int64) {
	fmt.Println("--------", point, "点-----")
	fmt.Println(isBool, number)
}

func (t *TestObject) SumWin(description string, One, Two, Three int64) {
	sicData := &Sic{}
	sicData.One = One
	sicData.Two = Two
	sicData.Three = Three
	fmt.Println("开奖结果======", sicData)
	fmt.Println("总点数4", t.obj.SumWin(sicData,SUM4))
	fmt.Println("总点数17", t.obj.SumWin(sicData,SUM17))
	fmt.Println("总点数5", t.obj.SumWin(sicData,SUM5))
	fmt.Println("总点数16", t.obj.SumWin(sicData,SUM16))
	fmt.Println("总点数6", t.obj.SumWin(sicData,SUM6))
	fmt.Println("总点数15", t.obj.SumWin(sicData,SUM15))
	fmt.Println("总点数7", t.obj.SumWin(sicData,SUM7))
	fmt.Println("总点数14", t.obj.SumWin(sicData,SUM14))
	fmt.Println("总点数8", t.obj.SumWin(sicData,SUM8))
	fmt.Println("总点数13", t.obj.SumWin(sicData,SUM13))
	fmt.Println("总点数9", t.obj.SumWin(sicData,SUM9))
	fmt.Println("总点数10", t.obj.SumWin(sicData,SUM10))
	fmt.Println("总点数11", t.obj.SumWin(sicData,SUM11))
	fmt.Println("总点数12", t.obj.SumWin(sicData,SUM12))
}

func (t *TestObject) PairToWin(description string, One, Two, Three int64) {
	sicData := &Sic{}
	sicData.One = One
	sicData.Two = Two
	sicData.Three = Three

	fmt.Println("开奖结果======", sicData)

	fmt.Println("一点对子", t.obj.PairToWin(sicData,DOUBLE1))
	fmt.Println("二点对子", t.obj.PairToWin(sicData,DOUBLE2))
	fmt.Println("三点对子", t.obj.PairToWin(sicData,DOUBLE3))
	fmt.Println("四点对子", t.obj.PairToWin(sicData,DOUBLE4))
	fmt.Println("五点对子", t.obj.PairToWin(sicData,DOUBLE5))
	fmt.Println("六点对子", t.obj.PairToWin(sicData,DOUBLE6))
}

func (t *TestObject) WrapSicIsWin(description string, One, Two, Three int64) {
	sicData := &Sic{}
	sicData.One = One
	sicData.Two = Two
	sicData.Three = Three

	fmt.Println("开奖结果======", sicData)

	fmt.Println("一点围骰", t.obj.WrapSicIsWin(sicData,SAMESIC1))
	fmt.Println("二点围骰", t.obj.WrapSicIsWin(sicData,SAMESIC2))
	fmt.Println("三点围骰", t.obj.WrapSicIsWin(sicData,SAMESIC3))
	fmt.Println("四点围骰", t.obj.WrapSicIsWin(sicData,SAMESIC4))
	fmt.Println("五点围骰", t.obj.WrapSicIsWin(sicData,SAMESIC5))
	fmt.Println("六点围骰", t.obj.WrapSicIsWin(sicData,SAMESIC6))

}

func (t *TestObject) SingleIsWin(description string, One, Two, Three int64) {
	sicData := &Sic{}
	sicData.One = One
	sicData.Two = Two
	sicData.Three = Three

	fmt.Println("开奖结果======", sicData)

	fmt.Println("单", t.obj.SingleIsWin(sicData,ODD))
	fmt.Println("双", t.obj.SingleIsWin(sicData,EVEN))
}

func (t *TestObject) BigSmallIsWin(description string, One, Two, Three int64) {
	sicData := &Sic{}
	sicData.One = One
	sicData.Two = Two
	sicData.Three = Three

	fmt.Println("开奖结果======", sicData)

	fmt.Println("小", t.obj.BigSmallIsWin(sicData,SMALL))
	fmt.Println("大", t.obj.BigSmallIsWin(sicData,BIG))
}

func (t *TestObject) AllSameSic(description string, One, Two, Three int64) {
	sicData := &Sic{}
	sicData.One = One
	sicData.Two = Two
	sicData.Three = Three

	fmt.Println("开奖结果======", sicData)

	fmt.Println("全围", t.obj.WrapSicIsWin(sicData,ALLSAMESIC))
}




func (t *TestObject) SetBigOrSmall(num int64) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.SetBigOrSmall(num)
	utest.Commit()
	return
}
