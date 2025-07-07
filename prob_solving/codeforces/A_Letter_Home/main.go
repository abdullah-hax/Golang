/*

Problem link : https://codeforces.com/problemset/problem/2121/A


 read testcase(12)
 for i = 1 to (12) :
    read n
    read start
    arr[n]
    for j= 0 to n-1 :
       read arr[j]

       |start - arr[0]| + |arr[n-1] - arr[0]|  or
       |start - arr[n-1]| + |arr[n-1] - arr[0]|


*/

package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	var i, res int
	var arr_store = make([]int , testcase)

	for i = 0; i < testcase; i++ {
		var n, start int
		fmt.Scan(&n, &start)

		var arr = make([]int , n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}

		if abs(start-arr[0]) < abs(start-arr[n-1]) {
			res = abs(start-arr[0]) + abs(arr[n-1]-arr[0])
		} else {
			res = abs(start-arr[n-1]) + abs(arr[n-1]-arr[0])
		}

		arr_store[i] = res
	}

	for i = 0; i < testcase; i++ {
		fmt.Println(arr_store[i])
	}

}


/* === Easy Method ===


package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	var i, res int

	for i = 0; i < testcase; i++ {
		var n, start int
		fmt.Scan(&n, &start)

		var arr = make([]int , n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}

		if abs(start-arr[0]) < abs(start-arr[n-1]) {
			res = abs(start-arr[0]) + abs(arr[n-1]-arr[0])
		} else {
			res = abs(start-arr[n-1]) + abs(arr[n-1]-arr[0])
		}

		fmt.Println(res)
	}

} */

