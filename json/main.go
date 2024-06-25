package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int `json:"price"`
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("WAGWAN JSON")
	// encodeJson()
	DecodeJson()
}

func encodeJson() {
	courses := []course{
		{
			"REACT", 123, "YT", "123", []string{"Web dev", "JS"},
		},
		{
			"MERN", 123, "YT", "345", []string{"Web dev", "Google"},
		},
	}
	// package this as JSON data
	finalJson, err := json.MarshalIndent(courses, "", "  ")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}

func DecodeJson() {
	bhaiKaJson := []byte(`
  {
    "coursename": "REACT",
    "price": 123,
    "website": "YT",
    "tags": [
      "Web dev",
      "JS"
    ]
  }	
	`)
	var apnaCourse course

	isValid := json.Valid(bhaiKaJson)

	if isValid {
		fmt.Println("Valid")
		json.Unmarshal(bhaiKaJson, &apnaCourse)
		fmt.Println(apnaCourse, apnaCourse.Name)
	} else {
		fmt.Println("Valid nahi hai")
	}
}