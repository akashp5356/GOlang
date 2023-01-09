package main

import (
	"ABC/conn"
	"ABC/model"

	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chararch/gobatch"
	"github.com/chararch/gobatch/util"
	_ "github.com/go-sql-driver/mysql"
)

type MyReader struct {
}
type MyProcessor struct {
}
type MyWriter struct {
}

func main() {
	fmt.Println("Job started")
	var godb = conn.MyDBGO()
	fmt.Println("Connection Successful!")
	gobatch.SetDB(godb)
	step1 := gobatch.NewStep("Task").Handler(task).Build()
	step2 := gobatch.NewStep("Process").Reader(&MyReader{}).Processor(&MyProcessor{}).Writer(&MyWriter{}).ChunkSize(1).Build()
	fmt.Println("Job in chunk")
	job := gobatch.NewJob("Insert in Test DB with tablename:playerdata").Step(step1, step2).Build()
	gobatch.Register(job)
	params, _ := util.JsonString(map[string]interface{}{"rand": time.Now().Nanosecond()})
	gobatch.Start(context.Background(), job.Name(), params)
}

func task() {
	fmt.Println("Task Function")
}

func (r *MyReader) Read(chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	file, err := os.Open("names.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := csv.NewReader(file).ReadAll()

	//fmt.Println("this is data ", data)
	pdata := []model.Players{}
	_ = pdata

	for _, val := range data {
		pdata = append(pdata, model.Players{Pid: val[0], Name: val[1], Field: val[2]})
	}

	curr, _ := chunkCtx.StepExecution.StepContext.GetInt("read.num", 1)
	if curr <= 1 {
		chunkCtx.StepExecution.StepContext.Put("read.num", curr+1)
		return pdata, nil
	}
	fmt.Println("Read", pdata)
	return nil, nil
}

func (r *MyProcessor) Process(item interface{}, chunkCtx *gobatch.ChunkContext) (interface{}, gobatch.BatchError) {
	newitem := item.([]model.Players)
	return newitem, nil
}

func (r *MyWriter) Write(items []interface{}, chunkCtx *gobatch.ChunkContext) gobatch.BatchError {
	var Idb = conn.MyDBInsert()
	for _, v := range items {
		newv := v.([]model.Players)
		for idx, _ := range newv {
			fmt.Println("this is the val", newv[idx].Name)
			Idb.Exec("insert into playerdata(id,Name,Field) values(?,?,?)", newv[idx].Pid, newv[idx].Name, newv[idx].Field)
		}
	}
	return nil
}
