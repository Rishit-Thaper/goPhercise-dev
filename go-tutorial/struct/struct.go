package main

import "fmt"

//capitalized fields are exported
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	fmt.Println("Structs in Go")
	//no inheritance in go
	//no super or parent
	
	userDetail := []User{{"Rishit", "rishit@dev.in", 21}, {"Rishu", "rishu@dev.in", 21}}
	for index := range userDetail {
		fmt.Println("User ", index+1, " is ", userDetail[index])
	}
	rishit := User{"Rishit", "rishit@dev.in", 21}
	fmt.Println(rishit)
	fmt.Printf("Name is %v, Age is %v and Email is %v.\n", rishit.Name, rishit.Age, rishit.Email)
}
