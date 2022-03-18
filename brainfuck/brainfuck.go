package brainfuck

import (
	"fmt"
	"reflect"
)

var executableCommands = map[rune]command{
	//Increment operation
	'+': incrementOperation{},
	//Decrement operation
	'-': decrementOperation{},
	//Increment data pointer operation
	'>': incrementDataPointerOperation{},
	//Decrement data pointer operation
	'<': decrementDataPointerOperation{},
	//Output operation
	'.': outputOperation{},
	//Input operation
	',': inputOperation{},
	//The beginning of loop
	'[': loopOperation{},
	//The end of loop
	']': loopCheckBordersOperation{
		innerOperation: &currentLoop[0],
	},
	//All the functions implemented after the ']', are not implemented in original Brainfuck language
	//Clear operation
	'0': zeroOperation{},
	//Copy operation
	'c': copyOperation{},
	//Paste operation
	'p': pasteOperation{},
}

var memmorySet memmory
var codePointer int

var copyPasteAccumulator byte
var currentLoop = []loopOperation{
	{operation: operation{}},
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
			check := loopCheckBordersOperation{innerOperation: &currentLoop[0], operation: operation{}}
			currentLoop[1].innerLoop = append(currentLoop[1].innerLoop, check)
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
		com.execute(&memmorySet)

		//debug

		fmt.Print(reflect.TypeOf(com))
		fmt.Print(com)
		fmt.Printf("%d %d\n", memmorySet.pointer, memmorySet.cells[memmorySet.pointer])
		switch t := com.(type) {
		case loopOperation:
			fmt.Println(t.dontRepeat)
			for _, funct := range t.innerLoop {
				fmt.Print(reflect.TypeOf(funct))
				fmt.Println(funct)
				switch s := funct.(type) {
				case loopCheckBordersOperation:
					fmt.Println(s.innerOperation.dontRepeat)
				}
			}
		}
		/**/
	}
}
