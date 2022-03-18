package brainfuck

import "fmt"

type command interface {
	execute(mem *memmory)
}

type operation struct {
	execute func()
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
type loopCheckBordersOperation struct {                // ]
	innerOperation *loopOperation
	operation
}
type loopOperation struct { // [
	operation
	innerLoop []command
	repeat    bool
}

func (op incrementOperation) execute(mem *memmory) {
	mem.cells[mem.pointer]++
}

func (op decrementOperation) execute(mem *memmory) {
	mem.cells[mem.pointer]--
}
func (op incrementDataPointerOperation) execute(mem *memmory) {
	mem.pointer++
}

func (op decrementDataPointerOperation) execute(mem *memmory) {
	mem.pointer--
}

func (op outputOperation) execute(mem *memmory) {
	//fmt.Printf("%c", mem.cells[mem.pointer])
	//debug
	fmt.Printf("%d", mem.cells[mem.pointer])
}

func (op inputOperation) execute(mem *memmory) {
	fmt.Scanf("%c", &mem.cells[mem.pointer])
}

func (op zeroOperation) execute(mem *memmory) {
	mem.cells[mem.pointer] = 0
}

func (op copyOperation) execute(mem *memmory) {
	copyPasteAccumulator = mem.cells[mem.pointer]
}

func (op pasteOperation) execute(mem *memmory) {
	if copyPasteAccumulator != 0 {
		mem.cells[mem.pointer] = copyPasteAccumulator
	}
}

func (op loopOperation) execute(mem *memmory) {
	for op.repeat {
		for _, innerOperation := range op.innerLoop {
			innerOperation.execute(mem)
		}
	}
}

func (op loopCheckBordersOperation) execute(mem *memmory) {
	//fmt.Println(op.innerOperation.mem)
	if mem.cells[mem.pointer] == 0 {
		op.innerOperation.repeat = false
	} else {
		op.innerOperation.repeat = true
	}
}
