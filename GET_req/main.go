package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome")
	//performGetRequest()
	PerformPostReq()
}

func performGetRequest() {
	myurl := "http://localhost:3000"

	resp, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status", resp.StatusCode)
	fmt.Println("Content length", resp.ContentLength)
	
	var respStr strings.Builder
	content, _ := io.ReadAll(resp.Body)
	byteCount, _ := respStr.Write(content)
	fmt.Println("Byte count", byteCount)
	fmt.Println(respStr.String())
}

func PerformPostReq() {
	const myUrl = "http://localhost:3000"

	reqBody := strings.NewReader(`{
		"course": "GO Lang",
		"price": "100"
	}`)

	resp, _ := http.Post(myUrl, "application/json", reqBody)

	defer resp.Body.Close()
	var respStr strings.Builder
	content, _ := io.ReadAll(resp.Body)
	respStr.Write(content)
	fmt.Println(respStr.String())
}

func PerformFormRequest() {
	const myUrl = "http://localhost:3000"
}