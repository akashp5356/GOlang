package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type M map[string]interface{}

func main() {
	generateExcel()
}

func generateExcel() {
	xlsx, err := excelize.OpenFile("./Static.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	fmt.Println("sdsa", reflect.TypeOf(xlsx))
	sheet1Name := "Static Data"

	rows := make([]M, 0)
	for i := 2; i < 5; i++ {
		row := M{
			"Name":       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
			"Age":        xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
			"Language":   xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
			"Department": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i)),

			"Mail": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i)),
		}
		rows = append(rows, row)
	}

	fmt.Printf("%v \n", rows)
}
