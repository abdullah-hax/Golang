/*
Problem link : https://codeforces.com/problemset/problem/263/A

declare arr[5][5]
read arr[5][5]

for i = 0 to i = 4 :
   for j = 0 to j = 4:
     if(a[i][j] == 1)
       result = abs(2-i) + abs(2-j)
       return 0

*/

package main

import (
	"fmt"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var arr [5][5]int
	var i, j int
	var row, col int

	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			fmt.Scan(&arr[i][j]) // don't use Scanln ❌

			if arr[i][j] == 1 {
				row = i
				col = j
			}

		}
	}

	result := abs(2-row) + abs(2-col)

	fmt.Println(result)
}

/*

| Feature                        | `abs(x int) int` (custom) | `math.Abs()` (built-in)               |
| ------------------------------ | ------------------------- | ------------------------------------- |
| **Works with `int`**           | ✅ Yes                     | ❌ No (requires `float64`)             |
| **Simple & readable**          | ✅ Yes                     | 😐 No (needs type conversion)         |
| **Faster (no conversion)**       | ✅ Yes                     | ❌ Slightly slower due to conversions  |
| **Safer (no float precision)**     | ✅ Yes                     | ❌ Might lose precision for large ints |
| **Best for integer logic**        | ✅ Yes                     | ❌ No                                  |
| **Use in float math**           | ❌ No                      | ✅ Yes                                 |


*/
