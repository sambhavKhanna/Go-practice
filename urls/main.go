package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=react&pid=32iou903u2"

func main() {
	fmt.Println("The mandem celebrate eid")
	fmt.Println(myurl)
	res, _ := url.Parse(myurl)
	fmt.Println(res.RawQuery)
	qParams := res.Query()
	fmt.Println(qParams["coursename"])
}