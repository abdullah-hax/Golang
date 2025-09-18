

/*

Problem link : https://codeforces.com/problemset/problem/110/A

*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	var count int
	str := fmt.Sprint(num)
	for i := 0; i < len(str); i++ {
		if str[i] == '4' || str[i] == '7' {
			count++
		}
	}

	if count == 4 || count == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
