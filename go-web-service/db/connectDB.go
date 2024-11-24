package db

import (
	"database/sql"
	"fmt"
		//The _ is used as a "blank identifier." It means you are importing the package only for its side effects and not directly using any of its functions or types in your code.
	//import _	Used to enable side effects (e.g., registering a database driver) without referencing it.
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", GetDBConfig().Username, GetDBConfig().Password, GetDBConfig().DBUrl, GetDBConfig().DBName)

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
