package brainfuck

import "fmt"

// Interface command describes method execute() that recieves memmory pointer mem.
//
// The method execute() is called in compiling part of the interpretator,
// so inherit this interface if you want to add a new operation.
type command interface {
	execute(mem *memmory)
}

type incrementOperation struct{}            // +
type decrementOperation struct{}            // -
type incrementDataPointerOperation struct{} // >
type decrementDataPointerOperation struct{} // <
type outputOperation struct{}               // .
type inputOperation struct{}                // ,
type copyOperation struct{}                 // c
type pasteOperation struct{}                // p
type zeroOperation struct{}                 // 0
type endLoopOperation struct{}              // ]
type loopOperation struct {                 // [
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
	fmt.Printf("%c", mem.cells[mem.pointer])
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
	for mem.cells[mem.pointer] != 0 {
		for _, innerOperation := range op.innerLoop {
			innerOperation.execute(mem)
		}
	}
}

func (op endLoopOperation) execute(mem *memmory) {}
