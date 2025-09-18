/*

Problem link : https://codeforces.com/problemset/problem/977/A

*/

package main

import "fmt"

func main() {
	var num, k int
	fmt.Scan(&num, &k)

	for i := 0; i < k; i++ {
		if num%10 == 0 {
			num = num / 10
		} else {
			num = num - 1
		}
	}

	fmt.Println(num)
}
