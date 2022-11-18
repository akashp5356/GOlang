package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"cid"`
	CourseName  string  `json:"cname"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}
func main() {
	fmt.Println("API")
	r := mux.NewRouter()
	courses = append(courses, Course{CourseId: "1", CourseName: "C++", CoursePrice: 199, Author: &Author{Name: "James", Website: "c.org"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Java", CoursePrice: 299, Author: &Author{Name: "Mark", Website: "java.dev"}})
	
	//routing
	r.HandleFunc("/",servehome).Methods("GET")
	r.HandleFunc("/getallcourses",getallcourses).Methods("GET")
	r.HandleFunc("/getonecourse/{id}",getonecourse).Methods("GET")
	r.HandleFunc("/createcourse",createcourse).Methods("POST")
	r.HandleFunc("/updatecourse/{id}",updatecourse).Methods("PUT")
	r.HandleFunc("/deletecourse/{id}",deletecourse).Methods("DELETE")

// listen to a port
log.Fatal(http.ListenAndServe(":5000", r))
}

func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome</h1>"))
}

func getallcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get id from request
	//loop through courses and find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("No data found  with given id")
		return
	}
}

func createcourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	//body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Enter some data")
	}
	// {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
// for _, course := range courses {
// 	if course.CourseName== }
	
	//generate random id and append them
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updatecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")
	//get id from request
	params := mux.Vars(r)
	//loop
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deletecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")
	//get id from request
	params := mux.Vars(r)
	//loop
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Deleted Successfully")

}
