package main

import "fmt"

func sliceFromArray() {
	arr := [6]int{0, 1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println("Slice from array : ", slice)
}

func appendElementsToSlice() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	var anotherSlice = append(slice, 10, 20, 30)
	fmt.Println("Appended main slice : ", slice)           // [1, 2, 3, 4, 5]
	fmt.Println("Appended another slice : ", anotherSlice) // [1, 2, 3, 4, 5, 10, 20, 30]

	// Note : nill slice eo value append kora jabe

}

func removeElements(slice []int) {
	slice = append(slice[:2], slice[5:]...)
	fmt.Println("Slice after removal : ", slice)
}

func iterateOver(slice []int) {
	for i, v := range slice {
		fmt.Printf("index : %d, value : %d\n", i, v)
	}

}

func benefitOfSlice() {

	arr := [6]int{0, 1, 2, 3, 4, 5}             // number array
	arr2 := [3]string{"go", "slice", "closure"} // string array
	arr3 := [0]int{}                            // number array

	arr[2] = 16
	arr2[2] = "vogus"

	fmt.Println(arr)
	fmt.Println(arr2, arr3)

	// arr er index value modify kora jai, empty value neya jai
	/*

	   slice e xtra ja ache :
	   	=> size e varible use kra jai
	     	=> append() die add or remove kora jai
	      	=> onkgulo value k eksate neya jai  (slice form array topic)
	         ( fmt.Println(slice[2:5]),
	   	   arr te value gulo alada vabe nite hoi jmn
	   	   fmt.Println(str[2], str[3], str[4]) / fmt.Println(arr[2], arr[3], arr[4])


	*/

}

func main() {
	var slice = []int{0, 1, 2, 3, 4, 5}

	sliceFromArray()
	appendElementsToSlice()
	removeElements(slice)
	iterateOver(slice)
	benefitOfSlice()
}

// for _, v := range myslice {
//     fmt.Println(v)
// }

/*

func pass(slc []int)
 pass(slc)  // it is slice


func pass(slc int)
 pass(slc[2])  // it is not a slice , just a integer

*/
