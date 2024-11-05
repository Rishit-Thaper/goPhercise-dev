package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	//comma ok syntax || err ok syntax

	input, _ := reader.ReadString('\n')

	// input, err := reader.ReadString('\n') if anything goes wrong the error gets stored in err it's like try catch block in js
	
	fmt.Println("You entered: ", input)
	fmt.Printf("You entered: %s of type: %T", input, input)
}