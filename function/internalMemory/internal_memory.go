/*

	-> compilation phase ei compiler sb kisu ke binary te convert kore code segment e raka shuru kore
	-> code segment er jinisgulo binary file teke ase ,, egulo read-only , change kora jaina
	variable change kora jai tai var k code segment e rakena

*/

/*

   
	go build main.go => compilation phase start => compiler binary file create kore
	./main           => execution phase start => 'computer' binary code k execute kore
	go run main.go  => compile , execute 2 tai kore

	***** go build main.go => binary file(.exe file)(code segment) ******
)
	a = 10  // eta const
	call = func() {...}
	add = func() {...}
	main = func() {...}
	init = func() {...}
	egulo k binary hisebe rakhe, evabe english hisebe rake na.

*/

// real simulation koro (trpr abr agerta koro, 2tai prai same, but agerta ektu fast)

package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x, y int) {
		var z = x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}

func main() {
	call()

	fmt.Println(a)
}

func init() {
	fmt.Println("Assalamualaikum")
}

/*

***variable definition***
    var a = 20

***funtion definition***
	func main() {
		call()  ** eta funtion noi , eta hlo call **

		fmt.Println(a)
	}


***funtion definition***
	func add(a, b int)(int, int){
		result := a + b
		return a, result
	}

	add(45, 5)
	add(2, 6)  ** {} er bairer add 2 ta definition er under e pore na , era hlo call

***funtion definition***
	func call() {
		var add = func(a int , b int){
			fmt.Println(a + b)
		}

	** add variable er vetore jeta raka hoice setao 'funtion definition'.

		add(3, 4)
		add(5, 6)   ** era hlo call
	}

*/
