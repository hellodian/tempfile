package tiger

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of Tiger
func (t *Tiger) SetSdk(sdk sdk.ISmartContract) {
	t.sdk = sdk
}

//GetSdk This is a method of Tiger
func (t *Tiger) GetSdk() sdk.ISmartContract {
	return t.sdk
}