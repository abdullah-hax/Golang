package main

import "fmt"

func add(x *int) int {
	*x = 20 // dereference of x / value at x
	return *x
}

func main() {
	var a = 10
	fmt.Println(a)

	b := add(&a)
	fmt.Println(a)
	fmt.Println(b)

}

/*

  	in c, * goes before the variable
    => int *a;

  	in go, * goes before the data type
    => var a *int

	so go cleared some confusion
	example : int* x, y; => only x is a pointer, y is not.
	         var x, y * int => both are pointer

*/
