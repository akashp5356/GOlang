package main

import (
	// "Practice/CSVtoMYSQLbyGObatch/model"
	"context"
	"encoding/csv"
	"fmt"
	"hello/connconfig"
	"hello/model"
	"log"
	"os"
	"time"

	"github.com/chararch/gobatch"
	"github.com/chararch/gobatch/util"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Job started")
	var godb = connconfig.Mygodb()
	fmt.Println("Connection Successful!")
	gobatch.SetDB(godb)
	//step1 := gobatch.NewStep("Task").Handler(task).Build()
	step2 := gobatch.NewStep("Process").Reader(&MyReader{}).Processor(&MyProcessor{}).Writer(&MyWriter{}).ChunkSize(1).Build()
	fmt.Println("Job in chunk")
	job := gobatch.NewJob("playerdata").Step(step2).Build()
	gobatch.Register(job)
	params, _ := util.JsonString(map[string]interface{}{"rand": time.Now().Nanosecond()})
	gobatch.Start(context.Background(), job.Name(), params)
}

type MyReader struct {
}
type MyWriter struct {
}
type MyProcessor struct {
}

func (r *MyReader) Read(chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	var cnt int = 0
	var db = connconfig.MyDB()
	rows, err := db.Query("Select * from playerdata")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	pl := model.Players{}
	var temp []string
	var temp2 [][]string
	for rows.Next() {
		err = rows.Scan(&pl.Pid, &pl.Name, &pl.Field)
		if err != nil {
			log.Fatal(err)
		}
		temp = []string{pl.Pid, pl.Name, pl.Field}
		temp2 = append(temp2, temp)
		cnt++
	}
	fmt.Println("CNT", cnt)
	// fmt.Println("this is the value of temp", temp)
	// fmt.Println("this is the length of temp", len(temp))
	curr, _ := chunkCtx.StepExecution.StepContext.GetInt("read.num", 1)
	if curr < cnt {
		chunkCtx.StepExecution.StepContext.Put("read.num", curr+1)
		return temp2, nil
	}
	return nil, nil
}

func (r *MyProcessor) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	prows := item.([][]string)
	// fmt.Println("prows", prows)
	return prows, nil
}

func (r *MyWriter) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	file, err := os.Create("new.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	defer file.Close()
	for _, v := range items {
		newv := v.([][]string)
		for idx, _ := range newv {
			// fmt.Println("newv", newv[idx])
			err = writer.Write(newv[idx])
			if err != nil {
				log.Fatal(err)
			}
			defer writer.Flush()

		}
	}
	// newtemp := items.([]string)

	return nil
}
