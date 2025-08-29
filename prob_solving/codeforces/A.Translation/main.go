package main

import "fmt"

func main() {

	var str1 string
	var str2 string

	fmt.Scan(&str1)
	fmt.Scan(&str2)

	length := len(str1)
	runes := []rune(str1)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	if string(runes) == str2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
