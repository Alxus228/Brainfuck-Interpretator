package brainfuck

var executableCommands = map[rune]command{
	//Increment operation
	'+': increment{},
	//Decrement operation
	'-': decrement{},
	//Increment data pointer operation
	'>': incrementDataPointer{},
	//Decrement data pointer operation
	'<': decrementDataPointer{},
	//Output operation
	'.': output{},
	//Input operation
	',': input{},
	//The beginning of loop
	'[': loop{},
	//The end of loop
	']': endLoop{},
	//All the functions implemented after the ']', are not implemented in original Brainfuck language
	//Copy operation
	'c': copy{},
	//Paste operation
	'p': paste{},
	//Clear operation
	'0': zero{},
}

var memmorySet memmory
var codePointer int

var copyPasteAccumulator byte
var currentLoop = []loop{
	{},
}

func Brainfuck(code string) {
	interpretate(code)
	compile()
}

//function to call
func interpretate(code string) {
	for codePointer = 0; codePointer < len(code); codePointer++ {
		var newCommand = executableCommands[rune(code[codePointer])]
		switch t := newCommand.(type) {
		case loop:
			currentLoop = append([]loop{t}, currentLoop...)
		case endLoop:
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
	}
}
