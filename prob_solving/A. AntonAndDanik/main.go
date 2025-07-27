/*

Problem link : https://codeforces.com/problemset/problem/734/A
*/

package main

import "fmt"

func main() {
	var n int
	var str string
	fmt.Scan(&n, &str)

	var a, d int
	for i := 0; i < n; i++ {
		if str[i] == 'A' {
			a++
		} else if str[i] == 'D' {
			d++
		}
	}

	if a > d {
		fmt.Println("Anton")
	} else if d > a {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}

/*  Alternative (More right code but lengthy)

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Scanf("\n")    // consume the leftover newline

	str := make([]byte, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &str[i]) // no space needed now, we handled newline
	}

	var a, d int
	for i := 0; i < n; i++{
		if str[i] == 'A'{
			a++
		} else if str[i] == 'D'{
			d++
		}
	}

	if a > d{
		fmt.Println("Anton")
	} else if d > a{
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
} 
*/

// For print the string : 	fmt.Println(string(str))
