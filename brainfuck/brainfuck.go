package brainfuck

import "fmt"

var executableCommands = map[rune]func(*memmory){
	'+': func(mem *memmory) {
		mem.cells[mem.pointer]++
	},
	'-': func(mem *memmory) {
		mem.cells[mem.pointer]--
	},
	'>': func(mem *memmory) {
		mem.pointer++
	},
	'<': func(mem *memmory) {
		mem.pointer--
	},
	'.': func(mem *memmory) {
		fmt.Printf("%c", mem.cells[mem.pointer])
	},
	',': func(mem *memmory) {
		fmt.Scanf("%c", &mem.cells[mem.pointer])
	},
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
	']': func(mem *memmory) {
		if mem.cells[mem.pointer] == 0 || !loopsStack[0].executing {
			loopsStack = loopsStack[1:]
		} else {
			// - 1 here is necessary because we increment codePointer each iteration
			codePointer = loopsStack[0].openIndex - 1
		}
	},
}

var memmorySet memmory
var codePointer int
var loopsStack []loop

type loop struct {
	openIndex int
	executing bool
}

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
