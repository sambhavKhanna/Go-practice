package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices")

	var fruits = []string{"q", "3", "y"}
	fmt.Printf("Type of fruits %T", fruits)
	fruits = append(fruits, "34")

	highscore := make([]int, 4)
	highscore[3] = 123
}