package main

import (
	"fmt"
)

func solve() {
	var x int
	fmt.Scan(&x)

	min := 9
	for x > 0 {

		if x > 0 {
			c := x % 10
			if c < min {
				min = c
			}
			x = x / 10
		}

	}
	fmt.Println(min)

}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	var i int
	for i = 0; i < testcase; i++ {
		solve()
	}
}

/*

left x == 0
x = abs(x)

*/
