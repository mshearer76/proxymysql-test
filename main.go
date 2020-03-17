package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mshearer76/mysql"
	"log"
)

func main() {
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:16033)/mydb")
		if err != nil {
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}
		defer db.Close()

		rows, err := db.Query("select @@hostname")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			var hostname string
			for rows.Next() {
				err := rows.Scan(&hostname)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("hostname = %s\n", hostname)
			}
		}

		defer rows.Close()
}
