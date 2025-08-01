package main

import "fmt"

func normalMethod(){
	fmt.Println("Slices are used to store multiple values in a single variable of same data type.")
	slice1 := []int{2, 3, 5, 6, 7}  // 5ta value / 5ta int
	fmt.Println(slice1)
	fmt.Println("length : ", len(slice1))
	fmt.Println("capacity : ", cap(slice1))

	slice2 := []string{"Go", "slices", "are", "powerful"}  // 4ta value / 4ta string
	fmt.Println(slice2)
	fmt.Println("length : ", len(slice2))
	fmt.Println("capacity : ", cap(slice2))
}

func makeMethod(){
	slice := make([]string, 3, 5)
	fmt.Println(slice)
	fmt.Println("length : ", len(slice))
	fmt.Println("capacity : ", cap(slice))

	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)
	fmt.Println("length : ", len(slice2))
	fmt.Println("capacity : ", cap(slice2))
}

func nillSlice(){
	slice := []int{}
	slice2 := []int{1, 4, 7}
	fmt.Println("Emplty slice : ",slice)
	fmt.Println("Slice with initial values : ", slice2)
}

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

func main(){
	normalMethod()
	makeMethod()
	nillSlice()
	copyOneSliceToAnother()
	multiDSlice()

}

// array & slice er output [] er vetor hoi  ✅
// string er output [] er vetor hoina  ❌