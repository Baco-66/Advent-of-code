package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseFile(file_path string) ([][]rune, []rune, error) {
	// Open the file
	file, err := os.Open(file_path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Variables to store the matrix and the sequence
	var matrix [][]rune
	var sequence []rune
	isMatrix := true

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Empty line indicates the end of the matrix
			isMatrix = false
			continue
		}

		if isMatrix {
			// Add each line as a row in the matrix
			matrix = append(matrix, []rune(line))
		} else {
			// Parse the sequence
			sequence = append(sequence, []rune(line)...)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return matrix, sequence, nil
}

// Locate "@" in the matrix
func findRobot(matrix [][]rune) (int, int, bool) {
	for i, row := range matrix {
		for j, cell := range row {
			if cell == '@' {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

// Check if a position is within bounds
func isValidPosition(matrix [][]rune, x, y int) bool {
	return x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0])
}

// Recursive function to move an "O"
func moveBox(matrix [][]rune, x, y, dx, dy int) bool {
	newX, newY := x+dx, y+dy

	// Stop if out of bounds or a wall is encountered
	if !isValidPosition(matrix, newX, newY) || matrix[newX][newY] == '#' {
		return false
	}

	// If there's another "O" behind, try to move it recursively
	if matrix[newX][newY] == 'O' {
		if !moveBox(matrix, newX, newY, dx, dy) {
			return false // Blocked by a wall or end of grid
		}
	}

	// Move the current "O" to the new position
	matrix[newX][newY] = 'O'
	matrix[x][y] = '.'
	return true
}

// Move "@" in the specified direction
func moveRobot(matrix [][]rune, direction rune) bool {
	// Direction vectors for movement
	var directions = map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	// Locate the "@" position
	x, y, found := findRobot(matrix)
	if !found {
		fmt.Println("Error: '@' not found in the matrix")
		return false
	}

	// Get the movement vector for the direction
	dx, dy := directions[direction][0], directions[direction][1]
	newX, newY := x+dx, y+dy

	// Check if the new position is valid
	if !isValidPosition(matrix, newX, newY) || matrix[newX][newY] == '#' {
		return false // Cannot move into walls or out of bounds
	}

	// Handle interactions with "O"
	if matrix[newX][newY] == 'O' {
		// Try to move the "O" recursively
		if !moveBox(matrix, newX, newY, dx, dy) {
			return false // Blocked by a wall or end of grid
		}
	}

	// Normal move
	matrix[newX][newY] = '@'
	matrix[x][y] = '.'
	return true
}

func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func doubleMap(matrix [][]rune) [][]rune {

	// Iterate over each row and double its size
	for i := range matrix {
		// Create a new slice of runes with double the original width
		newRow := make([]rune, len(matrix[i])*2)

		// Fill the new line
		for j := range matrix[i] {
			if matrix[i][j] == '#' {
				newRow[j*2] = '#'
				newRow[j*2+1] = '#'
			} else if matrix[i][j] == '.' {
				newRow[j*2] = '.'
				newRow[j*2+1] = '.'
			} else if matrix[i][j] == 'O' {
				newRow[j*2] = '['
				newRow[j*2+1] = ']'
			} else if matrix[i][j] == '@' {
				newRow[j*2] = '@'
				newRow[j*2+1] = '.'
			}
		}

		// Replace the old row with the new row
		matrix[i] = newRow
	}

	return matrix
}

// Struct to represent a change
type Change struct {
	X, Y  int
	Value rune
}

// Apply the changes to the matrix
func applyChanges(matrix [][]rune, changes []Change) {
	// Create a set of positions where "[" or "]" will be placed
	occupied := make(map[[2]int]bool)

	// First pass: Mark positions for "[" and "]"
	for _, change := range changes {
		if change.Value == '[' || change.Value == ']' {
			occupied[[2]int{change.X, change.Y}] = true
		}
	}

	// Second pass: Apply changes, skipping "." if the position will have "[" or "]"
	for _, change := range changes {
		if change.Value == '.' && occupied[[2]int{change.X, change.Y}] {
			continue // Skip the "." if "[" or "]" will be placed here
		}
		matrix[change.X][change.Y] = change.Value
	}
}

// Recursive function to move an "[]"
func moveBigBoxVerticaly(matrix [][]rune, x, y, dx int) []Change {
	var changes []Change
	var yr, yl int
	if matrix[x][y] == '[' {
		yr, yl = y, y+1
	} else {
		yr, yl = y-1, y
	}
	newrY, newlY, newX := yr, yl, x+dx

	// Stop if out of bounds or a wall is encountered
	if !isValidPosition(matrix, newX, newrY) || matrix[newX][newrY] == '#' || !isValidPosition(matrix, newX, newlY) || matrix[newX][newlY] == '#' {
		return nil
	}

	// If there's another "[]" behind, try to move it recursively
	if matrix[newX][newrY] == '[' {
		if change := moveBigBoxVerticaly(matrix, newX, newrY, dx); change == nil {
			return nil // Blocked by a wall or end of grid
		} else {
			changes = append(changes, change...)
		}
	} else {
		if matrix[newX][newlY] == '[' {
			if change := moveBigBoxVerticaly(matrix, newX, newlY, dx); change == nil {
				return nil // Blocked by a wall or end of grid
			} else {
				changes = append(changes, change...)
			}
		}
		if matrix[newX][newrY] == ']' {
			if change := moveBigBoxVerticaly(matrix, newX, newrY, dx); change == nil {
				return nil // Blocked by a wall or end of grid
			} else {
				changes = append(changes, change...)
			}

		}
	}

	// Move the current "[]" to the new position

	//fmt.Println("x", x, "newrY", newrY, "x", x, "newlY", newlY, "newX", newX, "newrY", newrY, "newX", newX, "newlY", newlY)

	changes = append(changes, Change{x, newrY, '.'}, Change{x, newlY, '.'}, Change{newX, newrY, '['}, Change{newX, newlY, ']'})
	return changes
}

// Recursive function to move an "[]"
func moveBigBoxHorrizontaly(matrix [][]rune, x, y, dy int) []Change {
	var changes []Change
	var yr, yl int
	if matrix[x][y] == '[' {
		yr, yl = y, y+1
	} else {
		yr, yl = y-1, y
	}
	newrY, newlY := yr+dy, yl+dy

	// Stop if out of bounds or a wall is encountered
	if !isValidPosition(matrix, x, newrY) || matrix[x][newrY] == '#' || !isValidPosition(matrix, x, newlY) || matrix[x][newlY] == '#' {
		return nil
	}

	//fmt.Println("x", x, "yr", yr, "x", x, "yl", yl, "dy", dy)
	//fmt.Printf("%c, %c, %c, %c\n", matrix[x][yr-1], matrix[x][yr], matrix[x][yl], matrix[x][yl+1])

	// If there's another "[]" behind, try to move it recursively
	if newrY == yl && matrix[x][newlY] == '[' {
		if change := moveBigBoxHorrizontaly(matrix, x, newlY, dy); change == nil {
			return nil // Blocked by a wall or end of grid
		} else {
			changes = append(changes, change...)
		}
	}
	if newlY == yr && matrix[x][newrY] == ']' {
		if change := moveBigBoxHorrizontaly(matrix, x, newrY, dy); change == nil {
			return nil // Blocked by a wall or end of grid
		} else {
			changes = append(changes, change...)
		}
	}

	// Move the current "[]" to the new position
	changes = append(changes, Change{x, yr, '.'}, Change{x, yl, '.'}, Change{x, newrY, '['}, Change{x, newlY, ']'})
	return changes
}

// Move "@" in the specified direction
func moveRobotBig(matrix [][]rune, direction rune) bool {
	// Direction vectors for movement
	var directions = map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	// Locate the "@" position
	x, y, found := findRobot(matrix)
	if !found {
		fmt.Println("Error: '@' not found in the matrix")
		return false
	}

	// Get the movement vector for the direction
	dx, dy := directions[direction][0], directions[direction][1]
	newX, newY := x+dx, y+dy
	//fmt.Println("X", x, "Y", y, "DX", dx, "DY", dy, "NEWX", newX, "NEWY", newY)

	// Check if the new position is valid
	if !isValidPosition(matrix, newX, newY) || matrix[newX][newY] == '#' {
		return false // Cannot move into walls or out of bounds
	}

	// Handle interactions with "[]"
	if matrix[newX][newY] == '[' || matrix[newX][newY] == ']' {
		// Try to move the "[]" recursively
		if dy == 0 {
			if changes := moveBigBoxVerticaly(matrix, newX, newY, dx); changes == nil {
				return false // Blocked by a wall or end of grid
			} else {
				applyChanges(matrix, changes)
			}
		} else {
			if changes := moveBigBoxHorrizontaly(matrix, newX, newY, dy); changes == nil {
				return false // Blocked by a wall or end of grid
			} else {
				applyChanges(matrix, changes)
			}
		}

	}

	// Normal move
	matrix[newX][newY] = '@'
	matrix[x][y] = '.'
	return true
}

func main() {
	//file_path := "test_data_small.txt"
	//file_path := "test_data.txt"
	file_path := "data.txt"

	// Parse the file
	matrix, sequence, err := parseFile(file_path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := 0

	// Prompt the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Advent of code day 15!")
	fmt.Println("For part one type number 1")
	fmt.Println("For part one type number 2")
	fmt.Print("What is the query?\n")

	// Read the user input
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Remove extra whitespace or newlines

	// Look up and run the chosen function
	switch input {
	case "1":
		for _, direction := range sequence {
			moveRobot(matrix, direction)
		}

		for i, row := range matrix {
			for j, cell := range row {
				if cell == 'O' {
					result += j + (i * 100)
				}
			}
		}
	case "2":
		// Double the width of the matrix
		matrix = doubleMap(matrix)

		for _, direction := range sequence {
			moveRobotBig(matrix, direction)
		}
		printMatrix(matrix)

		for i, row := range matrix {
			for j, cell := range row {
				if cell == '[' {
					result += j + (i * 100)
				}
			}
		}

	default:
		fmt.Println("Invalid choice. Please select a valid function.")
	}

	fmt.Println("Result:", result)
}
