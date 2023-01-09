package conn

import (
	"database/sql"
	"log"
)

func MyDBInsert() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func MyDBGO() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gobatch?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
