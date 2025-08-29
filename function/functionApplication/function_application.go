package main 

import "fmt"

func printWelcomeMessage() {
	fmt.Println("Assalamualaikum.\nWelcome to the application")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name : ")
	fmt.Scanln(&name)
	return name
}

func get2numbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter first number : ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number : ")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int , num2 int) int {
	var sum = num1 + num2
	return sum
}

func displayResult(name string, sum int) {
	fmt.Println("Hello, ", name)
	fmt.Println("Summation = ", sum)
}

func printGoodbye() {
	fmt.Println("Thank you for using the application")
	fmt.Println("Allah Hafiz")
}
func main () {
	printWelcomeMessage()
	var name = getUserName()
	var num1, num2 = get2numbers()
	var sum = add (num1, num2)
	displayResult(name, sum)
	printGoodbye()
}



// Without Function

// package main

// import "fmt"

//  func main() {
// 	// Print welcome message
// 	fmt.Println("Assalamualaikum.\nWelcome to the application.")

// 	// Get username form user
// 	// var name string
// 	name := ""
// 	fmt.Println("Enter your name : ")
// 	fmt.Scanln(&name)   // ekane space deya jaccena

// 	// Get number form user
// 	var num1 int
// 	var num2 int
// 	fmt.Println("Enter first number : ")
// 	fmt.Scanln(&num1)
// 	fmt.Println("Enter second number : ")
// 	fmt.Scanln(&num2)
// 	var sum = num1 + num2

// 	// Display result
// 	fmt.Println("Hello, ", name)
// 	fmt.Println("Summation = ", sum)

// 	// Print a goodbye message
// 	fmt.Println("Thank you for using the application.")
// 	fmt.Println("Allah Hafiz")
//  }