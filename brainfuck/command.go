package brainfuck

import "fmt"

type command interface {
	execute(mem *memmory)
}

type operation struct {
	mem *memmory
}

type incrementOperation struct{ operation }            // +
type decrementOperation struct{ operation }            // -
type incrementDataPointerOperation struct{ operation } // >
type decrementDataPointerOperation struct{ operation } // <
type outputOperation struct{ operation }               // .
type inputOperation struct{ operation }                // ,
type zeroOperation struct{ operation }                 // 0
type copyOperation struct{ operation }                 // c
type pasteOperation struct{ operation }                // p

func (op *incrementOperation) execute() {
	op.mem.cells[op.mem.pointer]++
}

func (op *decrementOperation) execute() {
	op.mem.cells[op.mem.pointer]--
}
func (op *incrementDataPointerOperation) execute() {
	op.mem.pointer++
}

func (op *decrementDataPointerOperation) execute() {
	op.mem.pointer--
}

func (op *outputOperation) execute() {
	fmt.Printf("%c", op.mem.cells[op.mem.pointer])
}

func (op *inputOperation) execute() {
	fmt.Scanf("%c", &op.mem.cells[op.mem.pointer])
}

func (op *zeroOperation) execute() {
	op.mem.cells[op.mem.pointer] = 0
}

func (op *copyOperation) execute() {
	copyPasteAccumulator = op.mem.cells[op.mem.pointer]
}

func (op *pasteOperation) execute() {
	if copyPasteAccumulator != 0 {
		op.mem.cells[op.mem.pointer] = copyPasteAccumulator
	}
}
