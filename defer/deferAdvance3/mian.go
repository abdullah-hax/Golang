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

	result = 10

	defer fmt.Println("print1", result)

	fmt.Println("second", result)

	defer fmt.Println("print2(direct 5)", 5)

	return
}

func main() {
	punch := calculate() 
	fmt.Println("return", punch)
}



/* 

first 0
defer show[ref - result = 0, arg pass = 0]
defer p[ref - result = 5, arg pass = 5]
defer Println[ref - result]
second 10
defer Println[direct 5]

return

*/