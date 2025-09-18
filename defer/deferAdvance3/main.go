package main

import "fmt"

func calculate() (result int) {
	fmt.Println("first", result)  // first 0 print hobe

	show := func() {   // heap e show() & result joma hoise
		result = result + 10
		fmt.Println("show", result)
	}

	defer show()  // show() te result er valur referrence result, func call hobe pore

	result = 5

	p := func(a int) {  // heap e p() & result[a] joma hobe
		fmt.Println("p", a)
	}

	defer p(result)   // arg pass kora hoice tai p() te result er value 5, func call hobe pore

	defer fmt.Println("print1", result)  //  ekane result er value 5 , print hobe pore

	result = 10

	fmt.Println("second", result)  // second 10 print hobe

	defer fmt.Println("print2(direct 5)", 5)

	return

	/*

		first 0
		defer show()
		defer p(5)
		defer Println(5)
		second 10
		defer Println(direct 5)

		return result

	*/
}

func power() (result int) {
	fmt.Println("first", result)  // first 0 print hbe

	show := func(result int) {   // heap e show & result joma hoise
		result = result + 10
		fmt.Println("show", result)
	}

	defer show(result)  // arg pass korar karone show() te result er value 0, func call hbe pore

	result = 5

	p := func() {         // heap e p & result joma hoise
		fmt.Println("p", result)
	}

	defer p()  // arg pass kora hoini tai show() te result er value referrence, func call hbe pore

	defer fmt.Println("print1", result) // ekane result er value 5, print hobe pore

	result = 7

	fmt.Println("second", result)  // second 7 print hbe

	defer fmt.Println("print2(direct 5)", 5)

	return

	/*

		first 0
		defer show(0)
		defer p()
		defer println(5)
		second(10)
		defer println(direct 5)

		return result

	*/
}

func main() {

	punch := calculate()
	// hello := power()

	fmt.Println("calculate return", punch)
	// fmt.Println("power return", hello)
}
