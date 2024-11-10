package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversions")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating: ")
	input, err := reader.ReadString('\n')
	fmt.Println("You entered this number: ", input) //here it was printing and reading in string
	if err != nil {
		fmt.Println("Error Occured: ", err)
	}
	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64) //for converting string to float
	if error != nil {
		fmt.Println("Error Occured: ", error)
		panic(error) //it returns the error and stops the program
	} else {
		fmt.Println("You entered: ", numRating)
		fmt.Println("I increased the rating by 56: ", numRating+56)
		fmt.Printf("You entered: %f of type: %T", numRating, numRating)
	}
}
