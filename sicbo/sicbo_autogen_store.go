package sicbo

import (
	"fmt"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdk/bn"
)

//_setSecretSigner This is a method of SicBo
func (sb *SicBo) _setSecretSigner(v types.PubKey) {
	sb.sdk.Helper().StateHelper().McSet("/secretSigner", &v)
}

//_secretSigner This is a method of SicBo
func (sb *SicBo) _secretSigner() types.PubKey {

	return *sb.sdk.Helper().StateHelper().McGetEx("/secretSigner", new(types.PubKey)).(*types.PubKey)
}

//_clrSecretSigner This is a method of SicBo
func (sb *SicBo) _clrSecretSigner() {
	sb.sdk.Helper().StateHelper().McClear("/secretSigner")
}

//_chkSecretSigner This is a method of SicBo
func (sb *SicBo) _chkSecretSigner() bool {
	return sb.sdk.Helper().StateHelper().Check("/secretSigner")
}

//_McChkSecretSigner This is a method of SicBo
func (sb *SicBo) _McChkSecretSigner() bool {
	return sb.sdk.Helper().StateHelper().McCheck("/secretSigner")
}

//_setLockedInBets This is a method of SicBo
func (sb *SicBo) _setLockedInBets(k string, v bn.Number) {
	sb.sdk.Helper().StateHelper().McSet(fmt.Sprintf("/lockedInBets/%v", k), &v)
}

//_lockedInBets This is a method of SicBo
func (sb *SicBo) _lockedInBets(k string) bn.Number {
	temp := bn.N(0)
	return *sb.sdk.Helper().StateHelper().McGetEx(fmt.Sprintf("/lockedInBets/%v", k), &temp).(*bn.Number)
}

//_clrLockedInBets This is a method of SicBo
func (sb *SicBo) _clrLockedInBets(k string) {
	sb.sdk.Helper().StateHelper().McClear(fmt.Sprintf("/lockedInBets/%v", k))
}

//_chkLockedInBets This is a method of SicBo
func (sb *SicBo) _chkLockedInBets(k string) bool {
	return sb.sdk.Helper().StateHelper().Check(fmt.Sprintf("/lockedInBets/%v", k))
}

//_McChkLockedInBets This is a method of SicBo
func (sb *SicBo) _McChkLockedInBets(k string) bool {
	return sb.sdk.Helper().StateHelper().McCheck(fmt.Sprintf("/lockedInBets/%v", k))
}

//_setSettings This is a method of SicBo
func (sb *SicBo) _setSettings(v *SBSettings) {
	sb.sdk.Helper().StateHelper().McSet("/settings", v)
}

//_settings This is a method of SicBo
func (sb *SicBo) _settings() *SBSettings {

	return sb.sdk.Helper().StateHelper().McGetEx("/settings", new(SBSettings)).(*SBSettings)
}

//_clrSettings This is a method of SicBo
func (sb *SicBo) _clrSettings() {
	sb.sdk.Helper().StateHelper().McClear("/settings")
}

//_chkSettings This is a method of SicBo
func (sb *SicBo) _chkSettings() bool {
	return sb.sdk.Helper().StateHelper().Check("/settings")
}

//_McChkSettings This is a method of SicBo
func (sb *SicBo) _McChkSettings() bool {
	return sb.sdk.Helper().StateHelper().McCheck("/settings")
}

//_setRecFeeInfo This is a method of SicBo
func (sb *SicBo) _setRecFeeInfo(v []RecFeeInfo) {
	sb.sdk.Helper().StateHelper().McSet("/recFeeInfo", &v)
}

//_recFeeInfo This is a method of SicBo
func (sb *SicBo) _recFeeInfo() []RecFeeInfo {

	return *sb.sdk.Helper().StateHelper().McGetEx("/recFeeInfo", new([]RecFeeInfo)).(*[]RecFeeInfo)
}

//_clrRecFeeInfo This is a method of SicBo
func (sb *SicBo) _clrRecFeeInfo() {
	sb.sdk.Helper().StateHelper().McClear("/recFeeInfo")
}

//_chkRecFeeInfo This is a method of SicBo
func (sb *SicBo) _chkRecFeeInfo() bool {
	return sb.sdk.Helper().StateHelper().Check("/recFeeInfo")
}

//_McChkRecFeeInfo This is a method of SicBo
func (sb *SicBo) _McChkRecFeeInfo() bool {
	return sb.sdk.Helper().StateHelper().McCheck("/recFeeInfo")
}

//_setRoundInfo This is a method of SicBo
func (sb *SicBo) _setRoundInfo(k string, v *RoundInfo) {
	sb.sdk.Helper().StateHelper().Set(fmt.Sprintf("/roundInfo/%v", k), v)
}

//_roundInfo This is a method of SicBo
func (sb *SicBo) _roundInfo(k string) *RoundInfo {

	return sb.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/roundInfo/%v", k), new(RoundInfo)).(*RoundInfo)
}

//_chkRoundInfo This is a method of SicBo
func (sb *SicBo) _chkRoundInfo(k string) bool {
	return sb.sdk.Helper().StateHelper().Check(fmt.Sprintf("/roundInfo/%v", k))
}

//_setBetInfo This is a method of SicBo
func (sb *SicBo) _setBetInfo(k1 string, k2 string, v *BetInfo) {
	sb.sdk.Helper().StateHelper().Set(fmt.Sprintf("/betInfo/%v/%v", k1, k2), v)
}

//_betInfo This is a method of SicBo
func (sb *SicBo) _betInfo(k1 string, k2 string) *BetInfo {

	return sb.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/betInfo/%v/%v", k1, k2), new(BetInfo)).(*BetInfo)
}

//_chkBetInfo This is a method of SicBo
func (sb *SicBo) _chkBetInfo(k1 string, k2 string) bool {
	return sb.sdk.Helper().StateHelper().Check(fmt.Sprintf("/betInfo/%v/%v", k1, k2))
}

//_setBigOrSmall This is a method of SicBo
func (sb *SicBo) _setBigOrSmall(v int64) {
	sb.sdk.Helper().StateHelper().Set("/bigOrSmall", &v)
}

//_bigOrSmall This is a method of SicBo
func (sb *SicBo) _bigOrSmall() int64 {

	return *sb.sdk.Helper().StateHelper().GetEx("/bigOrSmall", new(int64)).(*int64)
}

//_chkBigOrSmall This is a method of SicBo
func (sb *SicBo) _chkBigOrSmall() bool {
	return sb.sdk.Helper().StateHelper().Check("/bigOrSmall")
}

