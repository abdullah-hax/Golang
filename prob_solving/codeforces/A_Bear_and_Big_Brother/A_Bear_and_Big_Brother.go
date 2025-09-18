/*

Problem link : https://codeforces.com/problemset/problem/791/A

declare limak, bob

read limak, bob


for i = 1 to i = 10:
    limak = 4*3^i
    bob = 7*2^i

    if(limak > bob)
    break

print(i)

=== alternative & best ===

for i = 1 to i = 10:
    limak = limak * 3
    bob = bob * 2

    if(limak > bob)
    break

*/


package main

import "fmt"

func main(){
	var limak, bob int
	var yrs int
	fmt.Scanln(&limak, &bob)

	for limak <= bob {
		limak = limak * 3
		bob = bob * 2

		yrs++
	}

	fmt.Println(yrs)

}




/* === Alternative (but lengthy) ===

package main

import (
	"fmt"
	"math"
)

func main(){
	var limak, bob, i, a, b int
	fmt.Scanln(&a, &b)

	for i = 1; i <= 10; i++{
		limak = a * int(math.Pow(3, float64(i)))
		bob = b * int(math.Pow(2, float64(i)))

		if limak > bob{
			break
		}
	}

	fmt.Println(i)
}

// math.Pow returns float64, ei function use krle sobaik float64 banate hbe 

*/