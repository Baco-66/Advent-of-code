package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//file_path := "test_data.txt"
	//file_path := "data.txt"

	result := 0

	// Prompt the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Advent of code day !")
	fmt.Println("For part one type number 1")
	fmt.Println("For part one type number 2")
	fmt.Print("What is the query?\n")

	// Read the user input
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Remove extra whitespace or newlines

	// Look up and run the chosen function
	switch input {
	case "1":

	case "2":

	default:
		fmt.Println("Invalid choice. Please select a valid function.")
	}

	fmt.Println("Result:", result)
}
