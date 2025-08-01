package main

import "fmt"

func main() {
	// pointer = address of memory (ram)

	x := 20

	p := &x // suppose &x = 15 , so p = 15

	fmt.Println("Address: ", p)
	fmt.Println("Value at this address: ", *p) // * = value at address,, p te j address ache sei address er value print koro


	fmt.Println("      2nd part     ")


	fmt.Println("x = ", x)

	*p = 300          // vlaue at address p = value at address x 
	fmt.Println("x = ", x)
}


