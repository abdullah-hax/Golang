package main

import "fmt"

func main() {
	arr := [6]int{0, 1, 2, 3, 4, 5}   // number array
	arr2 := [3]string{"go", "slice", "closure"}  // string array
	arr3 := [0]int{}   // number array

	arr[2] = 16
	arr2[2] = "vogus"

	fmt.Println(arr)
	fmt.Println(arr2, arr3)

}

// arr er index value change kora jai, empty value neya jai
/*

slice e xtra ja ache :
	=> size e varible use kra jai	
  	=> append() die add or remove kora jai
   	=> onkgulo value k eksate neya jai
      ( fmt.Println(slice[2:5]), 
	   arr te value gulo alada vabe nite hoi jmn
	   fmt.Println(str[2], str[3], str[4]) / fmt.Println(arr[2], arr[3], arr[4])


*/
