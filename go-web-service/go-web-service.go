package main

import (
	"fmt"
	db "go-web-service/db"
	controllers "go-web-service/controllers"
	// import (direct)	Used when you need to directly call methods, functions, or types from the package.
	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Go Web Service")
	router := gin.Default()

	db, err := db.ConnectToDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	router.GET("/users", controllers.GetUsers)
	router.POST("/add-users", controllers.AddUsers)
	router.PUT("/edit-users/:userId", controllers.EditUser)
	router.DELETE("/delete-users/:userId", controllers.DeletetUser)
	router.Run("localhost:8080")
}
