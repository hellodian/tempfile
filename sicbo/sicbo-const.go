package sicbo

const (
	PERMILLE = 1000
)

const (
	BIG        = iota + 1 //large
	SMALL                 //small
	ODD                   //single
	EVEN                  //double
	SAMESIC1              //Three points are the same 1
	SAMESIC2              //Three points are the same 2
	SAMESIC3              //Three points are the same 3
	SAMESIC4              //Three points are the same 4
	SAMESIC5              //Three points are the same 5
	SAMESIC6              //Three points are the same 6
	DOUBLE1               //pair 1
	DOUBLE2               //pair 2
	DOUBLE3               //pair 3
	DOUBLE4               //pair 4
	DOUBLE5               //pair 5
	DOUBLE6               //pair 6
	SUM4                  //Total points 4
	SUM17                 //Total points 17
	SUM5                  //Total points 5
	SUM16                 //Total points 16
	SUM6                  //Total points 6
	SUM15                 //Total points 15
	SUM7                  //Total points 7
	SUM14                 //Total points 14
	SUM8                  //Total points 8
	SUM13                 //Total points 13
	SUM9                  //Total points 9
	SUM10                 //Total points 10
	SUM11                 //Total points 11
	SUM12                 //Total points 12
	PAIGOW12              //Pai Gow 12
	PAIGOW13              //Pai Gow 13
	PAIGOW14              //Pai Gow 14
	PAIGOW15              //Pai Gow 15
	PAIGOW16              //Pai Gow 16
	PAIGOW23              //Pai Gow 23
	PAIGOW24              //Pai Gow 24
	PAIGOW25              //Pai Gow 25
	PAIGOW26              //Pai Gow 26
	PAIGOW34              //Pai Gow 34
	PAIGOW35              //Pai Gow 35
	PAIGOW36              //Pai Gow 36
	PAIGOW45              //Pai Gow 45
	PAIGOW46              //Pai Gow 46
	PAIGOW56              //Pai Gow 56
	POINT1                //point 1
	POINT2                //point 2
	POINT3                //point 3
	POINT4                //point 4
	POINT5                //point 5
	POINT6                //point 6
	ALLSAMESIC            //All three of them have the same number of points
)

//Winning multiple  first version
//const (
//	ONEMULTIPLE             = 1
//	TWOMULTIPLE             = 2
//	THREEMULTIPLE           = 3
//	FIVEMULTIPLE            = 5
//	SIXMULTIPLE             = 6
//	EIGHTMULTIPLES          = 8
//	TWELVEMULTIPLE          = 12
//	FOURTEENMULTIPLE        = 14
//	EIGHTEENMULTIPLE        = 18
//	TWENTYFOURMULTIPLE      = 24
//	FIFTYMULTIPLE           = 50
//	ONEHUNDREDFIFTYMULTIPLE = 150
//)

//Winning multiple  third version
const (
	BIGORSMALL              = 1
	ONEMULTIPLE             = 1
	TWOMULTIPLE             = 2
	THREEMULTIPLE           = 3
	FIVEMULTIPLE            = 5
	SIXMULTIPLE             = 6
	EIGHTMULTIPLES          = 8
	TWELVEMULTIPLE          = 12
	FOURTEENMULTIPLE        = 14
	EIGHTEENMULTIPLE        = 18
	TWENTYFOURMULTIPLE      = 25
	FIFTYMULTIPLE           = 50
	ONEHUNDREDFIFTYMULTIPLE = 150
)

//Winning multiple  second version
//const (
//	ONEMULTIPLE             = 1
//	TWOMULTIPLE             = 2
//	THREEMULTIPLE           = 3
//	FIVEMULTIPLE            = 5
//	SIXMULTIPLE             = 6
//	EIGHTMULTIPLES          = 10
//	TWELVEMULTIPLE          = 12
//	FOURTEENMULTIPLE        = 18
//	EIGHTEENMULTIPLE        = 30
//	TWENTYFOURMULTIPLE      = 30
//	FIFTYMULTIPLE           = 60
//	ONEHUNDREDFIFTYMULTIPLE = 180
//)




const (
	NOAWARD       = iota //Not the lottery
	AWARDED              //Has the lottery
	REFUNDED             //refunded
	OPENINGAPRIZE        //Is the lottery
)
