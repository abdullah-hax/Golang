package main

import "fmt"

func calculate() (result int) {
	fmt.Println("first", result)

	show := func(result int) {
		result = result + 10
		fmt.Println("show", result)
	}

	defer show(result)

	result = 5

	p := func(a int) {
		fmt.Println("p", a)
	}

	defer p(result)

	defer fmt.Println("print1", result)

	result = 10

	fmt.Println("second", result)

	defer fmt.Println("print2(direct 5)", 5)

	return

	/*

		first 0
		defer show(0)
		defer p(5)
		defer Println(5)
		second 10
		defer Println(direct 5)

		return result

	*/
}

func power() (result int) {
	fmt.Println("first", result)

	show := func(result int) {
		result = result + 10
		fmt.Println("show", result)
	}

	defer show(result)

	result = 5

	p := func() {
		fmt.Println("p", result)
	}

	defer p()

	defer fmt.Println("print1", result)

	result = 10

	fmt.Println("second", result)

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

	// punch := calculate()
	hello := power()

	// fmt.Println("calculate return", punch)
	fmt.Println("power return", hello)
}
