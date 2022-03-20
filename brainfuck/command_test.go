package brainfuck

import (
	"testing"
)

// Test of the interpreter
func TestIncrement(t *testing.T) {
	mem := memmory{}
	com := increment{}

	com.execute(&mem)

	if mem.cells[mem.pointer] != 1 {
		t.Errorf("Expected %d, but got %d", 1, mem.cells[mem.pointer])
	}
}
