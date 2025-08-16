/*

	Problem link : https://codeforces.com/contest/2131/problem/A

*/

package main

import "fmt"

var i, count int

func solve() {
	var n int
	fmt.Scan(&n)
	arr1 := make([]int, n)
	arr2 := make([]int, n)

	for i = 0; i < n; i++ {
		fmt.Scan(&arr1[i])
	}

	for i = 0; i < n; i++ {
		fmt.Scan(&arr2[i])
	}

	count = 0

	for i = 0; i < n; i++ {
		if arr1[i] > arr2[i] {
			count = count + arr1[i] - arr2[i]
		}
	}

	count++
	fmt.Println(count)

}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	for i := 0; i < testcase; i++ {
		solve()
	}
}
