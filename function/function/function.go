// package main

// import "fmt"

// func add (num1 int , num2 int ){
// 	sum := num1 + num2

// 	fmt.Println(sum)
// }

// func main() {
// 	var a = 20
// 	var b = 5

// 	add(a, b)
// 	add(5, 8)
// }
//-------------------------------------------------------
// package main

// import "fmt"

// func add(num1 int, num2 int) int {
// 	var sum = num1 + num2
// 	return sum
// }

// func get(num1 int, num2 int) (int, int) {

// 	var sum = num1 + num2
// 	var multi = num1 * num2

// 	return sum, multi
// }

// func main() {
// 	var a = 10
// 	var b = 20

// 	sum := add(a, b)
// 	fmt.Println(sum)

// 	var m, n = get(a, b)

// 	fmt.Println(m, n)
// 	// fmt.Println(n)

// }
//---------------------------------------------
package main

import "fmt"

func printsomething() {
	fmt.Println("Education must be free!")
}

func print2things(name string) { // compiler RAM teke ekta block nibe(RAM e onk block take), block er nam dibe name, name block e Abdullah store krbe
	fmt.Println("This is the month of Ramadan, ", name) // long string ta print krbe & name block e j string ta ache ta print krbe
}
func main() {

	printsomething()
	print2things("Abdullah")
}

//------------------------------------------
// package main

// import "fmt"

// var(
// 	a = 10
// 	b = 20
// )

// func add(x int , y int) {
// 	res := x + y
// 	printNum(res)
// }

// func main() {
// 	add(a, b)
// }

// func printNum(num int) {
// 	fmt.Println(num)
// }