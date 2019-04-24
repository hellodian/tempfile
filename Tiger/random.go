package tiger
//
//import (
//	"blockchain/algorithm"
//	_ "blockchain/smcsdk/sdk"
//	"blockchain/smcsdk/sdk/bn"
//	_ "blockchain/smcsdk/sdk/jsoniter"
//	"blockchain/smcsdk/sdk/types"
//	"blockchain/smcsdk/utest"
//	"common/keys"
//	"common/kms"
//	"encoding/hex"
//	"fmt"
//	"github.com/tendermint/go-amino"
//	"github.com/tendermint/go-crypto"
//	"github.com/tendermint/tmlibs/common"
//	"io/ioutil"
//	_ "math"
//	_"testing"
//)
//
//const (
//	ownerName1 = "local_owner"
//	password1  = "12345678"
//)
//
//var (
//	cdc = amino.NewCodec()
//)
//
//func init() {
//	crypto.RegisterAmino(cdc)
//	crypto.SetChainId("local")
//	kms.InitKMS("./.keystore", "local_mode", "", "", "0x1003")
//	kms.GenPrivKey(ownerName1, []byte(password1))
//}
//
////hempHeight 想对于下注高度和生效高度之间的差值
////acct 合约的owner
//func PlaceBetHelper1(tempHeight int64) (commitLastBlock int64, pubKey [32]byte, reveal, commit []byte, signData [64]byte, err types.Error) {
//	acct, err := Load("./.keystore/local_owner.wal", []byte(password1), nil)
//	if err.ErrorCode != types.CodeOK {
//		return
//	}
//
//	localBlockHeight := utest.UTP.ISmartContract.Block().Height()
//
//	pubKey = acct.PubKey.(crypto.PubKeyEd25519)
//
//	commitLastBlock = localBlockHeight + tempHeight
//	decode := crypto.CRandBytes(32)
//	revealStr := hex.EncodeToString(algorithm.SHA3256(decode))
//	reveal, _ = hex.DecodeString(revealStr)
//
//	commit = algorithm.SHA3256(reveal)
//
//	signByte := append(bn.N(commitLastBlock).Bytes(), commit...)
//	signData = acct.PrivKey.Sign(signByte).(crypto.SignatureEd25519)
//
//	return
//}
//
//func Load(keystorePath string, password, fingerprint []byte) (acct *keys.Account, err types.Error) {
//	if keystorePath == "" {
//		common.PanicSanity("Cannot loads account because keystorePath not set")
//	}
//
//	walBytes, mErr := ioutil.ReadFile(keystorePath)
//	if mErr != nil {
//		err.ErrorCode = types.ErrInvalidParameter
//		err.ErrorDesc = "account does not exist"
//		return
//	}
//
//	jsonBytes, mErr := algorithm.DecryptWithPassword(walBytes, password, fingerprint)
//	if mErr != nil {
//		err.ErrorCode = types.ErrInvalidParameter
//		err.ErrorDesc = fmt.Sprintf("the password is wrong err info : %s", mErr)
//		return
//	}
//
//	acct = new(keys.Account)
//	mErr = cdc.UnmarshalJSON(jsonBytes, acct)
//	if mErr != nil {
//		err.ErrorCode = types.ErrInvalidParameter
//		err.ErrorDesc = fmt.Sprintf("UnmarshalJSON is wrong err info : %s", mErr)
//		return
//	}
//
//	acct.KeystorePath = keystorePath
//	err.ErrorCode = types.CodeOK
//	return
//}
//
//func (t *Tiger) GetReval()(reval []byte){
//
//	PlaceBetHelper1(100)
//}
//
