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

// Model for course
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

// fake DB
var courses []Course 

// middleware, helper
func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

type Author struct {
	Name string `json:"name"`
	Website string `json:"website"`
}

func main() {
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId:"2", CourseName:"Reactjs", CoursePrice: 123, Author: &Author{Name: "Gilgamesh", Website: "http://"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Wagwan</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL COURSES")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop through the courses and find matching id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("{ message: Send data }")
		return
	}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		obj := struct {
			Message string `json:"banaras"`
		} {
			Message: "Send data",
		}
		// package this as JSON data
		//finalJson, _ := json.MarshalIndent(obj, "", "  ")
		json.NewEncoder(w).Encode(obj)
		return
	}

	// generate a UID
	// append new course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range(courses) {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1 :]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// Not found
	json.NewEncoder(w).Encode("Not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id, remove
	for index, course := range(courses) {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index +1 :]...)
			return
		}
	}
}