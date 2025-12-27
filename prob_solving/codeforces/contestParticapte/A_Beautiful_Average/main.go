package main

import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	var slice = make([]int, n)

	var greater = -1
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
		if slice[i] > greater {
			greater = slice[i]
		}
	}

	fmt.Println(greater)

}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	for i := 0; i < testcase; i++ {
		solve()
	}
}

/*

	Problem link : https://codeforces.com/contest/2162/problem/A

*/
