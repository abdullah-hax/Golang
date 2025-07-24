package main

import "fmt"

func copyOneSliceToAnother(){
	var src = []int{0, 1, 2, 3, 4, 5}
	var dest = make([]int, len(src)) 
	copy(dest, src)
	fmt.Println("Copied slice : ", dest)
}

func multiDSlice(){
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Multidimensional Slice : ", matrix)
}

func removeElements(slice []int){
	slice = append(slice[:2], slice[5:]...)
	fmt.Println("Slice after removal : ", slice)
}

func iterateOver(slice []int){
	for i, v := range slice{
		fmt.Printf("index : %d, value : %d\n", i, v)
	}
	
}

func main(){
	var slice = []int{0, 1, 2, 3, 4, 5}
	copyOneSliceToAnother()
	multiDSlice()
	removeElements(slice)
	iterateOver(slice)
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