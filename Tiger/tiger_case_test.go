package tiger

import (
	"blockchain/smcsdk/sdk"
	"blockchain/algorithm"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/utest"
	"common/kms"
	"encoding/hex"
	"common/keys"
	"fmt"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/go-crypto"
	"github.com/tendermint/tmlibs/common"
	"gopkg.in/check.v1"
	"io/ioutil"
	"math"
	"testing"
)

const (
	ownerName = "local_owner"
	password  = "12345678"
)

var (
	cdc = amino.NewCodec()
)

func init() {
	crypto.RegisterAmino(cdc)
	crypto.SetChainId("local")
	kms.InitKMS("./.keystore", "local_mode", "", "", "0x1003")
	kms.GenPrivKey(ownerName, []byte(password))
}

//Test This is a function
func Test(t *testing.T) { check.TestingT(t) }

//MySuite This is a struct
type MySuite struct{}

var _ = check.Suite(&MySuite{})

//TestSicBo_SetSecretSigner This is a method of MySuite
func (mysuit *MySuite) TestTiger_SetSecretSigner(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	pubKey, _ := kms.GetPubKey(ownerName, []byte(password))

	account := utest.NewAccount(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1000000000))

	var tests = []struct {
		account sdk.IAccount
		pubKey  []byte
		desc    string
		code    uint32
	}{
		{contractOwner, pubKey, "--正常流程--", types.CodeOK},
		{contractOwner, []byte("0xff"), "--异常流程--公钥长度不正确--", types.ErrInvalidParameter},
		{account, pubKey, "--异常流程--非owner调用--", types.ErrNoAuthorization},
	}

	for _, item := range tests {
		utest.AssertError(test.run().setSender(item.account).SetSecretSigner(item.pubKey), item.code)
	}
}

//TestSicBo_SetSecretSigner This is a method of MySuite
func (mysuit *MySuite) TestSicBo_SetOwner(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	//pubKey, _ := kms.GetPubKey(ownerName, []byte(password))
	//
	account := utest.NewAccount(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1000000000))
	ss := account.Address()
	fmt.Println(ss)
	var tests = []struct {
		account sdk.IAccount
		addr    types.Address
		desc    string
		code    uint32
	}{
		{contractOwner, types.Address("local5wNLBx3qhuQoSuA5ZGsyw8Fj6dJGwdFww"), "--正常流程--", types.CodeOK},
		{contractOwner, types.Address("local5wNLBx3qhuQoSuA5ZGsyw8Fj6dJGwdFwwrrr"), "--异常流程--公钥长度不正确--", types.ErrInvalidParameter},
		{account, types.Address("local5wNLBx3qhuQoSuA5ZGsyw8Fj6dJGwdFwr"), "--异常流程--非owner调用--", types.ErrNoAuthorization},
		//{account, types.Address("local5wNLBx3qhuQoSuA5ZGsyw8Fj6dJGwdFwr"), "--异常流程--非owner调用--", types.ErrInvalidParameter},
	}

	//utest.AssertError(test.run().setSender(tests[0].account).SetOwner(tests[0].addr), tests[0].code)
	//utest.AssertError(test.run().setSender(tests[1].account).SetOwner(tests[1].addr), tests[1].code)
	utest.AssertError(test.run().setSender(tests[2].account).SetOwner(tests[2].addr), tests[2].code)

}

//TestSicBo_SetSettings This is a method of MySuite
func (mysuit *MySuite) TestSicBo_SetSettings(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 1)
	if accounts == nil {
		panic("初始化newOwner失败")
	}

	settings := Settings{}

	settings.FeeRatio = 50
	settings.FeeMiniNum = 300000
	settings.SendToCltRatio = 100
	settings.BetExpirationBlocks = 250
	settings.TokenNames = []string{test.obj.sdk.Helper().GenesisHelper().Token().Name()}
	limitMaps := make(map[string]Limit, len(settings.TokenNames))
	for _, value := range settings.TokenNames {
		limit := Limit{
			MaxProfit: 2E12,
			MaxLimit:  2E10,
			MinLimit:  1E8,
		}
		limitMaps[value] = limit
	}
	settings.LimitMaps = limitMaps
	resBytes1, _ := jsoniter.Marshal(settings)

	for _, value := range settings.TokenNames {
		limit := Limit{
			MaxProfit: 2E12,
			MaxLimit:  2E9,
			MinLimit:  2E10,
		}
		limitMaps[value] = limit
	}
	settings.LimitMaps = limitMaps
	resBytes2, _ := jsoniter.Marshal(settings)

	for _, value := range settings.TokenNames {
		limit := Limit{
			MaxProfit: 2E12,
			MaxLimit:  2E10,
			MinLimit:  2E8,
		}
		limitMaps[value] = limit
	}
	settings.LimitMaps = limitMaps
	settings.TokenNames = []string{}
	resBytes3, _ := jsoniter.Marshal(settings)

	settings.TokenNames = []string{test.obj.sdk.Helper().GenesisHelper().Token().Name()}
	for _, value := range settings.TokenNames {
		limit := Limit{
			MaxProfit: 2E12,
			MaxLimit:  0,
			MinLimit:  2E10,
		}
		limitMaps[value] = limit
	}
	settings.LimitMaps = limitMaps
	resBytes4, _ := jsoniter.Marshal(settings)

	for _, value := range settings.TokenNames {
		limit := Limit{
			MaxProfit: 2E12,
			MaxLimit:  2E10,
			MinLimit:  -1,
		}
		limitMaps[value] = limit
	}
	settings.LimitMaps = limitMaps
	resBytes5, _ := jsoniter.Marshal(settings)

	for _, value := range settings.TokenNames {
		limit := Limit{
			MaxProfit: math.MinInt64,
			MaxLimit:  2E8,
			MinLimit:  2E10,
		}
		limitMaps[value] = limit
	}
	settings.LimitMaps = limitMaps
	resBytes6, _ := jsoniter.Marshal(settings)

	for _, value := range settings.TokenNames {
		limit := Limit{
			MaxProfit: 2E12,
			MaxLimit:  2E9,
			MinLimit:  2E10,
		}
		limitMaps[value] = limit
	}
	settings.LimitMaps = limitMaps
	settings.FeeMiniNum = -1
	resBytes7, _ := jsoniter.Marshal(settings)

	settings.FeeMiniNum = 300000
	settings.FeeRatio = -1
	resBytes8, _ := jsoniter.Marshal(settings)

	settings.FeeRatio = 1001
	resBytes9, _ := jsoniter.Marshal(settings)

	settings.FeeRatio = 50
	settings.SendToCltRatio = -1
	resBytes10, _ := jsoniter.Marshal(settings)

	settings.SendToCltRatio = 1001
	resBytes11, _ := jsoniter.Marshal(settings)

	settings.SendToCltRatio = 100
	settings.BetExpirationBlocks = -1
	resBytes12, _ := jsoniter.Marshal(settings)

	var tests = []struct {
		account  sdk.IAccount
		settings []byte
		desc     string
		code     uint32
	}{
		{contractOwner, resBytes1, "--正常流程--", types.CodeOK},
		{contractOwner, resBytes2, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes3, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes4, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes5, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes6, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes7, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes8, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes9, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes10, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes11, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes12, "--异常流程--", types.ErrInvalidParameter},
		{accounts[0], resBytes1, "--异常流程--", types.ErrNoAuthorization},
	}

	test.run().setSender(contractOwner).InitChain()
	for _, item := range tests {
		utest.AssertError(test.run().setSender(item.account).SetSettings(string(item.settings)), item.code)
	}
}

//TestSicBo_SetRecFeeInfo This is a method of MySuite
func (mysuit *MySuite) TestSicBo_SetRecFeeInfo(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 1)
	if accounts == nil {
		panic("初始化newOwner失败")
	}

	recFeeInfo := make([]RecFeeInfo, 0)
	resBytes2, _ := jsoniter.Marshal(recFeeInfo)
	item := RecFeeInfo{
		RecFeeRatio: 500,
		RecFeeAddr:  "local9ge366rtqV9BHqNwn7fFgA8XbDQmJGZqE",
	}
	recFeeInfo = append(recFeeInfo, item)
	resBytes1, _ := jsoniter.Marshal(recFeeInfo)

	item1 := RecFeeInfo{
		RecFeeRatio: 501,
		RecFeeAddr:  "local9ge366rtqV9BHqNwn7fFgA8XbDQmJGZqE",
	}
	recFeeInfo = append(recFeeInfo, item1)
	resBytes3, _ := jsoniter.Marshal(recFeeInfo)

	recFeeInfo = append(recFeeInfo[:1], recFeeInfo[2:]...)
	item2 := RecFeeInfo{
		RecFeeRatio: 450,
		RecFeeAddr:  "lo9ge366rtqV9BHqNwn7fFgA8XbDQmJGZqE",
	}
	recFeeInfo = append(recFeeInfo, item2)
	resBytes4, _ := jsoniter.Marshal(recFeeInfo)

	recFeeInfo = append(recFeeInfo[:1], recFeeInfo[2:]...)
	item3 := RecFeeInfo{
		RecFeeRatio: 500,
		RecFeeAddr:  test.obj.sdk.Helper().BlockChainHelper().CalcAccountFromName(contractName,""),
	}
	recFeeInfo = append(recFeeInfo, item3)
	//resBytes5, _ := jsoniter.Marshal(recFeeInfo)

	recFeeInfo = append(recFeeInfo[:1], recFeeInfo[2:]...)
	item4 := RecFeeInfo{
		RecFeeRatio: -1,
		RecFeeAddr:  "local9ge366rtqV9BHqNwn7fFgA8XbDQmJGZqE",
	}
	recFeeInfo = append(recFeeInfo, item4)
	resBytes6, _ := jsoniter.Marshal(recFeeInfo)

	var tests = []struct {
		account sdk.IAccount
		infos   []byte
		desc    string
		code    uint32
	}{

		{contractOwner, resBytes1, "--正常流程--", types.CodeOK},
		{contractOwner, resBytes2, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes3, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes4, "--异常流程--", types.ErrInvalidAddress},
		//{contractOwner, resBytes5, "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, resBytes6, "--异常流程--", types.ErrInvalidParameter},
		{accounts[0], resBytes1, "--异常流程--", types.ErrNoAuthorization},
	}

	for _, item := range tests {
		utest.AssertError(test.run().setSender(item.account).SetRecFeeInfo(string(item.infos)), item.code)
	}
}


//TestTiger_ClacByLine This is a method of MySuite
func (mysuit *MySuite) TestTiger_ClacByLine(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_ClacMul This is a method of MySuite
func (mysuit *MySuite) TestTiger_ClacMul(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_ClacFee This is a method of MySuite
func (mysuit *MySuite) TestTiger_ClacFee(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_GetNineLine This is a method of MySuite
func (mysuit *MySuite) TestTiger_GetNineLine(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}





//TestTiger_DigtalCurrency This is a method of MySuite
func (mysuit *MySuite) TestTiger_DigtalCurrency(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	//test.setSender(contractOwner).InitChain()
	//TODO
	contract := utest.UTP.Message().Contract()
	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)

	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), contract.Account(), bn.N(1E11))
	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 6)
	if accounts == nil {
		panic("初始化newOwner失败")
	}
	_, pubKey, _, _, _, _ := PlaceBetHelper(104)
	test.run().setSender(contractOwner).InitChain()
	utest.AssertError(test.run().setSender(contractOwner).SetSecretSigner(pubKey[:]), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[5]).transfer(bn.N(1000000009)).DigtalCurrency("LOC",1000000009), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[5]).transfer(bn.N(1000000002)).DigtalCurrency("LOC",1000000002), types.CodeOK)


}

//TestTiger_ShuffleMainCards This is a method of MySuite
func (mysuit *MySuite) TestTiger_ShuffleMainCards(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_ShuffleRound This is a method of MySuite
func (mysuit *MySuite) TestTiger_ShuffleRound(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_AssemblePoker This is a method of MySuite
func (mysuit *MySuite) TestTiger_AssemblePoker(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_ShuffleFeeCards This is a method of MySuite
func (mysuit *MySuite) TestTiger_ShuffleFeeCards(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_PlaceBet This is a method of MySuite
func (mysuit *MySuite) TestTiger_PlaceBet(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	//test.setSender(contractOwner).InitChain()
	//TODO
	contract := utest.UTP.Message().Contract()
	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)

	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), contract.Account(), bn.N(1E11))
	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 6)
	if accounts == nil {
		panic("初始化newOwner失败")
	}
	// DigtalCurrency(tk string, num int64)
	test.run().setSender(contractOwner).InitChain()


	_, _, reval, _, _, _ := PlaceBetHelper(100)
	_, _, reval1, _, _, _ := PlaceBetHelper(101)
	_, _, reval2, _, _, _ := PlaceBetHelper(102)
	_, _, reval3, _, _, _ := PlaceBetHelper(103)
	commitLastBlock, pubKey, reval4, commit, signData, _ := PlaceBetHelper(104)
	utest.AssertError(test.run().setSender(contractOwner).SetSecretSigner(pubKey[:]), types.CodeOK)
	//utest.AssertError(err, types.CodeOK)
	reveal:= [][]byte{reval,reval1,reval2,reval3,reval4}

	utest.AssertError(test.run().setSender(accounts[5]).transfer(bn.N(1000000009000)).DigtalCurrency("LOC",1000000009000), types.CodeOK)
	test.run().setSender(accounts[5]).transfer(bn.N(1000000001)).PlaceBet(reveal,"LOC", 100000001, commitLastBlock,commit, signData[:], "")

for i:=1;i<1000;i++{
	s1,s2,s3,s4,s5:= i+1,i+2,i+3,i+4,i+5
	_, _, reval, _, _, _ := PlaceBetHelper(int64(s1))
	_, _, reval1, _, _, _ := PlaceBetHelper(int64(s2))
	_, _, reval2, _, _, _ := PlaceBetHelper(int64(s3))
	_, _, reval3, _, _, _ := PlaceBetHelper((int64(s4)))
	commitLastBlock, _, reval4, commit, signData, _ := PlaceBetHelper(int64(s5))
	reveal:= [][]byte{reval,reval1,reval2,reval3,reval4}
	test.run().setSender(accounts[5]).PlaceBet(reveal,"LOC", 1000000001, commitLastBlock,commit, signData[:], "")



}


}

//TestTiger_PlaceFeeBet This is a method of MySuite
//func (mysuit *MySuite) TestTiger_PlaceFeeBet(c *check.C) () {
//	utest.Init(orgID)
//	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
//	test := NewTestObject(contractOwner)
//	//test.setSender(contractOwner).InitChain()
//	//TODO
//	contract := utest.UTP.Message().Contract()
//	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
//	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)
//
//	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), contract.Account(), bn.N(1E11))
//	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 6)
//	if accounts == nil {
//		panic("初始化newOwner失败")
//	}
//
//	_, _, reval, _, _, _ := PlaceBetHelper(100)
//	_, _, reval1, _, _, _ := PlaceBetHelper(101)
//	_, _, reval2, _, _, _ := PlaceBetHelper(102)
//	_, _, reval3, _, _, _ := PlaceBetHelper(103)
//	commitLastBlock, pubKey, reval4, commit, signData, _ := PlaceBetHelper(104)
//	//utest.AssertError(err, types.CodeOK)
//	reveal:= [][]byte{reval,reval1,reval2,reval3,reval4}
//	// DigtalCurrency(tk string, num int64)
//	test.run().setSender(contractOwner).InitChain()
//	utest.AssertError(test.run().setSender(contractOwner).SetSecretSigner(pubKey[:]), types.CodeOK)
//	utest.AssertError(test.run().setSender(accounts[5]).transfer(bn.N(10000000000)).DigtalCurrency("LOC",10000000000), types.CodeOK)
//	utest.AssertError(test.run().setSender(accounts[5]).transfer(bn.N(1000000001)).PlaceBet(reveal,"LOC", 100000001, commitLastBlock,commit, signData[:], ""), types.CodeOK)
//	utest.AssertError(test.run().setSender(accounts[5]).PlaceFeeBet(reveal,"LOC", 100000001, commitLastBlock,commit, signData[:], ""), types.CodeOK)
//
//}

//TestTiger_GetRandomNum This is a method of MySuite
func (mysuit *MySuite) TestTiger_GetRandomNum(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}


//hempHeight 想对于下注高度和生效高度之间的差值
//acct 合约的owner
func PlaceBetHelper(tempHeight int64) (commitLastBlock int64, pubKey [32]byte, reveal, commit []byte, signData [64]byte, err types.Error) {
	acct, err := Load("./.keystore/local_owner.wal", []byte(password), nil)
	if err.ErrorCode != types.CodeOK {
		return
	}

	localBlockHeight := utest.UTP.ISmartContract.Block().Height()

	pubKey = acct.PubKey.(crypto.PubKeyEd25519)

	commitLastBlock = localBlockHeight + tempHeight
	decode := crypto.CRandBytes(32)
	revealStr := hex.EncodeToString(algorithm.SHA3256(decode))
	reveal, _ = hex.DecodeString(revealStr)

	commit = algorithm.SHA3256(reveal)

	signByte := append(bn.N(commitLastBlock).Bytes(), commit...)
	signData = acct.PrivKey.Sign(signByte).(crypto.SignatureEd25519)

	return
}

func Load(keystorePath string, password, fingerprint []byte) (acct *keys.Account, err types.Error) {
	if keystorePath == "" {
		common.PanicSanity("Cannot loads account because keystorePath not set")
	}

	walBytes, mErr := ioutil.ReadFile(keystorePath)
	if mErr != nil {
		err.ErrorCode = types.ErrInvalidParameter
		err.ErrorDesc = "account does not exist"
		return
	}

	jsonBytes, mErr := algorithm.DecryptWithPassword(walBytes, password, fingerprint)
	if mErr != nil {
		err.ErrorCode = types.ErrInvalidParameter
		err.ErrorDesc = fmt.Sprintf("the password is wrong err info : %s", mErr)
		return
	}

	acct = new(keys.Account)
	mErr = cdc.UnmarshalJSON(jsonBytes, acct)
	if mErr != nil {
		err.ErrorCode = types.ErrInvalidParameter
		err.ErrorDesc = fmt.Sprintf("UnmarshalJSON is wrong err info : %s", mErr)
		return
	}

	acct.KeystorePath = keystorePath
	err.ErrorCode = types.CodeOK
	return
}

////TestTiger_SetSecretSigner is a method of MySuite
//func (mysuit *MySuite) TestTiger_SetSecretSigner(c *check.C) () {
//	utest.Init(orgID)
//	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
//	test := NewTestObject(contractOwner)
//	test.setSender(contractOwner).InitChain()
//	//TODO
//}
//
////TestTiger_SetOwner is a method of MySuite
//func (mysuit *MySuite) TestTiger_SetOwner(c *check.C) () {
//	utest.Init(orgID)
//	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
//	test := NewTestObject(contractOwner)
//	test.setSender(contractOwner).InitChain()
//	//TODO
//}
//
////TestTiger_SetSettings is a method of MySuite
//func (mysuit *MySuite) TestTiger_SetSettings(c *check.C) () {
//	utest.Init(orgID)
//	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
//	test := NewTestObject(contractOwner)
//	test.setSender(contractOwner).InitChain()
//	//TODO
//}
//TestSicBo_SetRecFeeInfo
//TestTiger_SetRecFeeInfo is a method of MySuite
//func (mysuit *MySuite) TestTiger_SetRecFeeInfo(c *check.C) () {
//	utest.Init(orgID)
//	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
//	test := NewTestObject(contractOwner)
//	test.setSender(contractOwner).InitChain()
//	//TODO
//}
//TestTiger_SetOwner is a method of MySuite
func (mysuit *MySuite) TestTiger_SetOwner(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_SetSettings is a method of MySuite
func (mysuit *MySuite) TestTiger_SetSettings(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_SetRecFeeInfo is a method of MySuite
func (mysuit *MySuite) TestTiger_SetRecFeeInfo(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_SetPoker is a method of MySuite
func (mysuit *MySuite) TestTiger_SetPoker(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_WithdrawFunds is a method of MySuite
func (mysuit *MySuite) TestTiger_WithdrawFunds(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}

//TestTiger_PlaceFeeBet is a method of MySuite
func (mysuit *MySuite) TestTiger_PlaceFeeBet(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	test.setSender(contractOwner).InitChain()
	//TODO
}