package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tinrab/retry"
	"log"
	"time"
)


func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/caiyun")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	retry.ForeverSleep(2*time.Second, func(_ int) (_ error) {
		err = db.Ping()
		if err != nil {
			log.Println(err)
		}
		return
	})

}