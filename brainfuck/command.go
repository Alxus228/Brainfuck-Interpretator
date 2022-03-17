package brainfuck

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

func (op *incrementOperation) execute() {
	op.mem.cells[op.mem.pointer]++
}

func (op *decrementOperation) execute() {
	op.mem.cells[op.mem.pointer]--
}
