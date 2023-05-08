package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Rno     int    `json:"rno"`
	Name    string `json:"name"`
	Faculty string `json:"faculty"`
}

func myDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	fmt.Println("Inserting JSON data to table")
	var db = myDB()
	fmt.Println("Reading JSON data")
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	data := []student{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Writing data in DB")
	for i := 0; i < len(data); i++ {
		//checking data
		//	fmt.Println("ID:", data[i].Rno, "Name", data[i].Name, "Fac", data[i].Faculty)
		db.Exec("insert into studentdata (rno,name,faculty) values (?,?,?)", data[i].Rno, data[i].Name, data[i].Faculty)
	}
	fmt.Println("Data Inserted")
}
