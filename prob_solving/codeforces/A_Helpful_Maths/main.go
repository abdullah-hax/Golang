/*

 Problem link : https://codeforces.com/problemset/problem/339/A

declare str[101]
read str

for i = 0; str[i] != '\0'; i = i+2
   for  j = i + 2; str[j] != '\0'; j = j+2
      if(str[j] < str[i])
         str[i] = str[j]


print(str)

*/


package main

import "fmt"

func main() {
	var str string

	fmt.Scan(&str)

	var i, j int

	byt := []byte(str)   // see note
	
	var len = len(byt)

	for i = 0; i < len; i += 2 {
		for j = i + 2; j < len; j += 2 {
			if byt[j] < byt[i] {

				byt[i], byt[j] = byt[j], byt[i]
			}

		}
	}

	fmt.Println(string(byt))
}


/* === Note ===
    // or,  r := []rune(str)    // []rune = int32
	// str e emoji dile byte slice e seta koekti vage vag hoye jabe(because []byte = uint8, jekane emoji er bit aro beshi)
	// -ve number dile problem nai, karon -123 k []byte 4ta character hisebe dhorbe, ekta -ve nmbr hisebe dorbena.
*/

/* 

  => []byte / []rune string k char typer string banie dei (C er mto hoye jai)
  => ei code e len(str) & len(b) same, tbe jdi slice byt teke kno character remove kora hoi tkn kintu 2ta same na

  =>        var str string:               var str [5]string:
            -------------------------            ---------------------------------
            single string                array of 5 strings
            non modifyable           every string modifyable
            modify krte []byte / []rune
            slice e convert krte hbe

  => 
      arr1 := [6]string{"This", "is", "a", "Go", "interview", "questions"}
      arr1[2] = "hello"
      fmt.Println(arr1)

      arr2 := []string{"interview question"}
      arr2[0] = "hello"
      fmt.Println(arr2)

*/



/* === ( In C ) wrong condition ❌ ===

#include <stdio.h>
#include <string.h>

int main() {
    
    char str[101];
    scanf("%s", str);

    int i, j;

    int len = strlen(str);

    for(i = 0; str[i] != '\0'; i += 2){
        for(j = i + 2; str[j] != '\0'; j += 2){
            if(str[j] < str[i]){
               char temp = str[i];   // Note: temp must be character
               str[i] = str[j];
               str[j] = temp;
            }
        }
    }

    printf("%s", str);
    return 0;
}

==========
 jkn only ekta digit (such as 2) input dibo tkn-
       i = 0 er jnno vetore dukbe, j = 2 hbe, 
       condition check kore dkbe str[2] != '\0' hoise, but str[2] = grabage value
       hoito tkn 2 & garbage value swap hbe ,tkn str[0] = garbage value, so str = garbage value
       (programiz er compiler e problem hoinai, codeforces & vs code er compiler e emnta ghotse)

*/

// so condition hisebe " str[j] != '\0' " use na kore " len " use koro ✅
