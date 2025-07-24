package main

import "fmt"

func intAndStringSlice() {
	fmt.Println("Taking input using slice (take 4 input) :")
	var slice1 = make([]int, 4)
	var slice2 = make([]string, 4)

	// fmt.Scan(&slice) ❌
	for i := 0; i < 4; i++ {
		fmt.Scan(&slice1[i]) // 4ta integer input nibe
	}
	fmt.Println(slice1)

	// fmt.Scan(&slice) ❌
	for i := 0; i < 4; i++ {
		fmt.Scan(&slice2[i]) // whitespace chara 4 ta string input nibe
	}
	fmt.Println(slice2)
}

func intAndStringArray() {
	fmt.Println("Taking input using slice (take 3 input) :")
	var arr1 [3]int    // var arr []int ❌
	var arr2 [3]string // var str []string ❌

	for i := 0; i < 3; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Println(arr1)

	for i := 0; i < 3; i++ {
		fmt.Scan(&arr2[i])
	}
	fmt.Println(arr2)
}

func inputOneValueWithArrayOrSlice() {

	fmt.Println("Taking input one string : ")
	slice1 := make([]string, 1)
	// var slice1 [1]string
	fmt.Scan(&slice1[0])
	fmt.Println(slice1)

	fmt.Println("Taking input one int : ")
	slice2 := make([]int, 1)
	// var slice2 [1]int
	fmt.Scan(&slice2[0])
	fmt.Println(slice2)
}

func input() {
	fmt.Println("Which is better, array or slice? : (take input first)")
	var test int
	fmt.Scan(&test)

	var arr = make([]string, test)
	// var arr [test]int ❌  (invalid array lenth, "lenth must be constant in array", but not in "Slice")

	for i := 0; i < test; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr, "Ans : slice")
}

func inputOneValueWithVarible() {
	fmt.Println("Taking input one string : ")
	var str string
	fmt.Scan(&str)
	fmt.Println(str)

	fmt.Println("Taking input one int : ")
	var num int
	fmt.Scan(&num)
	fmt.Println(num)
}

func main() {
	intAndStringSlice()
	intAndStringArray()
	inputOneValueWithArrayOrSlice()
	input()
	inputOneValueWithVarible()
}

// In C , string = char type array
// In Go, string = onkgulo character er somosti but not array , int , float e jmn onkgulo digit thake tmni ekane onkgulo character thake ,
// go er string eo index ache but modify krte hole rune or byte e convert krte hbe.



// num   -> ei varible e 1ta number input nite parbe (fmt.Scan(&num))
// str     -> ei varible e 1ta string input nite parbe (fmt.Scan(&str))
// number array or string array  -> loop ghurie onkgulo number or onkgulo string input nite parbe ( fmt.Scan(&num[i]) / fmt.Scan(&str[i]) )  [ variable use kora jabena ❌ ]
// number slice or string slice  -> loop ghurie onkgulo number or onkgulo string input nite parbe ( fmt.Scan(&num[i]) / fmt.Scan(&str[i]) )   [ variable use kora jabe ✅ ]

