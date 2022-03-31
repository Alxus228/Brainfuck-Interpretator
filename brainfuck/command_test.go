package brainfuck

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	mem := memory{}

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
		{new(zero), 3, 13, 0},
		{new(zero), 87, 87, 0},
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

func TestCopyPasteExecute(t *testing.T) {
	mem := memory{}

	executeTests := []struct {
		com             command
		pointer         int
		value, expected byte
	}{
		{new(copy), 0, 14, 14},
		{new(paste), 1, 20, 14},
		{new(copy), 2, 255, 255},
		{new(paste), 3, 254, 255},
		{new(copy), 10, 70, 70},
		{new(paste), 20, 0, 70},
	}

	for _, test := range executeTests {
		mem.pointer = test.pointer
		mem.cells[mem.pointer] = test.value

		switch c := test.com.(type) {
		case copy:
			c.execute(&mem)
			if copyPasteAccumulator != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, copyPasteAccumulator)
			}
		case paste:
			c.execute(&mem)
			if mem.cells[mem.pointer] != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, mem.cells[mem.pointer])
			}
		}
	}
}

func TestOutput(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var com output
	var mem memory

	mem.pointer = 0
	mem.cells[mem.pointer] = 100

	com.execute(&mem)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if out[0] != 100 {
		t.Errorf("Expected %d, but got %d", 100, mem.cells[mem.pointer])
	}
}

func TestLoopExecute(t *testing.T) {
	mem := memory{}

	executeTests := []struct {
		com             command
		pointer         int
		value, expected byte
	}{
		{
			//move value from one cell to another
			loop{
				[]command{incrementDataPointer{}, increment{}, decrementDataPointer{}, decrement{}},
			},
			0,
			3,
			0},

		{
			//increment so many times that value becomes zero
			loop{
				[]command{increment{}},
			},
			0,
			100,
			0},
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
