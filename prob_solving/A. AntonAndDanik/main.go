/*

Problem link : https://codeforces.com/problemset/problem/734/A
*/

package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	var a, d int
	for i := 0; i < len(str); i++{
		if str[i] == 'A'{
			a++
		} else if str[i] == 'D'{
			d++
		}
	}

	if a > d{
		fmt.Println("Anton")
	} else if d < a{
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}


/* 
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	str := make([]byte, n)

	for i := 0; i < n; i++ {
		fmt.Scanf(" %c", &str[i]) // space before %c to skip newline/space
	}

	fmt.Println(string(str))
}
 */