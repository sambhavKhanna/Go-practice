package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	val := 6
	ptr := &val
	fmt.Println("Value of ptr", ptr)

	*ptr *= 5
	fmt.Println("Value of val", val)
}