/* 

Problem link : https://codeforces.com/problemset/problem/59/A


*/


package main

import (
	"fmt"
	"strings"
)


func main() {
	var str string
	fmt.Scan(&str)  // str er vetor HoUse ache.

	var upr, lwr int
	r := []rune(str)  // r er vetoreo HoUse ache tbe vanga vanga obstai ache. Ar str er vetor jura obstai ache.

	length := len(r)

	for i := 0; i < length; i++{
		if r[i] < 'a'{     // or r[i] < 97
			upr++
		} else {
			lwr++
		}
	}

	if upr > lwr {
		str = strings.ToUpper(str)
	} else {
		str = strings.ToLower(str)
	}

	fmt.Println(str)
}