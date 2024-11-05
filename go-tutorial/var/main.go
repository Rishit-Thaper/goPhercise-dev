package main

import "fmt"

// outerVar := 2000000 use it inside only functions or methods
// var outerVar = 2000000 use it globally
const OuterVar = 200000 // it is a public variable as it starts with capital letter
func main() {
	var username string = "Rishit"
	fmt.Println(username)
	fmt.Printf("Variable type is %T", username)

	// implicit type
	age := 20
	fmt.Println(age)
	fmt.Printf("Variable type is %T", age)

	fmt.Printf("OuterVar is %d", OuterVar)
}
