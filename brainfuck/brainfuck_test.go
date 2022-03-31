package brainfuck

import (
	"io/ioutil"
	"os"
	"testing"
)

// Test of the interpreter
func TestInterpret(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Interpret("+++++c>p++++++[-<+>]<.")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if out[0] != 16 {
		t.Errorf("The output of programm +++++c>p+++++[-<+>]<. should be 16, but it equals %d", out[0])
	}
}
