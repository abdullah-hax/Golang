//----------simulate the code-------------


package main

import "fmt"

func call() func (x,y int){
	return add
}

func add(x,y int) {
	res := x + y
	fmt.Println(res)
}

func main(){
	sum := call()

	sum(4,3)
}