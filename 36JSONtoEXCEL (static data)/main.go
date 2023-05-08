package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//map string interface
type M map[string]interface{}

var data = []M{
	M{"Full_Name": "John Doe", "Age": 30, "Dept": "HR", "Lang": "Eng/Ger", "Email": "john.doe@example.com"},
	M{"Full_Name": "Jane ", "Age": 25, "Dept": "Finance", "Lang": "French", "Email": "jane@example.com"},
	M{"Full_Name": "Bob Smith", "Age": 35, "Dept": "IB", "Lang": "Eng/Ger/Spanish", "Email": "bob.smith@example.com"},
	M{"Full_Name": "Alen", "Age": 35, "Dept": "Audit", "Lang": "Eng", "Email": "alen@example.com"},
	M{"Full_Name": "Shakira", "Age": 25, "Dept": "GM", "Lang": "Eng", "Email": "shaki@example.com"},
	M{"Full_Name": "Bobby", "Age": 53, "Dept": "Business", "Lang": "Italian", "Email": "bobby@example.com"},
	M{"Full_Name": "James", "Age": 30, "Email": "james@example.com", "Dept": "Sports", "Lang": "Columbian"},
	M{"Age": 25, "Full_Name": "Jamie", "Dept": "Staff", "Lang": "Russian", "Email": "jamie@example.com"},
	M{"Dept": "Member", "Lang": "Eng", "Email": "Blake@example.com", "Full_Name": "Blake", "Age": 28},
}

func main() {
	generateExcel()
}

func generateExcel() {
	//createnew file
	xlsx := excelize.NewFile()
	//sheet name
	sheet1Name := "Static Data"
	//style bold for heading
	style, _ := xlsx.NewStyle(`{
		"font": {
			"bold": true,
			"size": 12
		},
		"fill": {
			"type": "pattern",
			"pattern": 1
		}
	}`)

	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)
	//set the headings for columns
	xlsx.SetCellValue(sheet1Name, "A1", "Name")
	xlsx.SetCellValue(sheet1Name, "B1", "Age")
	xlsx.SetCellValue(sheet1Name, "C1", "Language")
	xlsx.SetCellValue(sheet1Name, "D1", "Department")
	xlsx.SetCellValue(sheet1Name, "E1", "Mail")
	//set style
	xlsx.SetCellStyle(sheet1Name, "A1", "E1", style)
	//set filter property
	err := xlsx.AutoFilter(sheet1Name, "A1", "E1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	//range over data and set appropriate fields
	for i, each := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["Full_Name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["Age"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each["Lang"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), each["Dept"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), each["Email"])
	}
	//save the file
	err = xlsx.SaveAs(`./Static.xlsx`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Excel file created successfully")
}
