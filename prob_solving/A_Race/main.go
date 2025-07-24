/*

Problem link : https://codeforces.com/contest/2112/problem/A

*/
/*

var j = 0
i = 0 to 100 :
  if |b - x| < |a - x| : arr[j++] = i
  else:
     cotinue

j = 0
i = 0 to 100 :
  if |b - y| < |a - y| : arr[j++] = i
  else:
     cotinue

or,

var j = 0
first = x - (a - x - 1)
last = x + (a - x - 1)

for i := first; i <= last; i++ :
    arr[j++] = i

j = 0
first = y - (a - y - 1)
last = y + (a - y - 1)

for i := first; i <= last; i++ :
    arr[j++] = i

*/

package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main(){
	var testcase int
	fmt.Scan(&testcase)

	var i int
	for i = 0; i < testcase; i++{
		var alice, x, y int
		fmt.Scan(&alice, &x, &y)

		var first = x - (abs(alice - x) - 1)
		var last = x + (abs(alice - x) - 1)

		var i = 0
		var j = 0
		var arr1 = make([]int, (last - first + 1))

		for i = first; i <= last; i++{
			arr1[j] = i
			j++
		}

		first = y - (abs(alice - y) - 1)
		last = y + (abs(alice - y) - 1)
		j = 0
		var arr2 = make([]int, (last - first + 1))

		for i = first; i <= last; i++{
			arr2[j] = i
			j++
		}

		var count bool
		for i = 0; i != len(arr1); i++{

			for j = 0; j != len(arr2); j++{

				if arr1[i] == arr2[j]{
					fmt.Println("YES")
					count = true
					break
				}

			}
			if count == true{
				break
			}
		}

		if count == false {
			fmt.Println("NO")
		}
	}
}