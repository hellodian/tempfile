package tiger

import "blockchain/smcsdk/sdk/types"

//GetNineLine 数字是多少  几个数相连
func (t *Tiger) ClacByLine(p []Poker)(mul ,sum int64,s []Poker){
	s=make([]Poker,0)
	if p[0].Value==0{
		return
	}else {
		mul=p[0].Value
		s=append(s,p[0])
	}
	ww:=p[1:]
	for _,pv:=range ww {

		switch pv.Value {
		case 0:
			{
				if mul!=1{//w 不能代替s
					s = append(s, pv)
				}else {
					sum = int64(len(s))
					return
				}

			}
		case 1:
			{
				if mul == 1 {
					s = append(s, pv)
				} else {
					sum = int64(len(s))

					return
				}
			}
		case 2:
			{
				if mul == 2 {
					s = append(s, pv)
				} else {
					sum = int64(len(s))
					return
				}
			}
		case 3:
			{
				if mul == 3 {
					s = append(s, pv)
				} else {
					sum = int64(len(s))
					return
				}
			}

		case 4:
			{
				if mul == 4 {
					s = append(s, pv)
				} else {
					sum = int64(len(s))
					return
				}
			}

		case 9:
			{
				if mul == 9 {
					s = append(s, pv)
				} else {
					sum = int64(len(s))
					return
				}
			}

		case 10:
			{
				if mul == 10 {
					s = append(s, pv)
				} else {
					sum = int64(len(s))
					return
				}

			}
		}

	}

	sum = int64(len(s))
	return

}

//ClacMul  This is a sample method   计算赔率
func (t *Tiger) ClacMul(mul,sum int64) (multiple int64) {

	if sum<3||mul==0{
		return
	}
	//W=0
	//K=4
	//Q=3
	//J=2
	//S=1
	switch mul {
	case 1:
		if sum==3{
			multiple =2
		}else if sum==4{
			multiple =5
		}else if sum==5{
			multiple =100
		}
	case 2:
		if sum==3{
			multiple =6
		}else if sum==4{
			multiple =20
		}else if sum==5{
			multiple =30
		}
	case 3:
		if sum==3{
			multiple =8
		}else if sum==4{
			multiple =25
		}else if sum==5{
			multiple =30
		}
	case 4:
		if sum==3{
			multiple =10
		}else if sum==4{
			multiple =30
		}else if sum==5{
			multiple =50
		}
	case 9:
		if sum==3{
			multiple =3
		}else if sum==4{
			multiple =5
		}else if sum==5{
			multiple =10
		}
	case 10:
		if sum==3{
			multiple =5
		}else if sum==4{
			multiple =8
		}else if sum==5{
			multiple =15
		}


	}
	return
}

//ClacMul  This is a sample method   计算免费游戏次数
func (t *Tiger) ClacFee(p []Poker) (sum int64){

	for _,v:=range p {
		if v.Value==1{
			sum++
		}

	}
	if sum<3{
		sum=0
		return
	}
	if sum==3 {
		sum=3
		return
	}
	if sum==4 {
		sum=5
		return
	}
	if sum==5 {
		sum=15
		return
	}
	return
}



//GetNineLine 获取9连线
func (t *Tiger) GetNineLine(address types.Address) (totalMul, totalFee int64, pokeList [][]Poker) {
	//t.poker=[3][5]Poker{
	//	[5]Poker{{0, 0 ,2},{0 ,1 ,3}, {0, 2, 3}, {0, 3 ,4} ,{0, 4 ,3}},
	//	[5]Poker{{1 ,0 ,9} ,{1, 1 ,10} ,{1 ,2, 9} ,{1 ,3 ,10} ,{1, 4 ,9}},
	//	[5]Poker{{2, 0, 9} ,{2, 1, 9}, {2, 2, 9}, {2 ,3 ,9} ,{2, 4 ,9}},
	//}
	poker := t._betInfo(address).Poker
	firstLine := []Poker{poker[0][0], poker[0][1], poker[0][2], poker[0][3], poker[0][4]}
	secondeLine := []Poker{poker[1][0], poker[1][1], poker[1][2], poker[1][3], poker[1][4]}
	thirdLine := []Poker{poker[2][0], poker[2][1], poker[2][2], poker[2][3], poker[2][4]}
	fourLine := []Poker{poker[0][0], poker[1][1], poker[2][2], poker[1][3], poker[0][4]}
	fiveLine := []Poker{poker[2][0], poker[1][1], poker[0][2], poker[1][3], poker[2][4]}
	sixLine := []Poker{poker[0][0], poker[0][1], poker[1][2], poker[0][3], poker[0][4]}
	sevenLine := []Poker{poker[2][0], poker[2][1], poker[1][2], poker[2][3], poker[2][4]}
	eightLine := []Poker{poker[1][0], poker[2][1], poker[2][2], poker[2][3], poker[1][4]}
	nineLine := []Poker{poker[1][0], poker[0][1], poker[0][2], poker[0][3], poker[1][4]}
	pokeList = make([][]Poker, 0) //为了返回中奖的连线
	totalMul = 0                  //倍率
	totalFee = 0                  //免费游戏次数
	mul, sum, p := t.ClacByLine(firstLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	//得到赔率
	mulNum := t.ClacMul(mul, sum)
	//获得免费游戏次数
	fee := t.ClacFee(firstLine)
	totalMul += mulNum
	totalFee += fee

	mul, sum, p = t.ClacByLine(secondeLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	mulNum = t.ClacMul(mul, sum)
	fee = t.ClacFee(secondeLine)
	totalMul += mulNum
	totalFee += fee
	mul, sum, p = t.ClacByLine(thirdLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	mulNum = t.ClacMul(mul, sum)
	fee = t.ClacFee(thirdLine)
	totalMul += mulNum
	totalFee += fee
	mul, sum, p = t.ClacByLine(fourLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	mulNum = t.ClacMul(mul, sum)
	fee = t.ClacFee(fourLine)
	totalMul += mulNum
	totalFee += fee
	mul, sum, p = t.ClacByLine(fiveLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	mulNum = t.ClacMul(mul, sum)
	fee = t.ClacFee(fiveLine)
	totalMul += mulNum
	totalFee += fee
	mul, sum, p = t.ClacByLine(sixLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	mulNum = t.ClacMul(mul, sum)
	fee = t.ClacFee(sixLine)
	totalMul += mulNum
	totalFee += fee
	mul, sum, p = t.ClacByLine(sevenLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	mulNum = t.ClacMul(mul, sum)
	fee = t.ClacFee(sevenLine)
	totalMul += mulNum
	totalFee += fee
	mul, sum, p = t.ClacByLine(eightLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	mulNum = t.ClacMul(mul, sum)
	fee = t.ClacFee(eightLine)
	totalMul += mulNum
	totalFee += fee
	mul, sum, p = t.ClacByLine(nineLine)
	if sum >= 3 {
		pokeList = append(pokeList, p)
	}
	mulNum = t.ClacMul(mul, sum)
	fee = t.ClacFee(nineLine)
	totalMul += mulNum
	totalFee += fee

	return
}


