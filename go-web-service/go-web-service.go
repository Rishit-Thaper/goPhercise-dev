package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

var users = []User{

	{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
	},
	{
		ID:    2,
		Name:  "Bob",
		Email: "bob@test.com",
		Age:   30,
	},
	{
		ID:    3,
		Name:  "Charlie",
		Email: "charlie@random.org",
		Age:   22,
	},
	{
		ID:    4,
		Name:  "Diana",
		Email: "diana@example.com",
		Age:   28,
	},
	{
		ID:    5,
		Name:  "Ethan",
		Email: "ethan@test.com",
		Age:   35,
	},
}

func getUsers(c *gin.Context) {
	fmt.Println(users)
	c.IndentedJSON(http.StatusOK, users)
}

func addUsers(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	users = append(users, newUser)

	fmt.Println(users)

	c.JSON(http.StatusCreated, users)
}

func main() {
	fmt.Println("Go Web Service")
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/add-users", addUsers)
	router.Run("localhost:8080")
}
