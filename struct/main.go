package main

import "fmt"

type user struct {   // member variable or property
	name string
	age  int
}

func main() {
	var info1 user

	info1 = user{       // instance or object    // instantiate
		name: "abdullah",
		age:  20,
	}

	var info2 = user{       
		name: "abdur rahman",
		age:  30,
	}

	fmt.Println("Name : ", info1.name)
	fmt.Println("Age : ", info1.age)
	fmt.Println("Name : ", info2.name)
	fmt.Println("Age : ", info2.age)
}
