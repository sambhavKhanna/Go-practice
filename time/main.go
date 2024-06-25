package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Study of Golang")

	present := time.Now()
	fmt.Println(present.Format("01-02-2006 15:04:05 Monday"))
}