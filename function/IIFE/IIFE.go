package main

import "fmt"

func main() {

	//Anonymus function
	//Immediately Invoked(called) Function Expression (IIFE)
	func(a,b int){
		var res = a + b
		fmt.Println("Summation = ",res)
	} (7,8)
}

func init(){
	fmt.Println("I'll be executed first.")
}