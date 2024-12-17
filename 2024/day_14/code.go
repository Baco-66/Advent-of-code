package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse_data(file_path string) ([][]string, error) {
	// Open the file
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err // Return error if file cannot be opened
	}
	defer file.Close() // Ensure the file is closed when the function ends

	// Initialize a matrix
	var matrix [][]string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read a line
		line := scanner.Text()

		// Split the line into elements (adjust delimiter if needed)
		elements := strings.Split(line, "") // Splits into individual characters

		// Append the slice to the matrix
		matrix = append(matrix, elements)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err // Return error if reading fails
	}

	return matrix, nil
}

func find_plots(x, y int, matrix [][]string, visited map[[2]int]bool) (int, int) {
	// If the cell is already visited, skip it
	if visited[[2]int{x, y}] {
		return 0, 0
	}

	// Mark the current cell as visited
	visited[[2]int{x, y}] = true

	// Initial area and perimeter
	var area, perimeter int = 1, 0
	type_of_plot := matrix[x][y]

	// Define the directions for neighbors (up, left, down, right)
	directions := [][2]int{
		{-1, 0}, // up
		{0, -1}, // left
		{1, 0},  // down
		{0, 1},  // right
	}

	// Check all 4 neighbors
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]

		// Check if the new position is within bounds
		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) {
			if matrix[newX][newY] == type_of_plot {
				// Recursively call find_plots for the neighbor
				res_area, res_perimeter := find_plots(newX, newY, matrix, visited)
				area += res_area
				perimeter += res_perimeter
			} else {
				// Increase the perimeter for non-matching plot
				perimeter++
			}
		} else {
			// Increase perimeter for out-of-bounds
			perimeter++
		}
	}

	return area, perimeter
}

// Helper function to check and count angles
func checkAngle(matrix [][]string, x, y, dx, dy int, type_of_plot string) int {

	angles := 0
	newX, newY := x+dx, y+dy

	if newX >= 0 && newX < len(matrix) {
		if newY >= 0 && newY < len(matrix[0]) { // Case 1: Both newX and newY are valid
			if matrix[newX][y] == type_of_plot && matrix[x][newY] == type_of_plot && matrix[newX][newY] != type_of_plot {
				angles++
			} else if matrix[newX][y] != type_of_plot && matrix[x][newY] != type_of_plot {
				angles++
			}
		} else { // Case 2: Only newX is valid (newY is out of bounds)
			if matrix[newX][y] != type_of_plot {
				angles++
			}
		}
	} else if newY >= 0 && newY < len(matrix[0]) { // Case 3: Only newY is valid (newX is out of bounds)
		if matrix[x][newY] != type_of_plot {
			angles++
		}
	} else { // Case 4: Both newX and newY are out of bounds
		angles++
	}

	return angles
}

func find_plots2(x, y int, matrix [][]string, visited map[[2]int]bool) (int, int) {
	// If the cell is already visited, skip it
	if visited[[2]int{x, y}] {
		return 0, 0
	}

	// Mark the current cell as visited
	visited[[2]int{x, y}] = true

	// Initial area and perimeter
	var area, angles int = 1, 0
	type_of_plot := matrix[x][y]

	// Define the directions for neighbors (up, left, down, right)
	directions := [][2]int{
		{-1, 0}, // up
		{0, -1}, // left
		{1, 0},  // down
		{0, 1},  // right
	}

	angles += checkAngle(matrix, x, y, -1, -1, type_of_plot) // Top-left
	angles += checkAngle(matrix, x, y, 1, -1, type_of_plot)  // Bottom-left
	angles += checkAngle(matrix, x, y, 1, 1, type_of_plot)   // Bottom-right
	angles += checkAngle(matrix, x, y, -1, 1, type_of_plot)  // Top-right

	// Check all 4 neighbors
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]

		// Check if the new position is within bounds
		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) && matrix[newX][newY] == type_of_plot {
			// Recursively call find_plots for the neighbor
			res_area, res_angles := find_plots2(newX, newY, matrix, visited)
			area += res_area
			angles += res_angles
		}
	}

	return area, angles
}

func main() {
	//file_path := "test_data.txt"
	file_path := "data.txt"

	// Use the function to read the file
	region, err := parse_data(file_path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := 0
	visited := make(map[[2]int]bool)

	// Prompt the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Advent of code day 12!")
	fmt.Println("For part one type number 1")
	fmt.Println("For part one type number 2")
	fmt.Print("What is the query?\n")

	// Read the user input
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Remove extra whitespace or newlines

	// Look up and run the chosen function
	switch input {
	case "1":
		// Process the matrix
		for i, row := range region {
			for j := range row {
				area, perimeter := find_plots(i, j, region, visited)
				result += area * perimeter
			}
		}
	case "2":
		// Process the matrix
		for i, row := range region {
			for j := range row {
				area, perimeter := find_plots2(i, j, region, visited)
				result += area * perimeter
				/*if area != 0 {
					fmt.Println("Area:", area, "Sides:", perimeter)
				}*/
			}
		}

	default:
		fmt.Println("Invalid choice. Please select a valid function.")
	}

	fmt.Println("Result:", result)
}
