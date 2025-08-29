/*

 Problem link : https://codeforces.com/problemset/problem/71/A

 (1 ≤ row ≤ 100)
 letter lenth 1 to 100
 if letter lenth greater than 10 then organize the word

input :
4
word
localization
internationalization
pneumonoultramicroscopicsilicovolcanoconiosis
 
output :
word
l10n
i18n
p43s

*/

package main

import "fmt"

func method1() {
	var testcase int
	fmt.Scan(&testcase)

	for i := 0; i < testcase; i++ {
		slice := make([]string, 1)
		fmt.Scan(&slice[0])

		byt := []byte(slice[0])

		// var str string
		// fmt.Scan(&str)
		// byt := []byte(str)

		len := len(byt)

		if len > 10 {
			fmt.Printf("%c%d%c\n", byt[0], len-2, byt[len-1])
		} else {
			fmt.Println(slice[0])
		}
	}
}

func method2() {
	var testcase int
	fmt.Scan(&testcase)

	var slice = make([]string, testcase)
	for i := 0; i < testcase; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 0; i < testcase; i++ {
		byt := []byte(slice[i])
		len := len(byt)

		if len > 10 {
			fmt.Printf("%c%d%c\n", byt[0], len-2, byt[len-1])
		} else {
			fmt.Println(slice[i])
		}
	}
}

func main() {
	// You can use any one method in codeforces
	method1()
	method2()
}
