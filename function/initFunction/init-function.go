package main

import "fmt"

var a = 10

func main() {
	fmt.Println(a)
}

func init() {
	fmt.Println(a)
	a = 50
}


// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello init function.")
// }

// func init() {
// 	fmt.Println("I am the first function that is executed initially.")
// }
