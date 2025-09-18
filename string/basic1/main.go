package main

import (
	"fmt"
	"strings"
)

func defualtValue() {
	var x int                         // x is automatically 0
	var s string                      // s is ""
	var b bool                        // b is false
	fmt.Printf("%d,%s,%v\n", x, s, b) // Output: 0 "" false
}

func concatenate() {
	fmt.Println("After concatenation : ", "Hello "+"World")
}

func substring() {
	message := "Hello World!"
	fmt.Println("Substring of Hello World! : ", message[0:5]) // slice,  "Hello"
}

func compare() {
	var msg1 = "one"
	var msg2 = "two"

	fmt.Println("Equal?", msg1 == msg2)      // if eual then return true
	fmt.Println("Not equal?", msg1 != msg2)  // if not equal then return true
	fmt.Println(strings.Compare(msg1, msg2)) // if msg1 small then return -ve, if large then +ve else if equal then return 0
}

func checkingSubstring(msg string) {
	fmt.Println("checkingSubstring : ", strings.Contains(msg, "Wor"))
}

func upperLowerCase(msg string) {
	fmt.Println("Lowercase : ", strings.ToLower(msg))
	fmt.Println("Uppercase : ", strings.ToUpper(msg))
}


func main() {

	var msg = "Hello World!"

	defualtValue()
	concatenate()
	substring()
	compare()
	checkingSubstring(msg)
	upperLowerCase(msg)


}
