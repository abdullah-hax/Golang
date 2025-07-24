package main

import "fmt"

type user struct {
	Name    string
	Age     int
	Salary  float64
	FavFood string
}

func main() {
	var info = user{
		Name:    "Abdullah",
		Age:     20,
		Salary:  0,
		FavFood: "Dates",
	}

	var p = &info

	fmt.Println(*p)
}
