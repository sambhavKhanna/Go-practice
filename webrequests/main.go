package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://localhost:3000"

func main() {
	fmt.Println("Welcome to webrequests")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response %T\n", res)
	defer res.Body.Close()
	
	databytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}