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

// Clears the buffer
func (p *Printer) clear_buffer() {
	p.result = []int{}
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
	A = int(A >> combo(operand))
}

// opcode 1
func bxl(operand int) {
	B = B ^ operand
}

// opcode 2
func bst(operand int) {
	B = combo(operand) % 8
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
	B = int(A >> combo(operand))
}

// opcode 7
func cdv(operand int) {
	C = int(A >> combo(operand))
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

// Solves part 1 of the challange
func solve_program(debug bool) {
	var temp1, temp2, temp3 int = A, B, C
	//STDOUT.Device_State()
	for ; INSTRUCTIONS_POINTER+1 < len(INSTRUCTIONS); INSTRUCTIONS_POINTER += 2 {
		if debug {
			fmt.Print("Instruction: ", INSTRUCTIONS[INSTRUCTIONS_POINTER], INSTRUCTIONS[INSTRUCTIONS_POINTER+1], "\n")

		}
		run_instruction(INSTRUCTIONS[INSTRUCTIONS_POINTER], INSTRUCTIONS[INSTRUCTIONS_POINTER+1])
		if debug {
			STDOUT.Device_State()
			STDOUT.Print()
		}
	}
	//STDOUT.Print()

	INSTRUCTIONS_POINTER = 0
	A = temp1
	B = temp2
	C = temp3

}

/*
func find_size() {
	solve_program()
	step, direction := 100, 1 // Start with step = 100, moving forward

	for len(STDOUT.result) != len(INSTRUCTIONS) {
		adjustStep(&step, &direction, len(STDOUT.result), len(INSTRUCTIONS))

		// Move A in the chosen direction
		A += direction * step

		// Re-run the program to check results
		STDOUT.clear_buffer()
		solve_program()
	}
}

// adjustStep updates the step size and direction based on overshooting
func adjustStep(step, direction *int, currentSize, targetSize int) {
	if currentSize < targetSize { // Need to increase A
		if *direction != 1 { // Change direction to forward
			*direction = 1
			*step = (*step / 2) + 1
		} else {
			*step *= 2 // Exponential growth when in the right direction
		}
	} else { // Need to decrease A
		if *direction != -1 { // Change direction to backward
			*direction = -1
			*step = (*step / 2) + 1
		} else {
			*step *= 2
		}
	}
}

// Atempt at solving part 2
func search() {

	var originalA int = A

	for len(STDOUT.result) == len(INSTRUCTIONS) {

		STDOUT.clear_buffer()
		solve_program()
		matching := true
		for i := range STDOUT.result {
			if STDOUT.result[i] != INSTRUCTIONS[i] {
				matching = false
				break
			}
		}
		if matching {
			return // Found the correct value, exit
		}
		A += 1
		// Check

	}

	A = originalA
	STDOUT.clear_buffer()
	solve_program()


	for len(STDOUT.result) == len(INSTRUCTIONS) {
		STDOUT.clear_buffer()
		solve_program()
		matching := true
		for i := range STDOUT.result {
			if STDOUT.result[i] != INSTRUCTIONS[i] {
				matching = false
				break
			}
		}
		if matching {
			return // Found the correct value, exit
		}
		A -= 1
		// Check

	}
}
*/

func reserse_engenier() {
	A = 1
	var match bool

	//reader := bufio.NewReader(os.Stdin)

	for n := 0; n <= len(INSTRUCTIONS); {

		for i := A % 8; i < 8; i++ {

			match = true

			//fmt.Printf("Running with value of A = %b\n", A)
			//reader.ReadString('\n')

			solve_program(false)

			fmt.Print("N : ", n, "; len = ", len(STDOUT.result), STDOUT.result, "\n")

			//fmt.Print(INSTRUCTIONS, "\n")
			//fmt.Print(STDOUT.result, "\n")

			for j := 0; j <= n; j++ { // Change the order of the iteration to make this a tiny bit faster

				//fmt.Print("j = ", len(STDOUT.result)-j-1, "\n")

				if len(STDOUT.result)-j > 0 {
					if STDOUT.result[len(STDOUT.result)-j-1] != INSTRUCTIONS[len(INSTRUCTIONS)-j-1] {
						match = false
						break
					}
				} else {
					break
				}
			}

			if match {
				//fmt.Print("Found a Match and exited.\n", "A = ", A, "\nInstructions = ", INSTRUCTIONS, "\nResult = ", STDOUT.result, "\n")

				n++
				A = A << 3

				STDOUT.clear_buffer()

				break
			}
			A++
			STDOUT.clear_buffer()
		}
		if !match {
			A--
			//fmt.Print("Match not found.\n")
			A = A >> 3
			for A%8 == 7 {
				A = A >> 3
			}
			A++
			n--

		}

	}
	A = A >> 3
}

// Part 2 was a tought nut to crack.
// In my first atemtps, I was trying to brute force the correct result, but doing some math, i could see the complexety of the problem is 8^n, where n is the size of the output
// After looking online for help, i started trying to reverse engineer the program but I lost a lof of time trying to figure out how I could possibly reverse engenier all possible programs
// Turns out, you dont need to do that. Investigating the program further, we can do a miss match of both aproaches, reducing brute force timing to 8n, which is realy good.

func main() {
	// Example usage
	filename := "data.txt" // actual data
	//filename = "test_data.txt" // Test data, for faster computation
	err := parseFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	reserse_engenier()

	solve_program(false)
	STDOUT.Device_State()
	STDOUT.Print()
}
