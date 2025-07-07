/* 

 problem link : https://codeforces.com/problemset/problem/50/A

*/

package main

import "fmt"

func main(){
	var a, b int
	fmt.Scan(&a, &b)

	var area = a * b

	if area % 2 == 0{
		fmt.Println(area / 2)
	} else {
		fmt.Println((area - 1) / 2)
	}
}