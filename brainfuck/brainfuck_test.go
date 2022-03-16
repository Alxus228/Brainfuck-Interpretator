package brainfuck

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestInterpretate(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Interpetate("+++++c>p+++++[-<+>]<.")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if out[0] != 15 {
		t.Errorf("The output of programm +++++c>p+++++[-<+>]<. should be 15, but it equals %d", out[0])
	}
}
