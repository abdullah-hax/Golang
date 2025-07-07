/* 

 Problem link : https://codeforces.com/contest/2072/problem/A

*/


package main

import "fmt"

func abs(n int) int {
	if n < 0{
		return -n
	}
	return n
}

func main(){
	var testcase int
	fmt.Scan(&testcase)

	var op , i int

	var n = make([]int, testcase)
	var k = make([]int, testcase)
	var p = make([]int, testcase)
	
	for i = 0; i < testcase; i++{
		fmt.Scan(&n[i], &k[i], &p[i])
	}

	for i = 0; i < testcase; i++{
		if abs(k[i]) % p[i] == 0{
			op = abs(k[i]) / p[i]
		} else {
			op = abs(k[i]) / p[i] + 1
		}
		
		if op <= n[i]{
			fmt.Println(op)
		} else {
			fmt.Println("-1")
		}
	}

}