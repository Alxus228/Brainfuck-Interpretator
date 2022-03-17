package brainfuck

import (
	"fmt"
)

var executableCommands = map[rune]func(*memmory){
	//Increment operation
	'+': func(mem *memmory) {
		mem.cells[mem.pointer]++
	},
	//Decrement operation
	'-': func(mem *memmory) {
		mem.cells[mem.pointer]--
	},
	//Increment data pointer operation
	'>': func(mem *memmory) {
		mem.pointer++
	},
	//Decrement data pointer operation
	'<': func(mem *memmory) {
		mem.pointer--
	},
	//Output operation
	'.': func(mem *memmory) {
		fmt.Printf("%c", mem.cells[mem.pointer])
	},
	//Input operation
	',': func(mem *memmory) {
		fmt.Scanf("%c", &mem.cells[mem.pointer])
	},
	//The beginning of loop
	'[': func(mem *memmory) {
		if mem.cells[mem.pointer] == 0 {
			if loopsStack[0].openIndex == codePointer {
				loopsStack[0].executing = false
			}
		} else {
			if loopsStack[0].openIndex != codePointer {
				loopsStack = append(make([]loop, 1), loopsStack...)
				loopsStack[0].openIndex = codePointer
				loopsStack[0].executing = true
			}
		}
	},
	//The end of loop
	']': func(mem *memmory) {
		if mem.cells[mem.pointer] == 0 || !loopsStack[0].executing {
			loopsStack = loopsStack[1:]
		} else {
			// - 1 here is necessary because we increment codePointer each iteration
			codePointer = loopsStack[0].openIndex - 1
		}
	},
	//All the functions implemented after the ']', are not implemented in original Brainfuck language
	//Clear operation
	'0': func(mem *memmory) {
		mem.cells[mem.pointer] = 0
	},
	//Copy operation
	'c': func(mem *memmory) {
		copyPasteAccumulator = mem.cells[mem.pointer]
	},
	//Paste operation
	'p': func(mem *memmory) {
		if copyPasteAccumulator != 0 {
			mem.cells[mem.pointer] = copyPasteAccumulator
		}
	},
}

var memmorySet memmory
var codePointer int
var loopsStack []loop
var copyPasteAccumulator byte

type loop struct {
	openIndex int
	executing bool
}

//function to call
func Interpetate(code string) {
	//creating "main loop"
	loopsStack = append(loopsStack, *new(loop))
	loopsStack[0].executing = true

	for codePointer = 0; codePointer < len(code); codePointer++ {
		if loopsStack[0].executing {
			executableCommands[rune(code[codePointer])](&memmorySet)
		}
	}
}
