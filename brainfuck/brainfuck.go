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
var loopsStack []loop
var copyPasteAccumulator byte
var commands []command

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
		var newCommand = executableCommands[rune(code[codePointer])]

		switch t := newCommand.(type) {
		case loopOperation:
			t.innerLoop = append(t.innerLoop, newCommand)
		default:
			commands = append(commands, newCommand)
		}
	}
}
