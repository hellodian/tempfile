package tiger

import (
	"fmt"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/utest"
	"blockchain/smcsdk/sdkimpl/sdkhelper"
	"blockchain/smcsdk/sdkimpl/object"
	"blockchain/smcsdk/sdk/types"
)

var (
	contractName= "Tiger" //contract name
	contractMethods= []string{"ClacByLine([]Poker)(TYPE,TYPE,ARRAY_OR_SLICE_TYPE)", "ClacMul(int64,int64)TYPE", "ClacFee([]Poker)TYPE", "GetNineLine(types.Address)(TYPE,TYPE,ARRAY_OR_SLICE_TYPE)", "SetPoker(string,string)", "SetSecretSigner(types.PubKey)", "SetOwner(types.Address)", "SetSettings(string)", "SetRecFeeInfo(string)", "DigtalCurrency(string,int64)", "WithdrawFunds(string,types.Address,bn.Number)", "PlaceBet([]byte,string,int64,int64,[]byte,[]byte,types.Address)", "PlaceFeeBet([]byte,string,int64,int64,[]byte,[]byte,types.Address)", "GetRandomNum([]byte)TYPE"}
	contractInterfaces= []string{}
	orgID= "orgNUjCm1i8RcoW2kVTbDw4vKW6jzfMxewJHjkhuiduhjuikjuyhnnjkuhujk111"
)
//TestObject This is a struct for test
type TestObject struct {
	obj *Tiger
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
	return &TestObject{&Tiger{sdk: utest.UTP.ISmartContract}}
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

//ClacByLine This is a method of TestObject
func (t *TestObject) ClacByLine(p []Poker) (result0 int64, result1 int64, result2 []Poker, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0, result1, result2 = t.obj.ClacByLine(p)
	utest.Commit()
	return
}

//ClacMul This is a method of TestObject
func (t *TestObject) ClacMul(mul, sum int64) (result0 int64, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.ClacMul(mul, sum)
	utest.Commit()
	return
}

//ClacFee This is a method of TestObject
func (t *TestObject) ClacFee(p []Poker) (result0 int64, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.ClacFee(p)
	utest.Commit()
	return
}

//GetNineLine This is a method of TestObject
func (t *TestObject) GetNineLine(address types.Address) (result0 int64, result1 int64, result2 [][]Poker, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0, result1, result2 = t.obj.GetNineLine(address)
	utest.Commit()
	return
}

//SetPoker This is a method of TestObject
func (t *TestObject) SetPoker(main, fee string) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.SetPoker(main, fee)
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

//DigtalCurrency This is a method of TestObject
func (t *TestObject) DigtalCurrency(tk string, num int64) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.DigtalCurrency(tk, num)
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
func (t *TestObject) PlaceBet(reveals []byte, tokenName string, betNum, commitLastBlock int64, commit, signData []byte, refAddress types.Address) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.PlaceBet(reveals, tokenName, betNum, commitLastBlock, commit, signData, refAddress)
	utest.Commit()
	return
}

//PlaceFeeBet This is a method of TestObject
func (t *TestObject) PlaceFeeBet(reveals []byte, tokenName string, betNum, commitLastBlock int64, commit, signData []byte, refAddress types.Address) (err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	t.obj.PlaceFeeBet(reveals, tokenName, betNum, commitLastBlock, commit, signData, refAddress)
	utest.Commit()
	return
}

//GetRandomNum This is a method of TestObject
func (t *TestObject) GetRandomNum(reveal []byte) (result0 bn.Number, err types.Error) {
	err.ErrorCode = types.CodeOK
	defer FuncRecover(&err)
	utest.NextBlock(1)
	result0 = t.obj.GetRandomNum(reveal)
	utest.Commit()
	return
}
            

