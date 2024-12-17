package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Define a struct to hold the data for each entry
type Entry struct {
	ButtonA struct{ X, Y int }
	ButtonB struct{ X, Y int }
	Prize   struct{ X, Y int }
}

type ErrParseLine struct {
	Line string
}

func (e ErrParseLine) Error() string {
	return fmt.Sprintf("Failed to find X and Y in line: %s", e.Line)
}

// Function to parse a line with X and Y values
func parseLine(line string, regex *regexp.Regexp) (x, y int, err error) {
	matches := regex.FindStringSubmatch(line)
	if len(matches) == 3 {
		fmt.Sscanf(matches[1], "%d", &x)
		fmt.Sscanf(matches[2], "%d", &y)
		err = nil
	} else {
		x, y = -1, -1
		err = ErrParseLine{line}
	}
	return
}

// Function to process 3 lines into an Entry
func process_machine(line1, line2, line3 string, buttonRegex, prizeRegex *regexp.Regexp) (Entry, error) {
	var entry Entry
	var err error

	// Parse Button A
	entry.ButtonA.X, entry.ButtonA.Y, err = parseLine(line1, buttonRegex)
	if err != nil {
		return entry, err
	}

	// Parse Button B
	entry.ButtonB.X, entry.ButtonB.Y, err = parseLine(line2, buttonRegex)
	if err != nil {
		return entry, err
	}

	// Parse Prize
	entry.Prize.X, entry.Prize.Y, err = parseLine(line3, prizeRegex)
	if err != nil {
		return entry, err
	}

	return entry, nil
}

func parse_data(file_path string) ([]Entry, error) {
	// Open the file
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err // Return error if file cannot be opened
	}
	defer file.Close() // Ensure the file is closed when the function ends

	// Regular expressions for parsing
	buttonRegex := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`X=(\d+), Y=(\d+)`)

	// Slice to store all entries
	var entries []Entry
	lines := []string{}

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
			if len(lines) == 3 {
				// Process 3 lines into an entry
				entry, err := process_machine(lines[0], lines[1], lines[2], buttonRegex, prizeRegex)
				if err != nil {
					return nil, err // Return error if file cannot be opened
				}
				entries = append(entries, entry)
				lines = lines[:0] // Clear the lines slice
			}
		}
	}

	return entries, nil
}

func (e Entry) find_prize_combination(price_x, price_y int) (int, int, bool) {

	// Calculate the denominator
	denominator := e.ButtonA.Y*e.ButtonB.X - e.ButtonA.X*e.ButtonB.Y
	if denominator == 0 {
		return 0, 0, false
	}

	// Calculate n and check if it's an integer
	numeratorN := (e.Prize.Y*e.ButtonB.X - e.Prize.X*e.ButtonB.Y)
	if numeratorN%denominator != 0 {
		// If the numerator is not divisible by the denominator, n is not an integer
		return 0, 0, false
	}
	n := numeratorN / denominator

	// Calculate m and check if it's an integer
	numeratorM := (e.Prize.X - n*e.ButtonA.X)
	if numeratorM%e.ButtonB.X != 0 {
		// If the numerator is not divisible by the denominator, m is not an integer
		return 0, 0, false
	}
	m := numeratorM / e.ButtonB.X

	return n * price_x, m * price_y, true
}

func main() {

	// Open the file and get a scanner
	//file_path := "test_data.txt"
	file_path := "data.txt"
	entries, err := parse_data(file_path)

	// Check for errors during scanning
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := 0

	// Prompt the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Advent of code day 13!")
	fmt.Println("For part one type number 1")
	fmt.Println("For part one type number 2")
	fmt.Print("What is the query?\n")

	// Read the user input
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Remove extra whitespace or newlines

	// Look up and run the chosen function
	switch input {
	case "1":
		for _, entry := range entries {
			x, y, err := entry.find_prize_combination(3, 1)
			if err {
				result += x + y
			}
		}
	case "2":
		for _, entry := range entries {
			entry.Prize.X += 10000000000000
			entry.Prize.Y += 10000000000000
			x, y, err := entry.find_prize_combination(3, 1)
			if err {
				result += x + y
			}
		}

	default:
		fmt.Println("Invalid choice. Please select a valid function.")
	}

	fmt.Println("Result:", result)
}
