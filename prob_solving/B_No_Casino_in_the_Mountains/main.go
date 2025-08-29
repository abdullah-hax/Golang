package main

import (
	"fmt"
	"os"
)

func solve() {
	var n int
	fmt.Scan(&n)
	var arr = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var count = 0
	var k int
	fmt.Scan(&k)

	
	for i := 0; i < n; i++ {
		if i == k {
			continue
		} else {
			if arr[i] == 0 {
				count++
			} else {
				fmt.Println("0")
				os.Exit(0)
			}
		}

	}

	if count > 0{

		a := n / (k + 1)
	
		b := n - a
		c := b / k
	
		fmt.Println(c)
	}

}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	var i int
	for i = 0; i < testcase; i++ {
		solve()
	}
}
