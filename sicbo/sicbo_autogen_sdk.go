package sicbo

import (
	"blockchain/smcsdk/sdk"
)

//SetSdk This is a method of SicBo
func (sb *SicBo) SetSdk(sdk sdk.ISmartContract) {
	sb.sdk = sdk
}

//GetSdk This is a method of SicBo
func (sb *SicBo) GetSdk() sdk.ISmartContract {
	return sb.sdk
}