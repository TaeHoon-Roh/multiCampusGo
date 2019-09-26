package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:uxfac@tcp(127.0.0.1:3306)/multicampus")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Db Connect Error!!")
		os.Exit(1)
	}
	q := "show databases"
	result, exerror := db.Query(q)
	if exerror != nil {
		fmt.Println("ExError")
		os.Exit(1)
	} else {
		for result.Next() {
			var temp string
			err := result.Scan(&temp)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(temp)
			}

		}
	}

	defer func() {
		db.Close()
		result.Close()
	}()
}
