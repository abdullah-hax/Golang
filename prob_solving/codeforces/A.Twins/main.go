/*

	Problem link : https://codeforces.com/problemset/problem/160/A


*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	var i, j int

	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	var sum int
	for i = 0; i < n; i++ {
		sum = sum + arr[i]
	}

	var count int
	var exactValue int
	for i = n - 1; i >= 0; i-- {

		if exactValue <= sum/2 {
			exactValue = exactValue + arr[i]
			count++
		} else {
			break
		}
	}

	fmt.Println(count)
}
