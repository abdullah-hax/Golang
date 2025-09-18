/* 

Problem link : https://codeforces.com/contest/2124/problem/A

*/

package main

import "fmt"

func main(){
	var testcase int
	fmt.Scan(&testcase)

	for i := 0; i < testcase; i++{
		var len int
		fmt.Scan(len)

		var slice = make([]int, len)

		for i := 0; i < len ; i++{
			fmt.Scan(&slice)
		}

		if len % 2 == 0{
			fmt.Println("YES")
		} else {
			
		}
	}
}