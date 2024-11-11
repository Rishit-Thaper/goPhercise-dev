package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Guavava", "Banana"} // Slice declaration
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	length := len(fruitList)
	fmt.Println("length is: ", length)
	// Iterating over elements of a slice
	for index, fruit := range fruitList {
		fmt.Println("Fruit is: ", fruitList[index])
		fmt.Println("Fruit is: ", fruit)
		fmt.Println("index is: ", index)
	}

	//appending the slice
	fruitList = append(fruitList, "Mango")
	fmt.Println("Fruitlist: ", fruitList)

	//slice the slice
	fruitList = fruitList[1:] // Removing the first element
	fmt.Println("Fruitlist: ", fruitList)

	fruitList = fruitList[1:3]
	fmt.Println("Fruitlist: ", fruitList)

	fmt.Println("length is: ", length)

	// Creating a slice using make
	highScores := make([]int, 5)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	highScores[4] = 777
	highScores = append(highScores, 555, 888, 999) // The append function is used to add elements to a slice, and it automatically resizes the slice if necessary.
	// Directly assigning a value to an index (e.g., highScores[6] = 777) will cause an "index out of range" error if the index is beyond the current length of the slice.
	fmt.Println("High Scores: ", highScores)

	sort.Ints(highScores)
	fmt.Println("High Scores: ", highScores)
	last := highScores[len(highScores)-1]
	fmt.Println("last: ", last)

	//remove an element from a slice
	index := 2
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println("High Scores: ", highScores)
}
