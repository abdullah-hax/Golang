/*

	Problem link : https://codeforces.com/contest/2149/problem/A

*/

package main

import "fmt"

func solve() {
	var n int
	var i int
	fmt.Scan(&n)
	slice := make([]int, n)

	for i = 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	var count = 0
	var minusOne = 0
	for i = 0; i < n; i++ {
		if slice[i] == 0 {
			count++
		}
		if slice[i] == -1 {
			minusOne++
		}
	}

	if minusOne%2 != 0 {
		count = count + 2
	}

	fmt.Println(count)

}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	for i := 0; i < testcase; i++ {
		solve()
	}
}

/*



if zero found :
	count++

minusOne++
if minusOne % 2 != 0 :
	count 2 barbe


*/
