package tiger

import (
	"fmt"
	"blockchain/smcsdk/sdk/types"
	"blockchain/smcsdk/sdk/bn"
)

//_setSecretSigner This is a method of Tiger
func (t *Tiger) _setSecretSigner(v types.PubKey) {
	t.sdk.Helper().StateHelper().McSet("/secretSigner", &v)
}

//_secretSigner This is a method of Tiger
func (t *Tiger) _secretSigner() types.PubKey {

	return *t.sdk.Helper().StateHelper().McGetEx("/secretSigner", new(types.PubKey)).(*types.PubKey)
}

//_clrSecretSigner This is a method of Tiger
func (t *Tiger) _clrSecretSigner() {
	t.sdk.Helper().StateHelper().McClear("/secretSigner")
}

//_chkSecretSigner This is a method of Tiger
func (t *Tiger) _chkSecretSigner() bool {
	return t.sdk.Helper().StateHelper().Check("/secretSigner")
}

//_McChkSecretSigner This is a method of Tiger
func (t *Tiger) _McChkSecretSigner() bool {
	return t.sdk.Helper().StateHelper().McCheck("/secretSigner")
}

//_setLockedInBets This is a method of Tiger
func (t *Tiger) _setLockedInBets(k string, v bn.Number) {
	t.sdk.Helper().StateHelper().McSet(fmt.Sprintf("/lockedInBets/%v", k), &v)
}

//_lockedInBets This is a method of Tiger
func (t *Tiger) _lockedInBets(k string) bn.Number {
	temp := bn.N(0)
	return *t.sdk.Helper().StateHelper().McGetEx(fmt.Sprintf("/lockedInBets/%v", k), &temp).(*bn.Number)
}

//_clrLockedInBets This is a method of Tiger
func (t *Tiger) _clrLockedInBets(k string) {
	t.sdk.Helper().StateHelper().McClear(fmt.Sprintf("/lockedInBets/%v", k))
}

//_chkLockedInBets This is a method of Tiger
func (t *Tiger) _chkLockedInBets(k string) bool {
	return t.sdk.Helper().StateHelper().Check(fmt.Sprintf("/lockedInBets/%v", k))
}

//_McChkLockedInBets This is a method of Tiger
func (t *Tiger) _McChkLockedInBets(k string) bool {
	return t.sdk.Helper().StateHelper().McCheck(fmt.Sprintf("/lockedInBets/%v", k))
}

//_setSettings This is a method of Tiger
func (t *Tiger) _setSettings(v *Settings) {
	t.sdk.Helper().StateHelper().McSet("/settings", v)
}

//_settings This is a method of Tiger
func (t *Tiger) _settings() *Settings {

	return t.sdk.Helper().StateHelper().McGetEx("/settings", new(Settings)).(*Settings)
}

//_clrSettings This is a method of Tiger
func (t *Tiger) _clrSettings() {
	t.sdk.Helper().StateHelper().McClear("/settings")
}

//_chkSettings This is a method of Tiger
func (t *Tiger) _chkSettings() bool {
	return t.sdk.Helper().StateHelper().Check("/settings")
}

//_McChkSettings This is a method of Tiger
func (t *Tiger) _McChkSettings() bool {
	return t.sdk.Helper().StateHelper().McCheck("/settings")
}

//_setRecFeeInfo This is a method of Tiger
func (t *Tiger) _setRecFeeInfo(v []RecFeeInfo) {
	t.sdk.Helper().StateHelper().McSet("/recFeeInfo", &v)
}

//_recFeeInfo This is a method of Tiger
func (t *Tiger) _recFeeInfo() []RecFeeInfo {

	return *t.sdk.Helper().StateHelper().McGetEx("/recFeeInfo", new([]RecFeeInfo)).(*[]RecFeeInfo)
}

//_clrRecFeeInfo This is a method of Tiger
func (t *Tiger) _clrRecFeeInfo() {
	t.sdk.Helper().StateHelper().McClear("/recFeeInfo")
}

//_chkRecFeeInfo This is a method of Tiger
func (t *Tiger) _chkRecFeeInfo() bool {
	return t.sdk.Helper().StateHelper().Check("/recFeeInfo")
}

//_McChkRecFeeInfo This is a method of Tiger
func (t *Tiger) _McChkRecFeeInfo() bool {
	return t.sdk.Helper().StateHelper().McCheck("/recFeeInfo")
}

//_setPokerMainSet This is a method of Tiger
func (t *Tiger) _setPokerMainSet(v [5][20]int64) {
	t.sdk.Helper().StateHelper().Set("/pokerMainSet", &v)
}

//_pokerMainSet This is a method of Tiger
func (t *Tiger) _pokerMainSet() [5][20]int64 {

	return *t.sdk.Helper().StateHelper().GetEx("/pokerMainSet", new([5][20]int64)).(*[5][20]int64)
}

//_chkPokerMainSet This is a method of Tiger
func (t *Tiger) _chkPokerMainSet() bool {
	return t.sdk.Helper().StateHelper().Check("/pokerMainSet")
}

//_setPokerFeeSet This is a method of Tiger
func (t *Tiger) _setPokerFeeSet(v [5][20]int64) {
	t.sdk.Helper().StateHelper().Set("/pokerFeeSet", &v)
}

//_pokerFeeSet This is a method of Tiger
func (t *Tiger) _pokerFeeSet() [5][20]int64 {

	return *t.sdk.Helper().StateHelper().GetEx("/pokerFeeSet", new([5][20]int64)).(*[5][20]int64)
}

//_chkPokerFeeSet This is a method of Tiger
func (t *Tiger) _chkPokerFeeSet() bool {
	return t.sdk.Helper().StateHelper().Check("/pokerFeeSet")
}

//_setBetInfo This is a method of Tiger
func (t *Tiger) _setBetInfo(k types.Address, v *PlayerInfo) {
	t.sdk.Helper().StateHelper().Set(fmt.Sprintf("/betInfo/%v", k), v)
}

//_betInfo This is a method of Tiger
func (t *Tiger) _betInfo(k types.Address) *PlayerInfo {

	return t.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/betInfo/%v", k), new(PlayerInfo)).(*PlayerInfo)
}

//_chkBetInfo This is a method of Tiger
func (t *Tiger) _chkBetInfo(k types.Address) bool {
	return t.sdk.Helper().StateHelper().Check(fmt.Sprintf("/betInfo/%v", k))
}

