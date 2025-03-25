package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const TapeSize = 30000

// Interpreter represents a Brainfuck interpreter instance
type Interpreter struct {
	tape       [TapeSize]byte
	dataPtr    int
	programPtr int
	program    string
	input      io.Reader
	output     io.Writer
}

// New creates a new Brainfuck interpreter with the given program
func New(program string) *Interpreter {
	return &Interpreter{
		program: program,
		input:   os.Stdin,
		output:  os.Stdout,
	}
}

// WithIO sets custom input and output streams for the interpreter
func (i *Interpreter) WithIO(input io.Reader, output io.Writer) *Interpreter {
	i.input = input
	i.output = output
	return i
}

// Run executes the Brainfuck program
func (i *Interpreter) Run() error {
	for i.programPtr < len(i.program) {
		switch i.program[i.programPtr] {
		case '>':
			if i.dataPtr < TapeSize-1 {
				i.dataPtr++
			}
		case '<':
			if i.dataPtr > 0 {
				i.dataPtr--
			}
		case '+':
			i.tape[i.dataPtr]++
		case '-':
			i.tape[i.dataPtr]--
		case '.':
			fmt.Fprint(i.output, string(i.tape[i.dataPtr]))
		case ',':
			reader := bufio.NewReader(i.input)
			input, err := reader.ReadByte()
			if err != nil {
				return fmt.Errorf("failed to read input: %v", err)
			}
			i.tape[i.dataPtr] = input
		case '[':
			if i.tape[i.dataPtr] == 0 {
				bracketCount := 1
				for bracketCount > 0 {
					i.programPtr++
					if i.programPtr >= len(i.program) {
						return fmt.Errorf("unmatched [")
					}
					if i.program[i.programPtr] == '[' {
						bracketCount++
					} else if i.program[i.programPtr] == ']' {
						bracketCount--
					}
				}
			}
		case ']':
			if i.tape[i.dataPtr] != 0 {
				bracketCount := 1
				for bracketCount > 0 {
					i.programPtr--
					if i.programPtr < 0 {
						return fmt.Errorf("unmatched ]")
					}
					if i.program[i.programPtr] == ']' {
						bracketCount++
					} else if i.program[i.programPtr] == '[' {
						bracketCount--
					}
				}
			}
		}
		i.programPtr++
	}
	return nil
}
