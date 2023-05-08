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
	RollNo  string
	Name    string
	Faculty string
}

func myDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	fmt.Println("Reading data from table")
	var db = myDB()
	rows, err := db.Query("select * from studentdata")
	if err != nil {
		log.Fatal(err)
	}
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Colums", cols)
	// fmt.Println("",)
	// jsondata := make([]map[string]interface{}, 0)
	defer rows.Close()
	var temp []string
	var temp2 [][]string
	data := student{}
	for rows.Next() {
		err = rows.Scan(&data.RollNo, &data.Name, &data.Faculty)
		if err != nil {
			log.Fatal(err)
		}
		temp = []string{data.RollNo, data.Name, data.Faculty}
		temp2 = append(temp2, temp)
	}

	// fmt.Println("temp2", temp2)
	d, err := json.Marshal(temp2)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(d))
	ioutil.WriteFile("data.json", d, 100)
	fmt.Println("Completed")
}
