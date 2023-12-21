package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// main is the entry point of the program.
func main() {
	generateExcel()
}

// generateExcel reads data from an Excel file and maps it to a slice of maps.
func generateExcel() {
	xlsx, err := excelize.OpenFile("./Test.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	sheet1Name := "Static Data"

	rows, err := xlsx.GetRows(sheet1Name)
	if err != nil {
		fmt.Println(err)
		return
	}

	headerRow := rows[0]

	// Read the rows and map them to the headers
	data := make([]map[string]interface{}, 0)
	for i := 2; i < len(rows); i++ {
		row := make(map[string]interface{})
		for j, header := range headerRow {
			if header == "" {
				continue
			}
			value, _ := xlsx.GetCellValue(sheet1Name, fmt.Sprintf("%s%d", string('A'+j), i))
			row[header] = value
		}
		data = append(data, row)
	}

	fmt.Printf("%v \n", data)
}
