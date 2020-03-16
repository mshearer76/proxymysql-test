package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	for {
		db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:16033)/mydb")
		if err != nil {
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}

		rows, err := db.Query("select @@max_allowed_packet")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			var max_allowed_packet int
			for rows.Next() {
				err := rows.Scan(&max_allowed_packet)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("max_allowed_packet = %d\n", max_allowed_packet)
			}
		}
		defer rows.Close()
		//time.Sleep(time.Second)
		db.Close()
	}
}
