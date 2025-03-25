# Brainfuck Interpreter

A fast and lightweight Brainfuck interpreter written in Go. This interpreter comes with both a command-line interface and a reusable package that can be integrated into other Go projects.

## Features

- Full Brainfuck language support
- 30,000 cell memory tape
- Command-line interface
- Reusable interpreter package
- Configurable I/O streams
- Error handling for unmatched brackets

## Installation

```bash
go install github.com/kanishka-sahoo/brainfuck@latest
```

Or clone and build manually:

```bash
git clone https://github.com/kanishka-sahoo/brainfuck.git
cd brainfuck
go build
```

## Usage

### Command Line Interface

Run a Brainfuck program directly:

```bash
./brainfuck "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
```

Run a program from a file:

```bash
./brainfuck -f program.bf
```

### As a Package

The interpreter can be used as a package in your Go programs:

```go
package main

import (
    "github.com/kanishka-sahoo/brainfuck/pkg/interpreter"
    "strings"
)

func main() {
    // Create a new interpreter instance
    program := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++."
    bf := interpreter.New(program)

    // Optionally configure custom I/O
    input := strings.NewReader("some input")
    var output strings.Builder
    bf.WithIO(input, &output)

    // Run the program
    if err := bf.Run(); err != nil {
        panic(err)
    }
}
```

## Memory Model

The interpreter uses a tape of 30,000 cells, each containing an 8-bit unsigned integer (0-255). The tape pointer starts at cell 0 and can move left or right (but not beyond tape boundaries).

## Brainfuck Commands

- `>` Move the pointer right
- `<` Move the pointer left
- `+` Increment the current cell
- `-` Decrement the current cell
- `.` Output the current cell value as an ASCII character
- `,` Input a character and store its ASCII value
- `[` Jump forward to matching `]` if current cell is 0
- `]` Jump back to matching `[` if current cell is not 0

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
