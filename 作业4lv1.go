package main

import (
	"fmt"
	"strings"
	"time"
	_ "time"
)
func do(){
	var T1 =time.Unix(1009818061,0)
	var T2 =time.Unix(1009818061,0)
	var T3 =time.Unix(1041354061,0)
	var T4 =time.Unix(1072890061,0)
	var T5 =time.Unix(1104512461,0)
	fmt.Println(T1)
	fmt.Println(T2)
	fmt.Println(T3)
	fmt.Println(T4)
	fmt.Println(T5)
	fmt.Println("hello")
}
//noinspection GoUnresolvedReference
func main() {
	var timestemp = make(map[int64]int64, 20)
	timestemp[1] = 1978282061
	timestemp[2] = 1009818061
	timestemp[3] = 1041354061
	timestemp[4] = 1072890061
	timestemp[5] = 1104512461
	for k, v := range timestemp {
		fmt.Println(k, v)
	}
	keys := make([]int64, 0, 100)
	for k := range timestemp {
		//noinspection ALL
		keys = append(keys,k)

	}
	for _, keys := range timestemp {
		fmt.Println(keys)
	}
	str :=
		"---------------------------------------------------\n"
		fmt.Println(str)
	fmt.Println("input:\n")
	fmt.Println("1978282061\n1009818061\n1041354061\n10728900611\n104512461")
	str1 :=
		"---------------------------------------------------\n"
		fmt.Println(str1)



	//fmt.Println(T1.Format("2002-01-02 03:04:05 PM"))
	//fmt.Println(T1.Format("01"))
	//T,_:=time.Parse("01","021")
	//if ("T"=="021"){
	//	fmt.Println(T.Unix())
	//}

	fmt.Println("output:\n")
	if  strings.EqualFold("T1","2002-01-01 01:01:01 AM"){
		fmt.Println("wrong")
	} else {
		fmt.Println("inputok\ninputok\ninputok\ninputok\ninputok\n")
		fmt.Println(" the result are :")
	 	do()

	}

	//var i int64
	//var time stirng
	//for i=0;i<10 ;i++  {
	//	time =time.Unix(timestemp[i],0)
	//
	//}
	//fmt.Println(time)

}