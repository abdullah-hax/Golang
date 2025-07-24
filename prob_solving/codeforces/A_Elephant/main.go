/*

Problem link : https://codeforces.com/contest/617/problem/A


*/

package main

import "fmt"

func main() {
	var n, res int
	fmt.Scan(&n)

	if n < 5 {
		res = 1
	} else {
		if n%5 == 0 {
			res = n / 5
		} else {
			res = n/5 + 1
		}

	}

	fmt.Println(res)
}
