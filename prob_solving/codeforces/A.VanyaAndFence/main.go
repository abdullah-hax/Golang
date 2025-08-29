/* 

	Problem link : https://codeforces.com/problemset/problem/677/A

*/

package main

import "fmt"

func main() {
	var number , height int
	fmt.Scan(&number, &height)

	var arr = make([]int, number)

	for i := 0; i < number; i++{

		fmt.Scan(&arr[i])

	}

	var count int
	for i := 0; i < number; i++{
		count++

		if arr[i] > height{

			count++ 
		}
	}

	fmt.Println(count)
}