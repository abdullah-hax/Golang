/* 

This is project of Alim Shahriar

*/

package main

import "fmt"

func main() {
	var a, c float32
	var b string
	fmt.Scan(&a, &b, &c)

	if b == "+" {
		fmt.Print(a + c)
	} else if b == "-" {
		fmt.Print(a - c)
	} else if b == "*" {
		fmt.Print(a * c)
	} else if b == "/" {

		fmt.Print(a / c)

	}
}