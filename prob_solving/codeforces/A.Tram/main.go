/*
Problem link : https://codeforces.com/contest/116/problem/A
*/
package main

import "fmt"

func main() {

	var testcase int
	fmt.Scan(&testcase)

	var count int
	var arr = make([]int, testcase)

	for i := 0; i < testcase; i++ {
		var minus, plus int
		fmt.Scan(&minus, &plus)

		count = count + plus - minus
		arr[i] = count
	}

	for i := 0; i < testcase; i++ {
		for j := i + 1; j < testcase; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr[0])

}
