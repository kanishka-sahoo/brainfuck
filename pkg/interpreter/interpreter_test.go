package interpreter

import (
	"strings"
	"testing"
)

func TestBasicOperations(t *testing.T) {
	tests := []struct {
		name     string
		program  string
		input    string
		expected string
	}{
		{
			name:     "increment",
			program:  "+++.",
			expected: string([]byte{3}),
		},
		{
			name:     "decrement",
			program:  "+++--.",
			expected: string([]byte{1}),
		},
		{
			name:     "move right and increment",
			program:  "+>++>+++.",
			expected: string([]byte{3}),
		},
		{
			name:     "move left",
			program:  "+>++>+++<.<.",
			expected: string([]byte{2, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output strings.Builder
			bf := New(tt.program)
			bf.WithIO(strings.NewReader(""), &output)

			err := bf.Run()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("got %q, want %q", output.String(), tt.expected)
			}
		})
	}
}

func TestLoops(t *testing.T) {
	tests := []struct {
		name     string
		program  string
		expected string
	}{
		{
			name:     "simple loop",
			program:  "++[>+<-]>.",
			expected: string([]byte{2}),
		},
		{
			name:     "nested loops",
			program:  "++[>++[>+<-]<-]>>.",
			expected: string([]byte{4}),
		},
		{
			name:     "skip loop if zero",
			program:  "[+++].",
			expected: string([]byte{0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output strings.Builder
			bf := New(tt.program)
			bf.WithIO(strings.NewReader(""), &output)

			err := bf.Run()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("got %q, want %q", output.String(), tt.expected)
			}
		})
	}
}

func TestInput(t *testing.T) {
	tests := []struct {
		name     string
		program  string
		input    string
		expected string
	}{
		{
			name:     "read and write",
			program:  ",.",
			input:    "A",
			expected: "A",
		},
		{
			name:     "read and increment",
			program:  ",+.",
			input:    "A",
			expected: "B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output strings.Builder
			bf := New(tt.program)
			bf.WithIO(strings.NewReader(tt.input), &output)

			err := bf.Run()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("got %q, want %q", output.String(), tt.expected)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		name        string
		program     string
		expectedErr string
	}{
		{
			name:        "unmatched opening bracket",
			program:     "[",
			expectedErr: "unmatched [",
		},
		{
			name:        "unmatched closing bracket",
			program:     "]",
			expectedErr: "unmatched ]",
		},
		{
			name:        "multiple unmatched brackets",
			program:     "[[[]",
			expectedErr: "unmatched [",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bf := New(tt.program)
			err := bf.Run()

			if err == nil {
				t.Error("expected error but got none")
			}

			if err.Error() != tt.expectedErr {
				t.Errorf("got error %q, want %q", err.Error(), tt.expectedErr)
			}
		})
	}
}

func TestTapeBounds(t *testing.T) {
	tests := []struct {
		name     string
		program  string
		expected string
	}{
		{
			name:     "move right beyond bounds",
			program:  string(make([]byte, TapeSize+1)),
			expected: "",
		},
		{
			name:     "move left beyond bounds",
			program:  "<",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output strings.Builder
			bf := New(tt.program)
			bf.WithIO(strings.NewReader(""), &output)

			err := bf.Run()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("got %q, want %q", output.String(), tt.expected)
			}
		})
	}
}

func TestHelloWorld(t *testing.T) {
	program := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	expected := "Hello World!\n"

	var output strings.Builder
	bf := New(program)
	bf.WithIO(strings.NewReader(""), &output)

	err := bf.Run()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if output.String() != expected {
		t.Errorf("got %q, want %q", output.String(), expected)
	}
}
