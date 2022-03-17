package brainfuck

var executableCommands = map[rune]command{
	//Increment operation
	'+': new(incrementOperation),
	//Decrement operation
	'-': new(decrementOperation),
	//Increment data pointer operation
	'>': new(incrementDataPointerOperation),
	//Decrement data pointer operation
	'<': new(decrementDataPointerOperation),
	//Output operation
	'.': new(outputOperation),
	//Input operation
	',': new(inputOperation),
	//The beginning of loop
	'[': new(loopOperation),
	//The end of loop
	']': new(loopCheckLoopBordersOperation),
	//All the functions implemented after the ']', are not implemented in original Brainfuck language
	//Clear operation
	'0': new(zeroOperation),
	//Copy operation
	'c': new(copyOperation),
	//Paste operation
	'p': new(pasteOperation),
}

var memmorySet memmory
var codePointer int

var copyPasteAccumulator byte
var commands []command
var currentLoop []loopOperation

func Brainfuck(code string) {

	//creating main loop
	mainLoop := loopOperation{operation: operation{mem: &memmorySet}}
	currentLoop = append(currentLoop, mainLoop)

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
		case loopCheckLoopBordersOperation:
			currentLoop = currentLoop[1:]
			currentLoop[0].innerLoop = append(currentLoop[0].innerLoop, newCommand)
		default:
			currentLoop[0].innerLoop = append(currentLoop[0].innerLoop, newCommand)
		}
	}
}

func compile() {
	for _, com := range currentLoop[0].innerLoop {
		com.execute()
	}
}
