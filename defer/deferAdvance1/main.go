
package main

import "fmt"

func calculate() (result int){
	fmt.Println("first", result)   // first 0 print hobe

	show := func(){   // heap e show() & result joma hobe
		result = result + 10
		fmt.Println("defer", result)
	}

	defer show()    // result er value referrence, func call hobe pore

	result = 5
	fmt.Println("second", result)  // second 5 print hobe

	return
}

func calc() int {
	result := 0
	fmt.Println("first", result)   // first 0 print hobe

	show := func (){    // heap e show() & result joma hobe
		result = result + 10
		fmt.Println("defer", result)
	}

	defer show()  // result er value referrence, func call hobe pore

	result = 5
	fmt.Println("second", result)   // second 5 print hobe

	return result  // return evaluated to 5 (return evaluated bcz the func is unnamed)
}

func main() {
	a := calculate()
	fmt.Println("return", a)
	
	b := calc()
	fmt.Println("return", b)
}




/* 

	defer 2 kind
	  named & unnamed

	sobar seshe print hoi
	ekadik defer er khettre nijeder modde lifo structure follow kore
	closure create kore, ekta powerful tool  (ask : onno lang e ase kina, tool kina)

	aro kisu info takle dao




*/