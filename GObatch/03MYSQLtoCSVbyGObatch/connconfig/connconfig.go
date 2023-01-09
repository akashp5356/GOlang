package connconfig

import (
	"database/sql"
	"log"
)

func MyDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Mygodb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/gobatch?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
