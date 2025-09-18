package main

import(
	"fmt"
	"strings"
)

func splitIntoSubstrings(msg string) {
	fmt.Println("Splited into substring : ", strings.Split(msg, " "))  // [Hello World!], [] means it's a slice of strings(each elemet is a separate string)
}

func replacingSubstring(msg string){
	fmt.Println("After replacing : ",strings.Replace(msg, "Wor", "Go", 1))
}

func trimingSpace(){
	str := "   Hello World!  "
	fmt.Println("After Trimming space : ", strings.TrimSpace(str))
}

func anyChar(msg string){
	fmt.Printf("%c %c %c %c", msg[0], msg[1], msg[5], msg[11])
}


func main(){
	msg := "Hello World!"
	splitIntoSubstrings(msg)
	replacingSubstring(msg)
	trimingSpace()
	anyChar(msg)
}