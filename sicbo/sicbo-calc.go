package sicbo

import (
	"blockchain/smcsdk/sdk"
	"blockchain/smcsdk/sdk/bn"
	"blockchain/smcsdk/sdk/crypto/sha3"
	"blockchain/smcsdk/sdk/types"
)

type Sic struct {
	One   int64 `json:"One"`
	Two   int64 `json:"Two"`
	Three int64 `json:"Three"`
}

//Total points
func (sb *SicBo) Total(sic *Sic) int64 {
	return sic.One + sic.Two + sic.Three
}

//odd or even number
func (sb *SicBo) IsSingle(sic *Sic) bool {
	total := sb.Total(sic)
	if total%2 == 0 {
		return false
	} else {
		return true
	}
}

//Is it dice or not
func (sb *SicBo) isWrapSic(sic *Sic) bool {
	if sic.One == sic.Two && sic.Two == sic.Three {
		return true
	} else {
		return false
	}
}

//Wai roll a few
func (sb *SicBo) WrapSicTo(sic *Sic) int64 {
	to := int64(0)
	if sb.isWrapSic(sic) {
		to = sic.One
	}
	return to
}

//Presence of pairs
func (sb *SicBo) hasPair(sic *Sic) (bool, int64) {
	if sic.One^sic.Two == 0 {
		return true, sic.One
	}
	if sic.One^sic.Three == 0 {
		return true, sic.One
	}
	if sic.Two^sic.Three == 0 {
		return true, sic.Two
	}
	return false, int64(0)
}

//For a few
func (sb *SicBo)PairTo(sic *Sic) int64 {
	to := int64(0)
	b, i := sb.hasPair(sic)
	if b {
		to = i
	}
	return to
}

//Repeat the count several times
func (sb *SicBo) RepeatNum(sic *Sic,a int64) int64 {
	var n int64
	n = sb.hand(sic,a)
	return n
}

func (sb *SicBo) hand(sic *Sic,a int64) int64 {
	n := int64(0)
	switch a {
	case POINT1:
		if sic.One == 1 {
			n++
		}
		if sic.Two == 1 {
			n++
		}
		if sic.Three == 1 {
			n++
		}
	case POINT2:
		if sic.One == 2 {
			n++
		}
		if sic.Two == 2 {
			n++
		}
		if sic.Three == 2 {
			n++
		}
	case POINT3:
		if sic.One == 3 {
			n++
		}
		if sic.Two == 3 {
			n++
		}
		if sic.Three == 3 {
			n++
		}
	case POINT4:
		if sic.One == 4 {
			n++
		}
		if sic.Two == 4 {
			n++
		}
		if sic.Three == 4 {
			n++
		}
	case POINT5:
		if sic.One == 5 {
			n++
		}
		if sic.Two == 5 {
			n++
		}
		if sic.Three == 5 {
			n++
		}
	case POINT6:
		if sic.One == 6 {
			n++
		}
		if sic.Two == 6 {
			n++
		}
		if sic.Three == 6 {
			n++
		}
	}
	return n
}

func (sb *SicBo) SplitPaiJuClazz(sic *Sic,sum, b, c int64, pj []int64) []int64 {
	switch sum {
	case 3:
		if !sb.hasNonEle(sic,pj, PAIGOW12) {
			pj = append(pj, PAIGOW12)
		}
	case 4:
		if !sb.hasNonEle(sic,pj, PAIGOW13) {
			pj = append(pj, PAIGOW13)
		}
	case 5:
		if b == 1 || c == 1 {
			if !sb.hasNonEle(sic,pj, PAIGOW14) {
				pj = append(pj, PAIGOW14)
			}
		}
		if b == 2 || c == 2 {
			if !sb.hasNonEle(sic,pj, PAIGOW23) {
				pj = append(pj, PAIGOW23)
			}
		}
	case 6:
		if b == 1 || c == 1 {
			if !sb.hasNonEle(sic,pj, PAIGOW15) {
				pj = append(pj, PAIGOW15)
			}
		}
		if b == 2 || c == 2 {
			if !sb.hasNonEle(sic,pj, PAIGOW24) {
				pj = append(pj, PAIGOW24)
			}
		}
	case 7:
		if b == 1 || c == 1 {
			if !sb.hasNonEle(sic,pj, PAIGOW16) {
				pj = append(pj, PAIGOW16)
			}
		}
		if b == 2 || c == 2 {
			if !sb.hasNonEle(sic,pj, PAIGOW25) {
				pj = append(pj, PAIGOW25)
			}
		}
		if b == 3 || c == 3 {
			if !sb.hasNonEle(sic,pj, PAIGOW34) {
				pj = append(pj, PAIGOW34)
			}
		}
	case 8:
		if b == 2 || c == 2 {
			if !sb.hasNonEle(sic,pj, PAIGOW26) {
				pj = append(pj, PAIGOW26)
			}
		}
		if b == 3 || c == 3 {
			if !sb.hasNonEle(sic,pj, PAIGOW35) {
				pj = append(pj, PAIGOW35)
			}
		}
	case 9:
		if b == 3 || c == 3 {
			if !sb.hasNonEle(sic,pj, PAIGOW36) {
				pj = append(pj, PAIGOW36)
			}
		}
		if b == 4 || c == 4 {
			if !sb.hasNonEle(sic,pj, PAIGOW45) {
				pj = append(pj, PAIGOW45)
			}
		}
	case 10:
		if !sb.hasNonEle(sic,pj, PAIGOW46) {
			pj = append(pj, PAIGOW46)
		}
	case 11:
		if !sb.hasNonEle(sic,pj, PAIGOW56) {
			pj = append(pj, PAIGOW56)
		}
	}
	return pj
}

//Does the element exist
func (sb *SicBo) hasNonEle(sic *Sic,pj []int64, a int64) bool {
	flag := false
	for _, v := range pj {
		flag = v == a
		if flag {
			return flag
		}
	}
	return flag
}

func (sb *SicBo) PaiJuTo(sic *Sic) []int64 {
	var pj []int64
	sum1 := sic.One + sic.Two
	if sic.One != sic.Two && sum1 >= 3 && sum1 <= 11 {
		pj = sb.SplitPaiJuClazz(sic,sum1, sic.One, sic.Two, pj)
	}
	sum2 := sic.One + sic.Three
	if sic.One != sic.Three && sum2 >= 3 && sum2 <= 11 {
		pj = sb.SplitPaiJuClazz(sic,sum2, sic.One, sic.Three, pj)
	}
	sum3 := sic.Two + sic.Three
	if sic.Two != sic.Three && sum3 >= 3 && sum3 <= 11 {
		pj = sb.SplitPaiJuClazz(sic,sum3, sic.Two, sic.Three, pj)
	}
	return pj
}

//Calculate the actual winning amount
func (sb *SicBo) GetOneWin(sic *Sic,data BetData) (amount bn.Number) {
	amount = bn.N(0)
	isWin := false
	mode := int64(data.BetMode)
	switch {

	case BIG <= mode && mode <= SMALL: //Large and small 1 - 2
		isWin = sb.BigSmallIsWin(sic,mode)

	case ODD <= mode && mode <= EVEN: //Single and double 3 - 4
		isWin = sb.SingleIsWin(sic,mode)

	case SAMESIC1 <= mode && mode <= SAMESIC6: //Three of the same number of points 5 - 10
		isWin = sb.WrapSicIsWin(sic,mode)

	case DOUBLE1 <= mode && mode <= DOUBLE6: //pair 11 -16
		isWin = sb.PairToWin(sic,mode)

	case SUM4 <= mode && mode <= SUM12: //The bet is the total number of points 17 - 30
		isWin = sb.SumWin(sic,mode)

	case PAIGOW12 <= mode && mode <= PAIGOW56: //pai gow 31 - 45
		isWin = sb.PaiJuToIsWin(sic,mode)

	case POINT1 <= mode && mode <= POINT6: //point 46 - 51
		var n int64
		isWin, n = sb.PointIsWin(sic,mode)
		if isWin {
			amount = data.BetAmount.MulI(n) // +1 ?
			return amount
		}
	case ALLSAMESIC == mode:
		isWin = sb.isWrapSic(sic)
	}
	if !isWin {
		return
	}
	amount = sb.MaybeWinAmountByOne(data)

	return
}

//Calculate multiple entry bonus
func (sb *SicBo) GetBetWinAmount(sic *Sic,betList []BetData) (totalWinAmount bn.Number) {
	totalWinAmount = bn.N(0)
	if len(betList) <= 0 {
		return
	}
	for _, bet := range betList {
		//Calculate note bonus
		amount := sb.GetOneWin(sic,bet)
		//Bonus greater than 0
		if bn.N(0).Cmp(amount) < 0 {
			totalWinAmount = totalWinAmount.Add(amount)
		}
	}
	return
}

//Determine whether size wins or not
func (sb *SicBo) BigSmallIsWin(sic *Sic,mode int64) bool {

	//庄家围骰通吃
	if sb.isWrapSic(sic) {
		return false
	}

	total := sb.Total(sic)

	//little
	if 4 <= total && total <= 10 && mode == SMALL {
		return true
	}
	if 11 <= total && total <= 17 && mode == BIG {
		return true
	}
	return false
}

//Judge whether the singles and doubles win or not
func (sb *SicBo) SingleIsWin(sic *Sic,mode int64) bool {
	//庄家围骰通吃
	if sb.isWrapSic(sic) {
		return false
	}

	isSingle := sb.IsSingle(sic)

	//single
	if mode == ODD && isSingle {
		return true
	}
	//double
	if mode == EVEN && isSingle == false {
		return true
	}
	return false
}

//Determine whether the dice win or not
func (sb *SicBo) WrapSicIsWin(sic *Sic,mode int64) bool {
	point := sb.WrapSicTo(sic)
	isWin := false
	switch mode {
	case SAMESIC1:
		if point == 1 {
			isWin = true
		}
	case SAMESIC2:
		if point == 2 {
			isWin = true
		}
	case SAMESIC3:
		if point == 3 {
			isWin = true
		}
	case SAMESIC4:
		if point == 4 {
			isWin = true
		}
	case SAMESIC5:
		if point == 5 {
			isWin = true
		}
	case SAMESIC6:
		if point == 6 {
			isWin = true
		}
	case ALLSAMESIC:
		switch point {
		case 1, 2, 3, 4, 5, 6:
			isWin = true
		}
	}
	return isWin
}

//Judge whether the pair wins or not
func (sb *SicBo) PairToWin(sic *Sic,mode int64) bool {
	point := sb.PairTo(sic)
	isWin := false
	switch mode {
	case DOUBLE1:
		if point == 1 {
			isWin = true
		}
	case DOUBLE2:
		if point == 2 {
			isWin = true
		}
	case DOUBLE3:
		if point == 3 {
			isWin = true
		}
	case DOUBLE4:
		if point == 4 {
			isWin = true
		}
	case DOUBLE5:
		if point == 5 {
			isWin = true
		}
	case DOUBLE6:
		if point == 6 {
			isWin = true
		}
	}
	return isWin
}

//Determine points and whether to win or not
func (sb *SicBo) SumWin(sic *Sic,mode int64) bool {
	sum := sb.Total(sic)
	isWin := false
	switch mode {
	case SUM4:
		if sum == 4 {
			isWin = true
		}
	case SUM5:
		if sum == 5 {
			isWin = true
		}
	case SUM6:
		if sum == 6 {
			isWin = true
		}
	case SUM7:
		if sum == 7 {
			isWin = true
		}
	case SUM8:
		if sum == 8 {
			isWin = true
		}
	case SUM9:
		if sum == 9 {
			isWin = true
		}
	case SUM10:
		if sum == 10 {
			isWin = true
		}
	case SUM11:
		if sum == 11 {
			isWin = true
		}
	case SUM12:
		if sum == 12 {
			isWin = true
		}
	case SUM13:
		if sum == 13 {
			isWin = true
		}
	case SUM14:
		if sum == 14 {
			isWin = true
		}
	case SUM15:
		if sum == 15 {
			isWin = true
		}
	case SUM16:
		if sum == 16 {
			isWin = true
		}
	case SUM17:
		if sum == 17 {
			isWin = true
		}
	}
	return isWin
}

func (sb *SicBo) PaiJuToIsWin(sic *Sic,mode int64) bool {
	winModeList := sb.PaiJuTo(sic)
	isWin := false
	for _, value := range winModeList {
		if value == mode {
			isWin = true
		}
	}
	return isWin
}

func (sb *SicBo) PointIsWin(sic *Sic,mode int64) (bool, int64) {
	num := sb.RepeatNum(sic,mode)
	return num > 0, int64(num)
}

//Result list
type SicResSlice []*Sic

//Get the lottery results
func (sb *SicBo) GetSicRes(reveal []byte) *Sic {
	//rand.Seed(time.Now().UnixNano())
	bytes := bn.NBytes(sha3.Sum256(reveal, sb.sdk.Block().BlockHash(), sb.sdk.Block().RandomNumber()))
	sicResSlices := sb.Shuffle(reveal, TotalRes())

	le := len(sicResSlices)
	i := bytes.ModI(int64(le))
	//fmt.Println("--------->", i)
	return sicResSlices[i.V.Int64()]
}

//shuffle
func (sb *SicBo) Shuffle(reveal []byte, sic SicResSlice) SicResSlice {
	bytes := bn.NBytes(sha3.Sum256(reveal, sb.sdk.Block().BlockHash(), sb.sdk.Block().RandomNumber()))
	size := len(sic)
	for i := size; i > 1; i-- {
		ran := bytes.ModI(int64(i))
		sic.swap(int64(i-1), ran.V.Int64())
	}
	return sic
}

//swap
func (sic SicResSlice) swap(i int64, j int64) {
	sic[i], sic[j] = sic[j], sic[i]
}

//Get the result list
func TotalRes() SicResSlice {
	var bets []*Sic
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			for k := 1; k <= 6; k++ {
				bets = append(bets, &Sic{int64(i), int64(j), int64(k)})
			}
		}
	}
	return bets
}

//Calculate the amount of money a long bet is likely to win: total bet amount betList scheme
func (sb *SicBo) MaybeWinAmountAndFeeByList(amount bn.Number, betList []BetData) (totalMaybeWinAmount, fee bn.Number) {
	fee = bn.N(0)
	totalMaybeWinAmount = bn.N(0)
	settings := sb._settings()
	// fee
	fee = amount.MulI(settings.FeeRatio).DivI(PERMILLE)
	if fee.CmpI(settings.FeeMiniNum) < 0 {
		fee = bn.N(settings.FeeMiniNum)
	}

	sdk.Require(fee.Cmp(amount) <= 0,
		types.ErrInvalidParameter, "Bet doesn't even cover fee")

	for _, bet := range betList {
		totalMaybeWinAmount = totalMaybeWinAmount.Add(sb.MaybeWinAmountByOne(bet))
	}
	return
}

//Calculate the amount of money a bet is likely to win
func (sb *SicBo) MaybeWinAmountByOne(bet BetData) (amount bn.Number) {

	var multiple int64

	switch bet.BetMode {
	case BIG, SMALL:
		multiple = sb._bigOrSmall()
	case ODD, EVEN:
		multiple = ONEMULTIPLE
	case SAMESIC1, SAMESIC2, SAMESIC3, SAMESIC4, SAMESIC5, SAMESIC6:
		multiple = ONEHUNDREDFIFTYMULTIPLE
	case ALLSAMESIC:
		multiple = TWENTYFOURMULTIPLE
	case DOUBLE1, DOUBLE2, DOUBLE3, DOUBLE4, DOUBLE5, DOUBLE6:
		multiple = EIGHTMULTIPLES
	case SUM4, SUM17:
		multiple = FIFTYMULTIPLE
	case SUM5, SUM16:
		multiple = EIGHTEENMULTIPLE
	case SUM6, SUM15:
		multiple = FOURTEENMULTIPLE
	case SUM7, SUM14:
		multiple = TWELVEMULTIPLE
	case SUM8, SUM13:
		multiple = EIGHTMULTIPLES
	case SUM9, SUM10, SUM11, SUM12:
		multiple = SIXMULTIPLE
	case PAIGOW12, PAIGOW13, PAIGOW14, PAIGOW15, PAIGOW16, PAIGOW23, PAIGOW24, PAIGOW25, PAIGOW26, PAIGOW34, PAIGOW35, PAIGOW36, PAIGOW45, PAIGOW46, PAIGOW56:
		multiple = FIVEMULTIPLE
	case POINT1, POINT2, POINT3, POINT4, POINT5, POINT6:
		multiple = THREEMULTIPLE
	default:
		multiple = 0
	}

	if multiple > 0 {
		multiple++
	}
	amount = bet.BetAmount.MulI(multiple)
	return
}
