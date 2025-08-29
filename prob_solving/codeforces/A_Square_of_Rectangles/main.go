/* 

Problem link : https://codeforces.com/problemset/problem/2120/A

*/

package main

import "fmt"

func main() {

	var testcase int
	fmt.Scan(&testcase)

	var i = 0
	for i < testcase {

		var l1, l2, l3, b1, b2, b3 int
		fmt.Scan(&l1, &b1, &l2, &b2, &l3, &b3)

		if ((l1 == l2 && l2 == l3) && l1 == b1+b2+b3) || ((b1 == b2 && b2 == b3) && b1 == l1+l2+l3) || 
		  ((l1 == b1+b2 && b2 == b3) && l1 == l2+l3) || ((b1 == l1+l2 && l2 ==l3) && b1 == b2+b3) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		i++
	}

}
