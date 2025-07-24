package main

import "fmt"

func main() {
	// //part 1
	// num := 12345

	// var str = fmt.Sprint(num)

	// r := []rune(str)

	// r[0] = '8'
	// fmt.Println(string(r))

	//part 2
	var str2 = "12345"
	fmt.Println(str2[0], str2[1])
	fmt.Printf("%c  %c", str2[0], str2[1])

	fmt.Printf("\n")

	var str3 = "abcde"
	fmt.Println(str3[0], str3[1])
	fmt.Printf("%c  %c", str3[0], str3[1])

}

/* part 1 :
rune = unicode
byte = ascii code

rune = 97 likle se etake unicode hisebe dkbe , 97 er character = a
byte = 8 likle se etake ascii code hisebe dkbe , 8 er character = backspace

-> ascii is a subset of unicode
-> Unicode preserves all ASCII values in the first 128 code points (U+0000 to U+007F).

part 2 :
string er vetor ja ache prottekta ekekta character , index er value change kora jaina ,
change krte hle rune or byte e convert krte hbe.
*/
