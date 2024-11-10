package main

import "fmt"

func main() {
	var numArray = [6]int{1, 2, 3, 4, 8, 6} //compulsion to define the size of the array
	fmt.Println(numArray)

	var rangers [4]string

	rangers[0] = "SPD"
	rangers[1] = "NinjaStorm"
	rangers[3] = "DinoThunder"

	fmt.Println(rangers) //[SPD NinjaStorm  DinoThunder] the double emptyspace is the 2nd element (0 based indexing)

	//arrays are not much used in Go, slices are used more often coz they are more powerful and flexible
}
