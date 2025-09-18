/*
	Problem link : https://codeforces.com/problemset/submit

*/

package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)

	var str string
	fmt.Scan(&str)

	runes := []rune(str)

	for i := 0; i < t; i++ {
		for j := 0; j < (len(runes) - 1); j++ {
			if runes[j] == 'B' && runes[j+1] == 'G' {
				runes[j], runes[j+1] = runes[j+1], runes[j]
				j++
			}
		}
	}

	fmt.Println(string(runes))
}

/*  Alternative (More right code but lengthy)

package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)

	fmt.Scanf("\n")

	str := make([]byte, n)

	for i := 0; i < n; i++{
		fmt.Scanf("%c", &str[i])
	}

	for i := 0; i < t; i++ {
		for j := 0; j < (len(str) - 1); j++ {
			if str[j] == 'B' && str[j+1] == 'G' {
				str[j], str[j+1] = str[j+1], str[j]
				j++
			}
		}
	}

	fmt.Println(string(str))
} */
