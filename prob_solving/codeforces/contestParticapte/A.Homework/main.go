package main

import "fmt"

func solve() {
	var alen int
	fmt.Scan(&alen)
	var a string
	fmt.Scan(&a)

	var blen int
	fmt.Scan(&blen)
	var b string
	fmt.Scan(&b)

	var c string
	fmt.Scan(&c)

	// cRune := []rune(c)
	// bRune := []rune(b)

	for i := 0; i < len(c); i++ {
		if c[i] == 'V' {
			a = string(b[i]) + a  // String er sate kebol string i jug krte parbe
		} else if c[i] == 'D' {
			a = a + string(b[i])
		}
	}

	fmt.Println(a)

}

func main() {
	var testcase int
	fmt.Scan(&testcase)

	for i := 0; i < testcase; i++ {
		solve()
	}
}
