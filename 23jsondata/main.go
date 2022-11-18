package main

import (
	"encoding/json"
	"fmt"
)



func main() {
	fmt.Println("Json data")
	// Encjson()
	Dejson()
}

func Encjson() {
	ourcourse := []course{
		// {"ASD", 45, "123", "cpp"},
		// {"BNM", 35, "345", "c"},
		// {"CVB", 25, "asd45", "nil"},
		{"A", 45, "123", []string{"c", "cpp"}},
		{"B", 35, "345", []string{"java"}},
		{"C", 25, "asd45", nil},
	}
	fjson, err := json.MarshalIndent(ourcourse,"","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", fjson)
}
type course struct {
	//use capital names to make it public
	Name     string `json:"coursename"`
	Age      int `json:"age"`
	Password string `json:"-"`
	Skills  []string `json:"abilities,omitempty"`
}

func Dejson(){
	data:=[]byte(`
	{
		"coursename": "C",
		"age": 25
}
	`)
	var valid course
	check:=json.Valid(data)
	if check{
		fmt.Println("Json is valid")
		json.Unmarshal(data,&valid)	//convert byte data into the original data structure.
		fmt.Printf("%#v\n",valid)
	} else {
		fmt.Println("Json not valid")
	}

	var onlinedata map[string] interface{}
	json.Unmarshal(data,&onlinedata)
	fmt.Printf("%#v\n",onlinedata)

	for key,value := range onlinedata{
		fmt.Printf("Key is: %v ,Value is: %v and type is %T\n",key,value,value)
	}

}