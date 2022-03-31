package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	//"net/http"
)

/* func StartPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
} */

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// *******************************************************************************************
// *******************************************************************************************
// *******************************************************************************************

func main() {

	//http.HandleFunc("/", StartPage)
	//http.ListenAndServe(":5500", nil)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	/* insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUE ('Alex', 32)")
	if err != nil {
		panic(err)
	}
	defer insert.Close() */

	res, err := db.Query("SELECT `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}

	fmt.Println("Enter to MySQL")
}
