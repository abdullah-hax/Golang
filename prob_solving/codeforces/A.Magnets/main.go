/*

	Problem link: https://codeforces.com/problemset/problem/344/A

*/

package main

import "fmt"

var i int

var count int

func main() {
	var n int
	fmt.Scan(&n)

	var slice = make([]int, n)

	for i = 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	for i = 0; i < n-1; i++ {

		if slice[i]+slice[i+1] == 11 {
			count++
		}
	}

	fmt.Println(count + 1)
}

// NB : Go doesn't allows multiple assignmets.
