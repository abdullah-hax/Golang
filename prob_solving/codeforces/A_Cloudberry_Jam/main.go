/* 

Problem link : https://codeforces.com/problemset/problem/2086/A

*/

package main

import "fmt"

func main(){

	var testcase int 
	fmt.Scan(&testcase)

	for i := 0 ; i < testcase; i++{
		var jar int 
		fmt.Scan(&jar)

		berries := ((3 + 1) * jar) / 2
		fmt.Println(berries)
	}

}