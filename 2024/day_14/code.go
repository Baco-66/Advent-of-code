package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Entry struct {
	Point  Point
	Vector Point
}

type Point struct {
	X, Y int
}

type ErrParseLine struct {
	Line string
}

func (e ErrParseLine) Error() string {
	return fmt.Sprintf("Failed to find data line: %s", e.Line)
}

// Function to parse a line with X and Y values
func parseLine(line string, regex *regexp.Regexp) (Entry, error) {
	var entry Entry
	matches := regex.FindStringSubmatch(line)
	if len(matches) == 5 {
		fmt.Sscanf(matches[1], "%d", &entry.Point.X)
		fmt.Sscanf(matches[2], "%d", &entry.Point.Y)
		fmt.Sscanf(matches[3], "%d", &entry.Vector.X)
		fmt.Sscanf(matches[4], "%d", &entry.Vector.Y)
		return entry, nil
	}
	return entry, ErrParseLine{line}
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
	regex := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	// Slice to store all entries
	var entries []Entry

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			// Process 3 lines into an entry
			entry, err := parseLine(line, regex)
			if err != nil {
				return nil, err // Return error if file cannot be opened
			}
			entries = append(entries, entry)
		}
	}
	return entries, nil
}

func add_vector_to_point(x, y, dx, dy, w, h int) (int, int) {
	x += dx
	if x >= w {
		x -= w
	} else if x < 0 {
		x += w
	}
	y += dy
	if y >= h {
		y -= h
	} else if y < 0 {
		y += h
	}
	return x, y
}

func (e *Entry) patrol(width, high, seconds int) {
	for range seconds {
		e.Point.X, e.Point.Y = add_vector_to_point(e.Point.X, e.Point.Y, e.Vector.X, e.Vector.Y, width, high)
	}
}

func (e Entry) find_quadrant(x_limit, y_limit int) (quadrant int) {
	if e.Point.X < x_limit {
		if e.Point.Y < y_limit {
			quadrant = 1
		} else if e.Point.Y > y_limit {
			quadrant = 3
		}
	} else if e.Point.X > x_limit {
		if e.Point.Y < y_limit {
			quadrant = 2
		} else if e.Point.Y > y_limit {
			quadrant = 4
		}
	}
	return
}

func calculate_safety(entries []Entry, width, hight int) int {
	x_limit := width / 2
	y_limit := hight / 2
	var result []int = make([]int, 4, 4)

	for _, entry := range entries {
		i := entry.find_quadrant(x_limit, y_limit)
		if i != 0 {
			result[i-1] += 1
		}
		//fmt.Println(result, x_limit, y_limit, entry.Point.X, entry.Point.Y)
	}
	return result[0] * result[1] * result[2] * result[3]
}

// Function to generate and print the grid with point counts
func printPointMap(entries []Entry, width, height int) {
	// Create a 2D grid initialized with zeros
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	// Populate the grid with point counts
	for _, entry := range entries {
		x, y := entry.Point.X, entry.Point.Y
		if x >= 0 && x < width && y >= 0 && y < height { // Ensure points are within bounds
			grid[y][x]++
		}
	}

	// Print the grid
	for _, row := range grid {
		for _, count := range row {
			if count == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(count)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Function to check if a bunch of points are connected together (based on a threshold)
func arePointsTogether(entries []Entry, threshold int) bool {
	if len(entries) == 0 {
		return false
	}

	// Define all 8 possible directions: up, down, left, right, and diagonals
	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, // Up-left, Up, Up-right
		{0, -1}, {0, 1}, // Left, Right
		{1, -1}, {1, 0}, {1, 1}, // Down-left, Down, Down-right
	}

	// Convert entries into a point map for quick lookup
	pointMap := make(map[Point]bool)
	for _, e := range entries {
		pointMap[e.Point] = true
	}

	// Counter for points that have neighbors
	connectedCount := 0

	// Iterate through all points and check for neighbors
	for _, entry := range entries {
		hasNeighbor := false
		// Check all 8 directions for neighbors
		for _, d := range directions {
			neighbor := Point{entry.Point.X + d.X, entry.Point.Y + d.Y}
			if pointMap[neighbor] { // If a neighboring point exists
				hasNeighbor = true
				break
			}
		}
		if hasNeighbor {
			connectedCount++
		}
	}

	// Print connected points for debugging
	fmt.Printf("Connected Points: %d\n", connectedCount)

	// Compare the connected points with the threshold
	return connectedCount >= threshold
}

func main() {
	//file_path := "test_data.txt"
	//var w, h int = 11, 7

	file_path := "data.txt"
	var w, h int = 101, 103

	entries, err := parse_data(file_path)

	// Check for errors during scanning
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := 0

	// Prompt the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Advent of code day 14!")
	fmt.Println("For part one type number 1")
	fmt.Println("For part one type number 2")
	fmt.Print("What is the query?\n")

	// Read the user input
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Remove extra whitespace or newlines

	// Look up and run the chosen function
	switch input {
	case "1":
		for i := range entries {
			entries[i].patrol(w, h, 1)
		}
		result = calculate_safety(entries, w, h)
	case "2": // THIS IS SO DUMB; WHAT A DUMB CHALLANGE
		time := 0
		threshold := 300
		for {
			for i := range entries {
				entries[i].patrol(w, h, 1)
			}
			time++
			if arePointsTogether(entries, threshold) {
				fmt.Println("Time ", time)
				printPointMap(entries, w, h)
				return
			}
		}

	default:
		fmt.Println("Invalid choice. Please select a valid function.")
	}

	fmt.Println("Result:", result)
}
