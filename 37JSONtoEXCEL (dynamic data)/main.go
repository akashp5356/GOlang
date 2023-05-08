package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	generateExcel()
}

func generateExcel() {
	//json data
	jsonData := `[
		{"Full_Name":"John Doe","Age":30,"Dept":"HR","Lang":"Eng/Ger","Email":"john.doe@example.com"},
		{"Full_Name":"Jane ","Age":25,"Dept":"Finance","Lang":"French","Email":"jane@example.com"},
		{"Full_Name":"Bob Smith","Age":35,"Dept":"IB","Lang":"Eng/Ger/Spanish","Email":"bob.smith@example.com"},
		{"Full_Name":"Alen","Age":35,"Dept":"Audit","Lang":"Eng","Email":"alen@example.com"},
		{"Full_Name":"Shakira","Age":25,"Dept":"GM","Lang":"Eng","Email":"shaki@example.com"},
		{"Full_Name":"Bobby","Age":53,"Dept":"Business","Lang":"Italian","Email":"bobby@example.com"},
		{"Full_Name":"James","Age":30,"Email":"james@example.com","Dept":"Sports","Lang":"Columbian"},
		{"Age":25,"Full_Name":"Jamie","Dept":"Staff","Lang":"Russian","Email":"jamie@example.com"},
		{"Dept":"Member","Lang":"Eng","Email":"Blake@example.com","Full_Name":"Blake","Age":28}
	]`
	//heading to set in excel
	IPheading := `Name,Age,Department,Language,Mail ID`
	//key values to get data from json
	IPdata := `Full_Name,Age,Dept,Lang,Email`
	//remove any whitespaces after commas
	IPheading = strings.ReplaceAll(IPheading, ", ", ",")
	IPdata = strings.ReplaceAll(IPdata, ", ", ",")
	//convert string to array of string
	myStrings := strings.Split(IPheading, ",")
	myStringForData := strings.Split(IPdata, ",")

	// Parse the JSON data into a slice of maps
	var data []map[string]interface{}

	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("Failed to parse JSON data:", err)
		return
	}
	//create new excel file
	xlsx := excelize.NewFile()
	//set sheetname
	sheet1Name := "Data"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	// Define variables for column headers
	var headers []string

	// Loop through columns A to Z and add headers to the slice
	for i := 'A'; i <= 'Z'; i++ {
		headers = append(headers, string(i))
	}

	// Define variables for the next set of column headers
	var nextHeaders []string
	nextIndex := 0
	// Loop through columns AA to AZ, BA to BZ, and so on, and add headers to the slice until the end of the array or maximum length
	for i := 'A'; len(headers)+len(nextHeaders) < len(myStrings) && i <= 'Z'; i++ {
		for j := 'A'; len(headers)+len(nextHeaders) < len(myStrings) && j <= 'Z'; j++ {
			if nextIndex < len(myStrings) {
				nextHeaders = append(nextHeaders, fmt.Sprintf("%s%s", string(i), string(j)))
				nextIndex++
			} else {
				break
			}
		}
	}
	// Append the next set of column headers to the original slice
	headers = append(headers, nextHeaders...)

	// Print the headers along with the array values
	for i, val := range myStrings {
		if i < len(headers) {
			// fmt.Printf("%s %s\n", headers[i], val)
			xlsx.SetCellValue(sheet1Name, fmt.Sprintf("%s1", headers[i]), fmt.Sprintf("%s", val))

		} else {
			break
		}
	}

	// Define variables for column headers
	var headerdata []string

	// Loop through columns A to Z and add headerdata to the slice
	for i := 'A'; i <= 'Z'; i++ {
		headerdata = append(headerdata, string(i))
	}

	// Define variables for the next set of column headerdata
	var nextData []string
	nextIndexData := 0

	// Loop through columns AA to AZ, BA to BZ, and so on, and add headerdata to the slice until the end of the array or maximum length
	for i := 'A'; len(headerdata)+len(nextData) < len(myStringForData) && i <= 'Z'; i++ {
		for j := 'A'; len(headerdata)+len(nextData) < len(myStringForData) && j <= 'Z'; j++ {
			if nextIndexData < len(myStringForData) {
				nextData = append(nextData, fmt.Sprintf("%s%s", string(i), string(j)))
				nextIndexData++
			} else {
				break
			}
		}
	}

	// Append the next set of column headerdata to the original slice
	headerdata = append(headerdata, nextData...)

	// Print the headerdata along with the array values
	for idx, dt := range data {
		for i, val := range myStringForData {
			if i < len(headerdata) {
				// fmt.Printf("%s%d %s\n", headerdata[i], i+2, dt[val])
				xlsx.SetCellValue(sheet1Name, fmt.Sprintf("%s%d", headerdata[i], idx+2), dt[val])

			} else {
				break
			}
		}
	}
	//save file(you can give your location)
	err := xlsx.SaveAs(`./Dynamic.xlsx`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Excel file created successfully")
}
