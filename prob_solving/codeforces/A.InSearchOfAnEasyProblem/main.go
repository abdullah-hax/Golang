package main

import "fmt"

func main() {
	var people int
	fmt.Scan(&people)

	slice := make([]int , people)

	for i := 0; i < people; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 0; i < people; i++{
		if slice[i] == 1{
			fmt.Println("Hard")
			return
		}
	}

	fmt.Println("Easy")
}