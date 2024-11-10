package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Go")

	presentTime := time.Now()
	fmt.Println("Present Time: ", presentTime)

	fmt.Println("Present time with different format", presentTime.Format("01-02-2006 Monday")) // use this date to format the string for standard

	createdDate := time.Date(2024, time.July, 9, 03, 0, 0, 0, time.UTC)
	fmt.Println("Created Time: ", createdDate)

}
