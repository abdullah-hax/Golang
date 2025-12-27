package main

import "fmt"

func main() {
	var recent, perYear, total int
	fmt.Scan(&recent, &perYear, &total)

	fmt.Println((total - recent)/perYear)
}