package brainfuck

import "fmt"

// Interface command describes method execute() that recieves memmory pointer mem.
//
// The method execute() is called in compiling part of the interpretator,
// so inherit this interface if you want to add a new operation.
type command interface {
	execute(mem *memmory)
}

// Type increment implements increasing data by 1.
type increment struct{} // It relates to the '+' character.
// Type decrement implements decreasing data by 1.
type decrement struct{} // It relates to the '-' character.
// Type incrementDataPointer implements increasing data pointer by 1.
type incrementDataPointer struct{} // It relates to the '>' character.
// Type decrementDataPointer implements decreasing data pointer by 1.
type decrementDataPointer struct{} // It relates to the '<' character.
// Type output implements printing 1 character.
type output struct{} // It relates to the '.' character.
// Type input implements input value assigment to the current data cells.
type input struct{} // It relates to the ',' character.
// Type copy implements copying the current data byte into the buffer.
type copy struct{} // It relates to the 'c' character.
// Type paste implements copying the buffer value into the current data byte.
type paste struct{} // It relates to the 'p' character.
// Type zero implements setting to 0 the current data byte.
type zero struct{} // It relates to the '0' character.
// Type endLoop implements nothing and needed only for interpretate() to recognize the bounds of a loop.
type endLoop struct{} // It relates to the ']' character.
// Type loop implements innerLoop variable - a slice of commands which can be executed in execute() method.
type loop struct {
	innerLoop []command
} // It relates to the '[' character.

func (com increment) execute(mem *memmory) {
	mem.cells[mem.pointer]++
}

func (com decrement) execute(mem *memmory) {
	mem.cells[mem.pointer]--
}
func (com incrementDataPointer) execute(mem *memmory) {
	mem.pointer++
}

func (com decrementDataPointer) execute(mem *memmory) {
	mem.pointer--
}

func (com output) execute(mem *memmory) {
	fmt.Printf("%c", mem.cells[mem.pointer])
}

func (com input) execute(mem *memmory) {
	fmt.Scanf("%c", &mem.cells[mem.pointer])
}

func (com zero) execute(mem *memmory) {
	mem.cells[mem.pointer] = 0
}

func (com copy) execute(mem *memmory) {
	copyPasteAccumulator = mem.cells[mem.pointer]
}

func (com paste) execute(mem *memmory) {
	if copyPasteAccumulator != 0 {
		mem.cells[mem.pointer] = copyPasteAccumulator
	}
}

func (com loop) execute(mem *memmory) {
	for mem.cells[mem.pointer] != 0 {
		for _, innerCommand := range com.innerLoop {
			innerCommand.execute(mem)
		}
	}
}

func (com endLoop) execute(mem *memmory) {}

// This variable is needed for copy and paste commands and can't
// be deleted or renamed without modification of those commands.
var copyPasteAccumulator byte
