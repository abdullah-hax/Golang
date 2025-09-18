
//-------------------Simulate the code--------------

package main

import "fmt"

func sum() {
	add(2, 4)
}
func add(a, b int) {
	fmt.Println(a + b)
}

func main() {

	sum()
	add(20, 40)

	var add = func(a, b int) {   // variable er vetor function raklam.
		c := a + b
		fmt.Println(c)
	}
	add(300, 500)   // variable k call krlam

}

func init() {
	fmt.Println("I'll be executed initially")
}

//-----------------------------------error asbe-------------------------

// package main

// import "fmt"

// func add(a,b int) {
// 	fmt.Println(a + b)
// }

// func main(){

// 	add(20,40)

// 	var add = func(a,b int) {
// 		c := a + b
// 		fmt.Println(c)
// 	}
// }





