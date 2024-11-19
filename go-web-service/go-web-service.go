package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	fmt.Println("newUser", newUser)
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	} else {
		users = append(users, newUser)
		c.JSON(http.StatusCreated, users)
	}
}

func editUser(c *gin.Context) {
	id := c.Param("userId")
	fmt.Println("id", id)
	var updatedUser User
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.BindJSON(&updatedUser); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	} else {
		for i, user := range users {
			if userId == user.ID {
				fmt.Println("user to be updated", user)
				users[i] = updatedUser
				c.JSON(http.StatusCreated, users)
				return
			}
		}
	}
}

func deletetUser(c *gin.Context) {
	id := c.Param("userId")
	fmt.Println("id", id)
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	} else {
		for i, user := range users {
			if userId == user.ID {
				fmt.Println("user to be deleted", user)
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"Message": "Deleted Successfully", "users": users})
				return
			}
		}
	}
}

func main() {
	fmt.Println("Go Web Service")
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/add-users", addUsers)
	router.PUT("/edit-users/:userId", editUser)
	router.DELETE("/delete-users/:userId", deletetUser)
	router.Run("localhost:8080")
}
