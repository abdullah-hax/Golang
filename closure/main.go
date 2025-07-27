//------------ must simulate the code -----------


package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	var money = 100
	var age = 30

	fmt.Println("Age = ", age)

	var show = func(){
		money = money + a + p
		fmt.Println(money)
	}  

	return show
}

func call(){
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
	
}

func main(){
	call()
	
}

func init(){
	fmt.Println("=== Assalamualaikum ===")
}



/*

	== outer() jkn show k return krbe tar agei heap e show k joma raka hbe 
	karon go jane show k nxt e keo na keo call krbe, outer() jehetu muche gese tai show k kivabe access krbe?
	tai access krte parar jnno show er ekta copy heap e raka hbe & money k access krte parar jnno money er copy o rakha hbe.
	a , p jehetu global e ache tai ederk heap e rakar drkar nai.


*/


/* 

function er local variable show k access kra lagbe nxt e


*/