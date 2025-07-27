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


	var test int
	fmt.Scan(&test)

	var slice = make([]string, test)
	// var arr [test]int ❌  (invalid array lenth, "lenth must be constant in array", but not in "Slice")

	for i := 0; i < test; i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice, "Ans : slice")
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

func xtra(){
/* 	hello := make([]string, 3) this means hello will take 3 strings
	hello := make([]int, 3) this means hello will take 3 int numbers
	hello := make([]byte, 3) this means hello will take 3 ascii characters
	hello := make([]rune, 3) this means hello will take 3 unicode characters
*/

	hello := make([]string, 3)
	hello[0] = "world"
	hello[1] = "world"
	hello[2] = "world"

	hello2 := make([]int, 3)
	hello2[0] = 1234
	hello2[1] = 1234
	hello2[2] = 1234

	hello3 := make([]byte, 3)
	hello3[0] = 'a'
	hello3[1] = '3'
	hello3[2] = 'z'

	hello4 := make([]rune, 3)
	hello4[0] = 'a'
	hello4[1] = '3'
	hello4[2] = '👍'

	fmt.Println(hello)
	fmt.Println(hello2)
	fmt.Printf("%c%c%c\n", hello3[0], hello3[1],hello3[2])
	fmt.Printf("%c\n", hello4)

}

func main() {
	intAndStringSlice()
	intAndStringArray()
	inputOneValueWithArrayOrSlice()
	inputOneValueWithVarible()

	xtra()
}

// In C , string = char type array
// In Go, string = onkgulo character er somosti but not array , int , float e jmn onkgulo digit thake tmni ekane onkgulo character thake ,
// go er string eo index ache but specefic character modify krte hole rune or byte e convert krte hbe.

// num   -> ei varible e 1ta number input nite parbe (fmt.Scan(&num))
// str     -> ei varible e 1ta string input nite parbe (fmt.Scan(&str))
// number array or string array  -> loop ghurie onkgulo number or onkgulo string input nite parbe ( fmt.Scan(&num[i]) / fmt.Scan(&str[i]) )  [ variable use kora jabena ❌ ]
// number slice or string slice  -> loop ghurie onkgulo number or onkgulo string input nite parbe ( fmt.Scan(&num[i]) / fmt.Scan(&str[i]) )   [ variable use kora jabe ✅ ]

// number = onkgulo digit mile ekta value
// string = onkgulo character mile ekta value
// array or slice = onkgulo string value or number value er somosti



/* Disadvantage of array :
    -> 3rd bracket[] er vetor var use kora jaina
	-> khali rakha jaina
*/