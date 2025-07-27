package main

import (
	"fmt"
)

func character() {
	str := "abacas"
	runes := []rune(str)

	fmt.Println(str)              // characters | Output: abacas
	fmt.Println(str[0])           // ascii code of character | Output: 97
	fmt.Println(runes)            // unicodes of characters | Output: [97 98 97 99 97 115]
	fmt.Println(runes[0])         // unicode of character | Output: 97
	fmt.Println(string(runes[0])) // character | Output: a
	fmt.Println(string(runes))    // characters | Output: abacas
}

func length() {
	str := "abacas"      // str takes 6 bytes (ASCII)   [a ke 1 byte er modde rakhe, b ke 1 byte er modde rakhe, ...]
	runes := []rune(str) // runes has 6 runes, typically 4 * 6 = 24 bytes (int32)  [a ke 4 byte er modde rakhe, b ke 4 byte er modde rakhe, ...]

	fmt.Println(len(str))   // 6
	fmt.Println(len(runes)) // 6

	/* 	string একটা character কে একটা  byte এর মধ্যে রাখে,  তাই len(str) দিলে 1 byte  এর জন্য length দেখাবে 1,

	   	rune সেই character টিকে 4 byte এর মধ্যে রাখে।
	   	তাই len(runes) দিলে 4 byte  এর জন্য length দেখাবে 1

	   	🙂 1 byte এ রাখা যায়না, 4 byte area তে রাখতে হয়। তাই len(str) = 4 & len(runes) = 1

	   	summary :
	   	🙂 দাও বা Alphabet দাও rune সেটাকে 4 byte এ রাখে  so always length পাওয়া যায় 1
	   	but string Alphabet কে 1 byte এ রাখলেও 🙂 কে 1 byte এ রাখতে পারেনা,  4 byte লাগে,
	   	তাই 2 বার length 2 রকম হয়।    */

	str2 := "😊"
	runes2 := []rune(str2)

	fmt.Println(len(str2))   // 4 (takes 4 bytes!)
	fmt.Println(len(runes2)) // 1 (takes 1 rune)

	fmt.Println(str2[0], str2[1]) // 240 159 (just part of the emoji, not the full char!) [string একটা একটা পার্ট প্রিন্ট করলে ascii code দেখাবে]
	fmt.Println(str2)             // 😊 [পুরো string প্রিন্ট করলে character দেখাবে]

	fmt.Println(runes2[0])
	fmt.Println(runes2)
	fmt.Println(string(runes2))
}

func modify(){
	str := "hello"
	runes := []rune(str)

	// str[0] = 'k'  ❌  string cann't be modified, you should use rune/bute to modify any index
	runes[0] = 'k'
	fmt.Printf("%c\n", str[0])
	fmt.Printf("%c\n", runes[0])
	fmt.Println(string(runes))
}

func main() {

	fmt.Println("First part : ")
	character()
	
	fmt.Println("Second part : ")
	length()

	fmt.Println("Third part : ")
	modify()

}


/* Disadvantage of string 
     -> str modify kora jaina
	 -> character length 1 byte er besi hle real string length paoa jaina
*/