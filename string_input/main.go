package main

import "fmt"

func main(){

	var str string
	fmt.Scan(&str)   // Scan() space k buffer e patai dei
	fmt.Println(str)
	
	var r = 3
	var i int

	var str2 = make([]string, r)  
	for i = 0; i < 3; i++{
		fmt.Scan(&str2[i])
	}
	fmt.Println(str2)

	// fmt.Scan() -> C er scanf("%s") :
	// hello world golang input dile 1st string hello input nibe erpr space teke sb kisu buffer e patiye dibe,
	// 2nd string buffer teke sb niye nibe trpr world k reke bakider k abr buffer e patiye dibe, erpr 3rd string bakigulo nie nibe
	// leading whitespace skip kore

	var str3 []string
	var temp string
	for i = 0; i < r; i++{
		fmt.Scan(temp)
		str3 = append(str3, temp)
	}

	// var str4 [20]string ❌


	fmt.Println(str3)

}

// chatgpt ask koro : str3 kno input niccena
