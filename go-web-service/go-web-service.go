package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	// import (direct)	Used when you need to directly call methods, functions, or types from the package.
	"github.com/gin-gonic/gin"

	//The _ is used as a "blank identifier." It means you are importing the package only for its side effects and not directly using any of its functions or types in your code.
	//import _	Used to enable side effects (e.g., registering a database driver) without referencing it.
	_ "github.com/go-sql-driver/mysql"

	// import (direct)	Used when you need to directly call methods, functions, or types from the package.
	"github.com/joho/godotenv"
)

// @TODO: Remove this upto line 61
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

// @TODO: Refactor the code by dividing it into multiple files and folders.
// Each file will focus on a specific functionality or module,
// improving abstraction, readability, and maintainability.
// that code will be in a seperate repo.
type DBConfig struct {
	Username string
	Password string
	DBUrl    string
	DBName   string
}

// access the db credentials, like we do in node/express code
func getDBConfig() DBConfig {
	return DBConfig{
		Username: getEnvVariables("DB_USERNAME"),
		Password: getEnvVariables("DB_PASSWORD"),
		DBUrl:    getEnvVariables("DB_URL"),
		DBName:   getEnvVariables("DB_NAME"),
	}
}

// access env's in the go code
func getEnvVariables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	value := os.Getenv(key)
	return value
}

// connecting to db function
func connectToDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("120err", err)
		panic(err)
	}
	err = db.Ping()
	fmt.Println("124err", err)

	if err != nil {
		fmt.Println("127err", err)
		panic(err)
	}
	fmt.Print("Pong\n")
	return db, nil
}

// @TODO: The following function will be implemented with proper DB integration
func getUsers(c *gin.Context) {
	fmt.Println(users)
	c.IndentedJSON(http.StatusOK, users)
}

// @TODO: The following function will be implemented with proper DB integration
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

// @TODO: The following function will be implemented with proper DB integration
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

// @TODO: The following function will be implemented with proper DB integration
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

	dbConfig := getDBConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.DBUrl, dbConfig.DBName)

	db, err := connectToDB(dsn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	router.GET("/users", getUsers)
	router.POST("/add-users", addUsers)
	router.PUT("/edit-users/:userId", editUser)
	router.DELETE("/delete-users/:userId", deletetUser)
	router.Run("localhost:8080")
}
