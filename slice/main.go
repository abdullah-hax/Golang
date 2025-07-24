package main

import "fmt"

func main() {
	// arr := [6]string{"This", "is", "a", "Go", "interview", "questions"}
	// fmt.Println(arr)

	// var s = arr[1:4]   // [is a Go], ptr = r-1, len = 3, cap = 5
	// fmt.Println(s)

	// s1 := s[1:2]   // [a], ptr = r-2, len = 1, cap = 5
	// fmt.Println(s1)

	// fmt.Println(len(s1))   // 2 - 1 = 1
	// fmt.Println(cap(s1))  // 6 - 2(r-2) = 4

	// s2 := s[2:3]
	// fmt.Println(s2, len(s2), cap(s2))

	// slice literal
	// var s = []int{3, 4, 5, 6} // ptr = r-1, len = 4, cap = 4

	// s[0] = 10
	// fmt.Println("slice : ", s, len(s), cap(s))

	slice := []int{}
	slice = append(slice, 2, 3)
	anotherSlice := append(slice, 1)

	fmt.Println(slice)
	fmt.Println(anotherSlice)

}


// value : var str string
// array of strings : var str[100]string
// array of string slice : var str []string  // nill slice
// make : var str = make([]string, r)



