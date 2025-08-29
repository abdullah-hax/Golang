package main

import "fmt"

func main(){
	var testcase int
	fmt.Scan(&testcase)


	for i := 0; i < testcase; i++{
		var n int
		fmt.Scan(&n)

		for j := 1; j <= n; j++{

			if j == 2 {
				continue
			}
			fmt.Printf("%d ", j) 
		}
		fmt.Println("2")
		// see alternative
	}
}

/* === Alternative ===

    for j := 2; j <= n; j++{

			fmt.Printf("%d ", j) 
		}
		fmt.Println("1")

*/