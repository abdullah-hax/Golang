package mathlib

import "fmt"   // it is built-in library , so it can be imported directly.

func Sub(a, b int) {
	res := a - b
	fmt.Println(res)
}

var Ops = 20

func Victory() {
	fmt.Println(Ops)
}


// note: main package & main function chara execute hbe na