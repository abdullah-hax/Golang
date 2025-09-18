/*

 Problem link : https://codeforces.com/contest/2123/problem/A

*/

package main

import (
	"fmt"
)

func solve1() {
	var n int
	fmt.Scan(&n)

	var cnt = make([]int, 4)

	for i := 0; i < n; i++ {
		cnt[i%4]++ // array size e variable raka jaina, tbe index e var raka jai
	}

	if cnt[0] == cnt[3] && cnt[1] == cnt[2] {
		fmt.Println("Bob")
	} else {
		fmt.Println("Alice")
	}
}

func solve2() {
	var n int
	fmt.Scan(&n)

	if n%4 == 0 {
		fmt.Println("Bob")
	} else {
		fmt.Println("Alice")
	}
}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	for i := 0; i < testcase; i++ {
		solve1()
		solve2() // Easy method
	}
}

/* // === Using slice (taking all input together) ===

package main

import "fmt"

func main(){
	var testcase int
	fmt.Scan(&testcase)

	var slice = make([]int, testcase)
	for i:= 0; i < testcase; i++{
		fmt.Scan(&slice[i])
	}

	for i:= 0; i < testcase; i++{
		if slice[i] % 4 == 0 {
			fmt.Println("Bob")
		} else {
			fmt.Println("Alice")
		}
	}
}

*/
