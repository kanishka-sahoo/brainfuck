package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kanishka-sahoo/brainfuck/pkg/interpreter"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `Brainfuck Interpreter

Usage:
  %s [flags] <program>
  %s -f <filename>

Flags:
  -f    Read program from file instead of command line

Examples:
  %s "++++++++++>-----"
  %s -f program.bf
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "Read program from file")
	flag.Parse()

	var program string
	if filename != "" {
		// Read program from file specified by -f flag
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		program = string(content)
	} else {
		// Get program from command line argument
		args := flag.Args()
		if len(args) != 1 {
			printUsage()
			os.Exit(1)
		}
		program = args[0]
	}

	bf := interpreter.New(program)
	if err := bf.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
