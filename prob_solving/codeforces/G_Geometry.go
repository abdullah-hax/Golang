package main

import "fmt"

func main() {
	var width, height float64
	fmt.Scan(&width, &height)

	res := .5*width*height - 3.14159265359*height*height/4

	fmt.Printf("%.7f", res)
}

/*

	Problem link : https://codeforces.com/group/TyLNzrd4HT/contest/259963/problem/G


*/
