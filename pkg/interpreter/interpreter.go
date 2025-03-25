package interpreter

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const TapeSize = 30000

type Interpreter struct {
	tape       [TapeSize]byte
	dataPtr    int
	programPtr int
	program    string
	input      io.Reader
	output     io.Writer
	brackets   map[int]int
}

func New(program string) *Interpreter {
	i := &Interpreter{
		program:  program,
		input:    os.Stdin,
		output:   os.Stdout,
		brackets: make(map[int]int),
	}
	i.parseLoops()
	return i
}

func (i *Interpreter) parseLoops() {
	stack := make([]int, 0)
	for pos := 0; pos < len(i.program); pos++ {
		switch i.program[pos] {
		case '[':
			stack = append(stack, pos)
		case ']':
			if len(stack) == 0 {
				return // unmatched ']' will be caught during execution
			}
			openPos := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i.brackets[openPos] = pos
			i.brackets[pos] = openPos
		}
	}
}

func (i *Interpreter) WithIO(input io.Reader, output io.Writer) *Interpreter {
	i.input = input
	i.output = output
	return i
}

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
				if matchPos, ok := i.brackets[i.programPtr]; ok {
					i.programPtr = matchPos
				} else {
					return fmt.Errorf("unmatched [")
				}
			}
		case ']':
			if matchPos, ok := i.brackets[i.programPtr]; ok {
				if i.tape[i.dataPtr] != 0 {
					i.programPtr = matchPos
					continue
				}
			} else {
				return fmt.Errorf("unmatched ]")
			}
		}
		i.programPtr++
	}
	return nil
}
