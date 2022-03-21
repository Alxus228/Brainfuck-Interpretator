package brainfuck

import (
	"io/ioutil"
	"os"
	"testing"
)

// Test of the interpreter
func TestBrainfuck(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Brainfuck("+++++c>p++++++[-<+>]<.")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if out[0] != 16 {
		t.Errorf("The output of programm +++++c>p+++++[-<+>]<. should be 16, but it equals %d", out[0])
	}
}

func TestExecute(t *testing.T) {
	mem := memmory{}

	executeTests := []struct {
		com             command
		pointer         int
		value, expected byte
	}{
		{new(increment), 0, 0, 1},
		{new(increment), 0, 255, 0},
		{new(decrement), 0, 0, 255},
		{new(decrement), 1, 5, 4},
		{new(incrementDataPointer), 0, 0, 4},
		{new(incrementDataPointer), 199, 10, 0},
		{new(decrementDataPointer), 0, 0, 10},
		{new(decrementDataPointer), 2, 0, 4},
	}

	for _, test := range executeTests {
		mem.pointer = test.pointer
		mem.cells[mem.pointer] = test.value

		test.com.execute(&mem)

		if mem.cells[mem.pointer] != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, mem.cells[mem.pointer])
		}
	}
}
