package brainfuck

import (
	"fmt"
	"reflect"
)

var executableCommands = map[rune]command{
	//Increment operation
	'+': incrementOperation{operation{mem: &memmorySet}},
	//Decrement operation
	'-': decrementOperation{operation{mem: &memmorySet}},
	//Increment data pointer operation
	'>': incrementDataPointerOperation{operation{mem: &memmorySet}},
	//Decrement data pointer operation
	'<': decrementDataPointerOperation{operation{mem: &memmorySet}},
	//Output operation
	'.': outputOperation{operation{mem: &memmorySet}},
	//Input operation
	',': inputOperation{operation{mem: &memmorySet}},
	//The beginning of loop
	'[': loopOperation{operation: operation{mem: &memmorySet}},
	//The end of loop
	']': loopCheckBordersOperation{&currentLoop[0]},
	//All the functions implemented after the ']', are not implemented in original Brainfuck language
	//Clear operation
	'0': zeroOperation{operation{mem: &memmorySet}},
	//Copy operation
	'c': copyOperation{operation{mem: &memmorySet}},
	//Paste operation
	'p': pasteOperation{operation{mem: &memmorySet}},
}

var memmorySet memmory
var codePointer int

var copyPasteAccumulator byte
var currentLoop = []loopOperation{
	{operation: operation{mem: &memmorySet}},
}

func Brainfuck(code string) {
	interpetate(code)
	compile()
}

//function to call
func interpetate(code string) {
	for codePointer = 0; codePointer < len(code); codePointer++ {
		var newCommand = executableCommands[rune(code[codePointer])]
		switch t := newCommand.(type) {
		case loopOperation:
			currentLoop = append([]loopOperation{t}, currentLoop...)
			check := loopCheckBordersOperation{innerOperation: &currentLoop[0]}
			check.execute()
			//debug
			fmt.Println(currentLoop[0].repeat)
			fmt.Println(*check.innerOperation)
		case loopCheckBordersOperation:
			t.innerOperation = &currentLoop[0]
			currentLoop[0].innerLoop = append(currentLoop[0].innerLoop, t)
			currentLoop[1].innerLoop = append(currentLoop[1].innerLoop, currentLoop[0])
			currentLoop = currentLoop[1:]
		default:
			currentLoop[0].innerLoop = append(currentLoop[0].innerLoop, newCommand)
		}
	}
}

func compile() {
	for _, com := range currentLoop[0].innerLoop {
		com.execute()

		fmt.Print(reflect.TypeOf(com))
		fmt.Print(currentLoop[0].mem.cells[currentLoop[0].mem.pointer])
		fmt.Println(currentLoop[0].mem.pointer)
		switch t := com.(type) {
		case loopOperation:
			fmt.Println(t.mem, t.repeat)
			for _, funct := range t.innerLoop {
				fmt.Println(reflect.TypeOf(funct))
			}
		}
	}
}
