package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type players struct {
	Pid   string
	Name  string
	Field string
}

func myDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
func main() {
	var db = myDB()
	rows, err := db.Query("Select * from playerdata")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pl := players{}
	var temp []string
	var temp2 [][]string
	for rows.Next() {
		err = rows.Scan(&pl.Pid, &pl.Name, &pl.Field)
		if err != nil {
			log.Fatal(err)

		}
		temp = []string{pl.Pid, pl.Name, pl.Field}
		temp2 = append(temp2, temp)
	}
	writer := csv.NewWriter(file)
	err = writer.WriteAll(temp2)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Flush()
	fmt.Println("Completed!!")
}
