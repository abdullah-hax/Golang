package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Scanf("\n") // consume the leftover newline

	str := make([]string, n)
	for i := 0; i < n; i++ {
		// fmt.Scanf("%c", &str[i]) // no space needed now, we handled newline
		fmt.Scan(&str[i])
	}

	fmt.Println(str)
	// fmt.Println(string(str))
}
