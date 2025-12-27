package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	var n int
	var i int
	var programmer, mathmatician, physician int


	fmt.Scan(&n)
	slice := make([]int, n)

	for  i = 0; i < n; i++ {
		fmt.Scan(&slice[i])

		if slice[i] == 1 {
			programmer++
		} else if slice[i] == 2 {
			mathmatician++
		} else if slice[i] == 3 {
			physician++
		}
	}

	lowest := min(programmer, min(mathmatician, physician))
	fmt.Println(lowest)


	var j, k, temp int
	i = 0

	for temp = 0; temp < lowest; temp++ {
		for ; i < n; i++{
			if slice[i] == 1 {
				fmt.Print(i + 1, " ")
				i++
				break
			}
		}

		for ; j < n; j++{
			if slice[j] == 2 {
				fmt.Print(j + 1, " ")
				j++
				break
			}
		}

		for ; k < n; k++{
			if slice[k] == 3 {
				fmt.Println(k + 1, " ")
				k++
				break
			}
		}
	}

}


/* 

	Problem link: https://codeforces.com/problemset/problem/490/A

*/


/* 


var j, k, temp int
	

	for temp = 0; temp < lowest; temp++ {
		for i = 0 ; i < n; i++{
			if slice[i] == 1 {
				fmt.Print(i + 1, " ")
				slice[i] = 0  // Mark used
				break
			}
		}

		for j = 0; j < n; j++{
			if slice[j] == 2 {
				fmt.Print(j + 1, " ")
				slice[j] = 0
				break
			}
		}

		for k = 0; k < n; k++{
			if slice[k] == 3 {
				fmt.Println(k + 1, " ")
				slice[k] = 0
				break
			}
		}
	}

*/