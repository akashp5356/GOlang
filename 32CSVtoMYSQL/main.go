package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//  go get github.com/go-sql-driver/mysql

type players struct {
	Pid   string
	Name  string
	Field string
}

func myDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	var db = myDB()
	pdata := []players{}
	fmt.Println("Opening CSV file.")
	file, err := os.Open("names.csv")
	if err != nil {
		panic(err)
	}
	df := csv.NewReader(file)
	data, _ := df.ReadAll()
	for _, value := range data {
		pdata = append(pdata, players{Pid: value[0], Name: value[1], Field: value[2]})
	}
	for i := 0; i < len(pdata); i++ {
		id, _ := strconv.Atoi(pdata[i].Pid)
		db.Exec("insert into playerdata(id,Name,Field) values(?,?,?)", id, pdata[i].Name, pdata[i].Field)
	}
	fmt.Println("Inserted")
}
