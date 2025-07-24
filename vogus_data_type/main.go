package main


/*

  1) Good software engineer
     -> optimize time
	 -> optimize memory

  2) The name of go operating system (mini operating system) itself = go runtime
  3) format specifier of bool => %v
     format specifier of showing type (of variable) => %T
	 format specifier of rune => %c

  4) Vogus data type :

      int, int8, int16, int32, int64
      uint8, uint16, uint32, uint64 (cann't store -ve number)
      float32, float64

	  bool -> 8 bits
	  byte -> alias for uint8 (cann't store -ve number)
  *** rune -> alias for int32 -> contain unicode (a-z & number chara ja ache sb hlo unicode)
	                      -> %c

  5) C er int -> 16 or 32 bits
         char -> 8 bits
     go er int -> pc er bits
*/

import "fmt"
func main(){

	var r = '💪'   // rune = int32
	fmt.Println(r)
	fmt.Printf("%c\n", r)
	fmt.Printf("%T\n", r)

	dev := "Abdullah sh"
	fmt.Printf("%s\n", dev)
	fmt.Printf("%T\n", dev)
}