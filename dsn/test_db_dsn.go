package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	sqldb, err := sql.Open("postgres", "postgres://postgres:postgres@172.17.0.2:5432?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sqldb.Ping())

	/*ddl := fmt.Sprintf(`create database "%s"`, "jkraj")
	_, err = sqldb.Exec(ddl)
	if err != nil {
		fmt.Println(err)
	}*/

	user_query := fmt.Sprintf(`create user %s with password '%s' createrole createdb`, "koil", "koil@123")
	_, err = sqldb.Exec(user_query)
	if err != nil {
		fmt.Println(err)
	}
}
