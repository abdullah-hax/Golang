package main

import "fmt"

func a(){
	i := 0
	fmt.Println("first", i)

	defer fmt.Println("second", i)

	i++

	fmt.Println("third", i)

	defer fmt.Println("fourth", i)

	return
}

func main() {
	a()
}


/* Output :

	first 0
	third 1
	fourth 1
	second 0

*/