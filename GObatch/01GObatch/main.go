package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/chararch/gobatch"
	"github.com/chararch/gobatch/util"
	_ "github.com/go-sql-driver/mysql"
)

func mytask() {
	fmt.Println("Task Executed")
}

type myReader struct{}
type myProcessor struct{}
type myWriter struct{}

func (r *myReader) Read(chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	curr, _ := chunkCtx.StepExecution.StepContext.GetInt("read.num", 1)
	if curr <= 10 {
		chunkCtx.StepExecution.StepContext.Put("read.num", curr+1)
		fmt.Println("stepcount", curr)
		return fmt.Sprintf("value :-%v", curr), nil

	}
	return nil, nil
}

func (r *myProcessor) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	return fmt.Sprintf("Processed:-%v", item), nil
}

func (r *myWriter) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	fmt.Printf("Write:-%v\n", items)
	return nil
}

func main() {
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gobatch?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	gobatch.SetDB(db)

	step1 := gobatch.NewStep("mytask").Handler(mytask).Build()
	step2 := gobatch.NewStep("mystep").Reader(&myReader{}).Processor(&myProcessor{}).Writer(&myWriter{}).ChunkSize(10).Build()
	job := gobatch.NewJob("myjob").Step(step1, step2).Build()
	gobatch.Register(job)
	params, _ := util.JsonString(map[string]interface{}{"rand": time.Now().Nanosecond()})
	gobatch.Start(context.Background(), job.Name(), params)
}
