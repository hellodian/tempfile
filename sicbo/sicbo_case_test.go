package sicbo

import (
	"blockchain/algorithm"
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/jsoniter"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/utest"
	"common/keys"
	"common/kms"
	"encoding/hex"
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
func (mysuit *MySuite) TestSicBo_SetOwner(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	//pubKey, _ := kms.GetPubKey(ownerName, []byte(password))
	//
	account := utest.NewAccount(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1000000000))
	ss:=account.Address()
	fmt.Println(ss)
	var tests = []struct {
		account sdk.IAccount
		addr  types.Address
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


//TestSicBo_SetSecretSigner This is a method of MySuite
func (mysuit *MySuite) TestSicBo_SetSecretSigner(c *check.C) {
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

	settings := SBSettings{}

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

//TestSicBo_WithdrawFunds This is a method of MySuite
func (mysuit *MySuite) TestSicBo_WithdrawFunds(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	genesisToken := test.obj.sdk.Helper().GenesisHelper().Token()
	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
	contractAccount := utest.UTP.Helper().ContractHelper().ContractOfName(contractName).Account()

	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)

	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), contractAccount, bn.N(1E11))
	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 1)
	if accounts == nil {
		panic("初始化newOwner失败")
	}

	test.run().setSender(contractOwner).InitChain()

	var tests = []struct {
		account        sdk.IAccount
		tokenName      string
		beneficiary    types.Address
		withdrawAmount bn.Number
		desc           string
		code           uint32
	}{
		{contractOwner, genesisToken.Name(), contractOwner.Address(), bn.N(1E10), "--正常流程--", types.CodeOK},
		{contractOwner, genesisToken.Name(), accounts[0].Address(), bn.N(1E10), "--正常流程--", types.CodeOK},
		{contractOwner, genesisToken.Name(), contractOwner.Address(), bn.N(1E15), "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, genesisToken.Name(), contractOwner.Address(), bn.N(-1), "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, genesisToken.Name(), contractAccount, bn.N(1E10), "--异常流程--", types.ErrInvalidParameter},
		{contractOwner, "xt", contractOwner.Address(), bn.N(1E10), "--异常流程--", types.ErrInvalidParameter},
		{accounts[0], genesisToken.Name(), contractOwner.Address(), bn.N(1E10), "--异常流程--", types.ErrNoAuthorization},
	}

	for _, item := range tests {
		utest.AssertError(test.run().setSender(item.account).WithdrawFunds(item.tokenName, item.beneficiary, item.withdrawAmount), item.code)
	}
}

//TestSicBo_PlaceBet This is a method of MySuite
func (mysuit *MySuite) TestSicBo_PlaceBet(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	contract := utest.UTP.Message().Contract()
	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)

	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), contract.Account(), bn.N(1E11))
	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 5)
	if accounts == nil {
		panic("初始化newOwner失败")
	}

	commitLastBlock, pubKey, _, commit, signData, _ := PlaceBetHelper(100)
	//utest.AssertError(err, types.CodeOK)

	test.run().setSender(contractOwner).InitChain()
	utest.AssertError(test.run().setSender(contractOwner).SetSecretSigner(pubKey[:]), types.CodeOK)

	betData := []BetData{{1, bn.N(1000000000)}}
	betDataJsonBytes, _ := jsoniter.Marshal(betData)
	betData1 := []BetData{{2, bn.N(1000000000)}}
	betDataJsonBytes1, _ := jsoniter.Marshal(betData1)
	// PlaceBet(betInfoJson string, commitLastBlock int64,betIndex string, commit, signData []byte, refAddress types.Address)
	utest.AssertError(test.run().setSender(accounts[0]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes), commitLastBlock, "hello", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[0]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgd", commit, signData[:], ""), types.CodeOK)
	//	i++
	//}
}

//TestSicBo_SettleBet This is a method of MySuite
func (mysuit *MySuite) TestSicBo_SettleBet(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)
	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), test.obj.sdk.Message().Contract().Account(), bn.N(1E11))
	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 5)
	if accounts == nil {
		panic("初始化newOwner失败")
	}

	commitLastBlock, pubKey, reveal, commit, signData, _ := PlaceBetHelper(100)
	//utest.AssertError(err, types.CodeOK)

	test.run().setSender(contractOwner).InitChain()
	utest.AssertError(test.run().setSender(contractOwner).SetSecretSigner(pubKey[:]), types.CodeOK)
	//var i int64 = 0
	//for {
	//	if i > 2 {
	//		break
	//	}
	//	betData := []BetData{{i + 1, bn.N(1000000000)}}
	//	betDataJsonBytes, _ := jsoniter.Marshal(betData)
	//	//utest.AssertError(test.run().setSender(accounts[i]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes), commitLastBlock, commit, signData[:], ""), types.CodeOK)
	//	utest.AssertError(test.run().setSender(accounts[i]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes), commitLastBlock,"hello", commit, signData[:], ""), types.CodeOK)
	//	i++
	//
	//}
	betData := []BetData{{1, bn.N(1000000000)}}
	betDataJsonBytes, _ := jsoniter.Marshal(betData)
	betData1 := []BetData{{2, bn.N(1000000000)}}
	betDataJsonBytes1, _ := jsoniter.Marshal(betData1)
	// PlaceBet(betInfoJson string, commitLastBlock int64,betIndex string, commit, signData []byte, refAddress types.Address)
	utest.AssertError(test.run().setSender(accounts[0]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes), commitLastBlock, "hello", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[0]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgd", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[1]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgdf", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[2]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgdd", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[3]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgdw", commit, signData[:], ""), types.CodeOK)
	//	i++
	utest.AssertError(test.run().setSender(contractOwner).SettleBet(reveal, 1), types.CodeOK)
	utest.AssertError(test.run().setSender(contractOwner).SettleBet(reveal, 3), types.CodeOK)
}

//TestSicBo_RefundBets This is a method of MySuite
func (mysuit *MySuite) TestSicBo_RefundBets(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)
	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), test.obj.sdk.Message().Contract().Account(), bn.N(1E11))
	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 2)
	if accounts == nil {
		panic("初始化newOwner失败")
	}

	commitLastBlock, pubKey, _, commit, signData, _ := PlaceBetHelper(100)

	test.run().setSender(contractOwner).InitChain()
	utest.AssertError(test.run().setSender(contractOwner).SetSecretSigner(pubKey[:]), types.CodeOK)

	betData := []BetData{{1, bn.N(1000000000)}}
	betDataJsonBytes, _ := jsoniter.Marshal(betData)
	utest.AssertError(test.run().setSender(accounts[0]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes), commitLastBlock, "hiii", commit, signData[:], ""), types.CodeOK)
	betData = []BetData{{2, bn.N(1000000000)}}
	utest.AssertError(test.run().setSender(accounts[1]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes), commitLastBlock, "gggg", commit, signData[:], ""), types.CodeOK)
	// set bet time out
	count := 0
	for {
		utest.NextBlock(1)
		count++
		if count > 250 {
			break
		}
	}
	utest.AssertError(test.run().setSender(contractOwner).RefundBets(commit, 1), types.CodeOK)
	utest.AssertError(test.run().setSender(contractOwner).RefundBets(commit, 1), types.CodeOK)
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

//TestSicBo_WithdrawWin is a method of MySuite
func (mysuit *MySuite) TestSicBo_WithdrawWin(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)
	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), test.obj.sdk.Message().Contract().Account(), bn.N(1E11))
	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 5)
	if accounts == nil {
		panic("初始化newOwner失败")
	}
	//commitLastBlock, pubKey, reveal, commit, signData, _ := PlaceBetHelper(100)
	commitLastBlock, pubKey, reveal, commit, signData, _ := PlaceBetHelper(100)
	test.run().setSender(contractOwner).InitChain()
	utest.AssertError(test.run().setSender(contractOwner).SetSecretSigner(pubKey[:]), types.CodeOK)
	betData := []BetData{{1, bn.N(1000000000)}}
	betDataJsonBytes, _ := jsoniter.Marshal(betData)
	betData1 := []BetData{{2, bn.N(1000000000)}}
	betDataJsonBytes1, _ := jsoniter.Marshal(betData1)
	utest.AssertError(test.run().setSender(accounts[0]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes), commitLastBlock, "hello", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[0]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgd", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[1]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgdf", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[2]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgdd", commit, signData[:], ""), types.CodeOK)
	utest.AssertError(test.run().setSender(accounts[3]).transfer(bn.N(1000000000)).PlaceBet(string(betDataJsonBytes1), commitLastBlock, "hhgdw", commit, signData[:], ""), types.CodeOK)

	//utest.AssertError(test.run().setSender(contractOwner).SettleBet(reveal, 1), types.CodeOK)
	//utest.AssertError(test.run().setSender(contractOwner).SettleBet(reveal, 3), types.CodeOK)
	utest.AssertError(test.run().setSender(contractOwner).SettleBet(reveal, 1), types.CodeOK)
	//utest.AssertError(test.run().setSender(contractOwner).WithdrawWin(commit), types.ErrInvalidParameter)
	utest.AssertError(test.run().setSender(accounts[3]).WithdrawWin(commit), types.CodeOK)
}

//Test PaiJu
func (mysuit *MySuite) TestSicBo_PaiJu(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	//牌九式12
	description12 := "牌九式12 点数 one=1 two=2 three=3 "
	test.run().setSender(contractOwner).PaiJuCommon(description12, 1, 2, 3)
	//牌九式13
	description13 := "牌九式13 点数 one=1 two=2 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description13, 1, 2, 3)
	//牌九式14
	description14 := "牌九式14 点数 one=1 two=2 three=4"
	test.run().setSender(contractOwner).PaiJuCommon(description14, 1, 2, 4)
	//牌九式15
	description15 := "牌九式15 点数 one=2 two=1 three=5"
	test.run().setSender(contractOwner).PaiJuCommon(description15, 2, 1, 5)
	//牌九式16
	description16 := "牌九式16 点数 one=1 two=4 three=6"
	test.run().setSender(contractOwner).PaiJuCommon(description16, 1, 4, 6)
	//牌九式23
	description23 := "牌九式23 点数 one=1 two=2 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description23, 1, 2, 3)
	//牌九式24
	description24 := "牌九式24 点数 one=4 two=2 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description24, 4, 2, 3)
	//牌九式25
	description25 := "牌九式25 点数 one=5 two=2 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description25, 5, 2, 3)
	//牌九式26
	description26 := "牌九式26 点数 one=6 two=2 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description26, 6, 2, 3)
	//牌九式34
	description34 := "牌九式34 点数 one=4 two=2 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description34, 4, 2, 3)
	//牌九式35
	description35 := "牌九式35 点数 one=1 two=5 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description35, 1, 5, 3)
	//牌九式36
	description36 := "牌九式36 点数 one=1 two=6 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description36, 1, 6, 3)
	//牌九式45
	description45 := "牌九式45 点数 one=4 two=5 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description45, 4, 5, 3)
	//牌九式46
	description46 := "牌九式46 点数 one=6 two=4 three=3"
	test.run().setSender(contractOwner).PaiJuCommon(description46, 6, 4, 3)
	//牌九式56
	description56 := "牌九式56 点数 one=5 two=2 three=6"
	test.run().setSender(contractOwner).PaiJuCommon(description56, 5, 2, 6)
}

//Test PointIsWin
func (mysuit *MySuite) TestSicBo_PointIsWin(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	description1 := "1点 点数 one=1 two=2 three=6"
	test.run().setSender(contractOwner).PointIsWin(description1, 1, 2, 6)

	description2 := "2点 点数 one=1 two=2 three=6"
	test.run().setSender(contractOwner).PointIsWin(description2, 1, 2, 6)

	description3 := "3点 点数 one=1 two=2 three=3"
	test.run().setSender(contractOwner).PointIsWin(description3, 1, 2, 3)

	description4 := "4点 点数 one=4 two=2 three=6"
	test.run().setSender(contractOwner).PointIsWin(description4, 4, 2, 6)

	description5 := "5点 点数 one=1 two=2 three=5"
	test.run().setSender(contractOwner).PointIsWin(description5, 1, 2, 5)

	description6 := "6点 点数 one=6 two=2 three=6"
	test.run().setSender(contractOwner).PointIsWin(description6, 6, 2, 6)
}

//Test SumWin
func (mysuit *MySuite) TestSicBo_SumWin(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	description1 := "总点数 = 4点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description1, 1, 2, 1)

	description2 := "总点数 = 17点 点数 one=6 two=6 three=5"
	test.run().setSender(contractOwner).SumWin(description2, 6, 6, 5)

	description3 := "总点数 = 5点 点数 one=1 two=2 three=2"
	test.run().setSender(contractOwner).SumWin(description3, 1, 2, 2)

	description4 := "总点数 = 16点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description4, 6, 4, 6)

	description5 := "总点数 = 6点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description5, 1, 2, 3)

	description6 := "总点数 = 15点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description6, 5, 5, 5)

	description7 := "总点数 = 7点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description7, 1, 1, 5)

	description8 := "总点数 = 14点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description8, 4, 4, 6)

	description9 := "总点数 = 8点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description9, 1, 2, 5)

	description10 := "总点数 = 13点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description10, 1, 6, 6)

	description11 := "总点数 = 9点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description11, 1, 2, 6)

	description12 := "总点数 = 10点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description12, 3, 4, 3)

	description13 := "总点数 = 11点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description13, 4, 4, 3)

	description14 := "总点数 = 12点 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).SumWin(description14, 4, 4, 4)

}

//Test PairToWin
func (mysuit *MySuite) TestSicBo_PairToWin(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	description1 := "一点对子 点数 one=1 two=2 three=1"
	test.run().setSender(contractOwner).PairToWin(description1, 1, 2, 1)

	description2 := "二点对子 点数 one=2 two=2 three=1"
	test.run().setSender(contractOwner).PairToWin(description2, 2, 2, 1)

	description3 := "三点对子 点数 one=3 two=3 three=1"
	test.run().setSender(contractOwner).PairToWin(description3, 3, 3, 1)

	description4 := "四点对子 点数 one=4 two=4 three=4"
	test.run().setSender(contractOwner).PairToWin(description4, 4, 4, 4)

	description5 := "五点对子 点数 one=5 two=5 three=3"
	test.run().setSender(contractOwner).PairToWin(description5, 5, 5, 3)

	description6 := "六点对子 点数 one=6 two=1 three=6"
	test.run().setSender(contractOwner).PairToWin(description6, 6, 1, 6)
}

//Test WrapSicIsWin
func (mysuit *MySuite) TestSicBo_WrapSicIsWin(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	description1 := "一点围骰 点数 one=1 two=1 three=1"
	test.run().setSender(contractOwner).WrapSicIsWin(description1, 1, 1, 1)

	description2 := "二点围骰 点数 one=2 two=2 three=2"
	test.run().setSender(contractOwner).WrapSicIsWin(description2, 2, 2, 2)

	description3 := "三点围骰 点数 one=3 two=3 three=3"
	test.run().setSender(contractOwner).WrapSicIsWin(description3, 3, 3, 3)

	description4 := "四点围骰 点数 one=4 two=4 three=4"
	test.run().setSender(contractOwner).WrapSicIsWin(description4, 4, 4, 4)

	description5 := "五点围骰 点数 one=5 two=5 three=5"
	test.run().setSender(contractOwner).WrapSicIsWin(description5, 5, 5, 5)

	description6 := "六点围骰 点数 one=6 two=6 three=6"
	test.run().setSender(contractOwner).WrapSicIsWin(description6, 6, 6, 6)

	description7 := "非围骰 点数 one=5 two=6 three=6"
	test.run().setSender(contractOwner).WrapSicIsWin(description7, 5, 6, 6)
}

//Test SingleIsWin
func (mysuit *MySuite) TestSicBo_SingleIsWin(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	description1 := "单 点数 one=1 two=1 three=1"
	test.run().setSender(contractOwner).SingleIsWin(description1, 2, 1, 1)

	description2 := "双 点数 one=2 two=2 three=2"
	test.run().setSender(contractOwner).SingleIsWin(description2, 3, 2, 2)

}

//Test BigSmallIsWin
func (mysuit *MySuite) TestSicBo_BigSmallIsWin(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	description1 := "小 点数 one=1 two=1 three=1"
	test.run().setSender(contractOwner).BigSmallIsWin(description1, 2, 1, 1)

	description4 := "小 点数 one=6 two=2 three=2"
	test.run().setSender(contractOwner).BigSmallIsWin(description4, 6, 2, 2)

	description2 := "大 点数 one=6 two=6 three=6"
	test.run().setSender(contractOwner).BigSmallIsWin(description2, 5, 6, 6)

	description3 := "大 点数 one=6 two=2 three=3"
	test.run().setSender(contractOwner).BigSmallIsWin(description3, 6, 2, 3)

}

//Test AllSameSic
func (mysuit *MySuite) TestSicBo_AllSameSic(c *check.C) {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)

	description1 := "一点全围 点数 one=1 two=1 three=1"
	test.run().setSender(contractOwner).AllSameSic(description1, 1, 1, 1)

	description2 := "二点全围 点数 one=2 two=2 three=2"
	test.run().setSender(contractOwner).AllSameSic(description2, 2, 2, 2)

	description3 := "三点全围 点数 one=3 two=3 three=3"
	test.run().setSender(contractOwner).AllSameSic(description3, 3, 3, 3)

	description4 := "四点全围 点数 one=4 two=4 three=4"
	test.run().setSender(contractOwner).AllSameSic(description4, 4, 4, 4)

	description5 := "五点全围 点数 one=5 two=5 three=5"
	test.run().setSender(contractOwner).AllSameSic(description5, 5, 5, 5)

	description6 := "六点全围 点数 one=6 two=6 three=6"
	test.run().setSender(contractOwner).AllSameSic(description6, 6, 6, 6)

	description7 := "非全围 点数 one=3 two=1 three=1"
	test.run().setSender(contractOwner).AllSameSic(description7, 3, 1, 1)
}


//TestSicBo_SetBigOrSmall is a method of MySuite
func (mysuit *MySuite) TestSicBo_SetBigOrSmall(c *check.C) () {
	utest.Init(orgID)
	contractOwner := utest.DeployContract(c, contractName, orgID, contractMethods, contractInterfaces)
	test := NewTestObject(contractOwner)
	//test.setSender(contractOwner).InitChain()
	//TODO
	genesisOwner := utest.UTP.Helper().GenesisHelper().Token().Owner()
	utest.Assert(test.run().setSender(utest.UTP.Helper().AccountHelper().AccountOf(genesisOwner)) != nil)
	utest.Transfer(nil, test.obj.sdk.Helper().GenesisHelper().Token().Name(), test.obj.sdk.Message().Contract().Account(), bn.N(1E11))
	accounts := utest.NewAccounts(test.obj.sdk.Helper().GenesisHelper().Token().Name(), bn.N(1E13), 5)
	if accounts == nil {
		panic("初始化newOwner失败")
	}

	//commitLastBlock, pubKey, reveal, commit, signData, _ := PlaceBetHelper(100)
	//utest.AssertError(err, types.CodeOK)

	test.run().setSender(contractOwner).InitChain()
	//utest.AssertError(test.run().setSender(contractOwner).SetSecretSigner(pubKey[:]), types.CodeOK)
	utest.AssertError(test.run().setSender(contractOwner).SetBigOrSmall(10), types.CodeOK)
}