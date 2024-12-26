package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func parseFile(file_path string) ([][]rune, error) {
	// Open the file
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Variables to store the matrix and the sequence
	var matrix [][]rune

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Add each line as a row in the matrix
		matrix = append(matrix, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

// Locate a rune in the matrix
func findInMatrix(matrix [][]rune, target rune) (int, int, bool) {
	for i, row := range matrix {
		for j, cell := range row {
			if cell == target {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

// Print the matrix
func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func hasDuplicates(points []*Node_v2) bool {
	pointMap := make(map[Point]bool)

	for _, point := range points {
		if pointMap[point.Point] {
			fmt.Println("Repeated:", point)
			return true // Duplicate found
		}
		pointMap[point.Point] = true
	}
	return false // No duplicates
}

// printMapWithPath overlays the path onto the map and prints it
func printMapWithPath(grid [][]rune, path []*Node_v2) {
	// Create a copy of the original map to avoid modifying it
	mapCopy := make([][]rune, len(grid))
	for i := range grid {
		mapCopy[i] = make([]rune, len(grid[i]))
		copy(mapCopy[i], grid[i])
	}

	// Mark the path on the map
	for _, p := range path {
		// Ensure the point is within bounds
		if p.Point.X >= 0 && p.Point.X < len(mapCopy) && p.Point.Y >= 0 && p.Point.Y < len(mapCopy[0]) {
			mapCopy[p.Point.X][p.Point.Y] = p.Type
		}
	}

	// Print the modified map
	for _, row := range mapCopy {
		fmt.Println(string(row))
	}
}

func main() {
	//file_path := "test_data.txt"
	file_path := "data.txt"
	//file_path := "test_data_mini.txt"

	result := 0

	// Parse the file
	matrix, err := parseFile(file_path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Prompt the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Advent of code day !")
	fmt.Println("For part one type number 1")
	fmt.Println("For part one type number 2")
	fmt.Print("What is the query?\n")

	// Read the user input
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Remove extra whitespace or newlines

	var found bool

	// Look up and run the chosen function
	switch input {
	case "1":
		//var path []Point
		_, result, found = a_star(matrix, weight)
		if !found {
			fmt.Println("No path found.")
			return
		}
	case "2":
		var path []*Node_v2
		path, found = best_spots(matrix, weight_controll_reverse)
		if found {
			printMapWithPath(matrix, path)
			result = len(path)
		} else {
			fmt.Println("No spots found.")
			return
		}
		//fmt.Println(hasDuplicates(path))
	default:
		fmt.Println("Invalid choice. Please select a valid function.")
	}

	fmt.Println("Result:", result)
}
