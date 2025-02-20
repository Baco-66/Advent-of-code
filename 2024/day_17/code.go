package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Local Variables
var A, B, C, INSTRUCTIONS_POINTER int = 0, 0, 0, 0
var INSTRUCTIONS []int

//////////////////// Message printing ////////////////////

// Printer struct to hold state variables
type Printer struct {
	result []int
}

// Global logger instance
var STDOUT = NewPrinter()

// NewPrinter initializes a Printer with a given prefix
func NewPrinter() *Printer {
	return &Printer{result: []int{}}
}

// Stores program output for later
func (p *Printer) Store(message int, args ...interface{}) {
	p.result = append(p.result, message)
}

// Print program output
func (p *Printer) Print() {
	first := true
	for _, message := range p.result { // Assuming p.result is a slice or array
		if first {
			fmt.Printf("%d", message)
			first = false // Update first after printing the first element
		} else {
			fmt.Printf(",%d", message)
		}
	}
	fmt.Printf("\n") // Ensure a newline after printing
}

// Print applies special rules and maintains state
func (p *Printer) Device_State() {
	fmt.Printf("Register: %+v\n", A)
	fmt.Printf("Register: %+v\n", B)
	fmt.Printf("Register: %+v\n", C)
	fmt.Print("Queue: ")
	fmt.Print(INSTRUCTIONS, "\n")
}

///////////// Read and parse the data from the file /////////////

func parseFile(file_path string) error {

	// Read the entire file into memory
	data, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}

	// Convert file data into lines
	lines := strings.FieldsFunc(string(data), func(r rune) bool {
		return r == '\n' || r == '\r' // Filters out empty lines
	})

	// Expecting exactly 4 lines: 3 for registers and 1 for program
	if len(lines) != 4 {
		return fmt.Errorf("invalid file format: expected 4 lines, got %d", len(lines))
	}

	// Extract register values
	getValue := func(line string) (int, error) {
		parts := strings.Split(line, ":")
		return strconv.Atoi(strings.TrimSpace(parts[1]))
	}

	if A, err = getValue(lines[0]); err != nil {
		return fmt.Errorf("failed to parse Register A: %w", err)
	}
	if B, err = getValue(lines[1]); err != nil {
		return fmt.Errorf("failed to parse Register B: %w", err)
	}
	if C, err = getValue(lines[2]); err != nil {
		return fmt.Errorf("failed to parse Register C: %w", err)
	}

	// Parse the program queue
	programData := strings.Split(strings.TrimSpace(strings.Split(lines[3], ":")[1]), ",")
	for _, part := range programData {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return fmt.Errorf("failed to parse program queue: %w", err)
		}
		INSTRUCTIONS = append(INSTRUCTIONS, num)
	}

	return nil
}

///////////// Handle the system functioning /////////////

func combo(operand int) int {
	if 0 <= operand && operand <= 3 {
		return operand
	} else if operand == 4 {
		return A
	} else if operand == 5 {
		return B
	} else if operand == 6 {
		return C
	}

	// Trigger a panic if the operand is invalid
	panic(fmt.Sprintf("Panic: invalid operand %d", operand))
}

// opcode 0
func adv(operand int) {
	denominator := 1 << combo(operand)
	A = int(A / denominator)
}

// opcode 1
func bxl(operand int) {
	B = B ^ operand
}

// opcode 2
func bst(operand int) {
	B = (combo(operand) % 8) & 7
}

// opcode 3
func jnz(operand int) {
	if A != 0 {
		INSTRUCTIONS_POINTER = operand - 2
	}
}

// opcode 4
func bxc(operand int) {
	B = B ^ C
}

// opcode 5
func out(operand int) {
	message := int(combo(operand) % 8)
	STDOUT.Store(message)
}

// opcode 6
func bdv(operand int) {
	denominator := 1 << combo(operand)
	B = int(A / denominator)
}

// opcode 7
func cdv(operand int) {
	denominator := 1 << combo(operand)
	C = int(A / denominator)
}

func run_instruction(opcode, operand int) {
	switch opcode {
	case 0:
		adv(operand)
	case 1:
		bxl(operand)
	case 2:
		bst(operand)
	case 3:
		jnz(operand)
	case 4:
		bxc(operand)
	case 5:
		out(operand)
	case 6:
		bdv(operand)
	case 7:
		cdv(operand)
	default:
		panic(fmt.Sprintf("Panic: Unknown opcode %d", opcode))
	}
}

func main() {
	// Example usage
	filename := "data.txt" // actual data
	//filename = "test_data.txt" // Test data, for faster computation
	err := parseFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	STDOUT.Device_State()

	for ; INSTRUCTIONS_POINTER+1 < len(INSTRUCTIONS); INSTRUCTIONS_POINTER += 2 {
		run_instruction(INSTRUCTIONS[INSTRUCTIONS_POINTER], INSTRUCTIONS[INSTRUCTIONS_POINTER+1])
	}
	STDOUT.Print()

	STDOUT.Device_State()

	fmt.Println()
}
