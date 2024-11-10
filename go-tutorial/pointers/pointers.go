package main

import "fmt"

func main() {
	//memory alloc/dealloc happens automatically in Go.
	//Go has garbage collection, so we don't have to worry about deallocating memory. GOGC=off to disable garbage collection.
	//new() function is used to create a pointer to a variable, allocate memory but no INIT value, you'll get the memory address, we get zeroed storage.
	//new() is used to create pointers to types, and it returns a pointer to the zeroed value of type T (not T).

	//make() function is used to create a pointer to a variable, allocate memory and INIT value, you'll get the memory address, we get non-zeroed storage.
	//make() is used to create slices, maps, and channels, and it returns an initialized (not zeroed) value of type T (not *T).

	//Pointers are used to store the memory address of another variable.

	fmt.Println("Pointers in Go")

	var ptr *int // nil pointer
	// var ptr *int = new(int)

	number := 123
	ptr = &number //referencing the memory address of number variable
	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Value of actual ptr:", *ptr) //value inside the pointer
	
	*ptr = 1232 * *ptr //actual value getting updated in the memory address for number variable
	fmt.Println("Value of actual ptr:", number) //value inside the pointer
}
