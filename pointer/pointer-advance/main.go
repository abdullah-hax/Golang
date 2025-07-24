package main

import "fmt"

func golang(numbers [3]int){   // numbers array te amra arr er 3ta ful **copy** kore pass korlam
	fmt.Println(numbers)
}

func super(numbers *[3]int){  // numbers pointerti te amra sudumatro arr er address pass krlam
	fmt.Println(numbers)
}

func main() {

	var arr = [3]int{3, 5, 7}

	golang(arr)  // pass by value

	super(&arr)  // pass by reference

	var arr1 = [3]string{"la", "ilaha", "illalah"}

	focus(&arr1)


}

func focus(kalema *[3]string) {
	fmt.Println(kalema)
}


/*

	=> ekta 3 coti fuler array nilam, arrayti golang() e pass krlam, golang() e numbers name ekta array ache.
	numbers array te 3 coti ful **copy** kore pass kora hoyece, then numbers k print krlam, 3 coti ful print holo.

	=> ebar arraytir address k super() e pass krlam, super() e numbers name ekta pointer array ache.
	numbers pointer array te only array er address pass krlam, then numbers k print krte bollam,
	ekn numbers e array er address(ortat reference) ache, tai array te cole gelo and array k print korlo. (numbers -> array er address ortat reference -> array)

*/