package main

import "fmt"

var arr2 = [4]string{"Shafi", "is", "a", "businessman"} // In C, 4 means 4 char, but in Go , 4 means 4 char/ 4 words/ 4 sentences.
var arr3 = [3]string{"I", "am", "baker"}

func main() {
	var arr = [4]int{3, 4, 5, 6}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

/*

    *** array ***

  array = onkgulo variable er somosti
  array = fuler mala = onkgulo fuler somosti

*/

/*

In C , declaration & assignment same

int arr[3] = {1, 2, 3};  // declaration + initialization
arr[0] = 10;             // assignment


In Go , assignment is same as C , but declaration is different

arr = [3]int{4, 5, 6} // declaration + initialization
arr[0] = 10            // assignment


*/