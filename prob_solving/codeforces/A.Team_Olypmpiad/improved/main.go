package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	//// groups store the indices of each type
	var programmers, mathematicians, physicians []int // at first the slices are empty , not 0 at index 0

	for i := 0; i < n; i++ {
		var skill int
		fmt.Scan(&skill)

		if skill == 1 {
			programmers = append(programmers, i+1)
		}

		if skill == 2 {
			mathematicians = append(mathematicians, i+1)
		}

		if skill == 3 {
			physicians = append(physicians, i+1)
		}
	}

	//// find minimum team count
	lowest := min(len(physicians), min(len(programmers), len(mathematicians)))
	fmt.Println(lowest)

	//// print each team
	for i := 0; i < lowest; i++ {
		fmt.Println(programmers[i], mathematicians[i], physicians[i])
	}
}


/* 

	Problem link: https://codeforces.com/problemset/problem/490/A

*/