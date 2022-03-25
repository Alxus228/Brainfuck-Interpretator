package brainfuck

// This map contains empty structures that inherit command interface
var executableCommands = map[rune]command{
	// increment
	'+': increment{},
	// decrement
	'-': decrement{},
	// increment data pointer
	'>': incrementDataPointer{},
	// decrement data pointer
	'<': decrementDataPointer{},
	// output
	'.': output{},
	// input
	',': input{},
	// the beginning of loop
	'[': loop{},
	// the end of loop
	']': endLoop{},
	// All the functions after the ']', are not implemented in original Brainfuck language
	// copy
	'c': copy{},
	// paste
	'p': paste{},
	// clear
	'0': zero{},
}

// The memmory set we're going to use everywhere.
var memmorySet memmory

// This slice of loops we're going to use as a stack for
// loop commands in which we will append commands inside "[]".
//
// After the interpretation we will execute it, but not with the
// currentLoop[0].execute(memmorySet), because we don't need to repeat it
var currentLoop = []loop{
	{}, // this is the main loop
}

// This function interpretates and compiles brainfuck code.
//
// Specification of the language can be found on Wiki: https://en.wikipedia.org/wiki/Brainfuck.
func Interpret(code string) {
	for codePointer := 0; codePointer < len(code); codePointer++ {
		var newCommand = executableCommands[rune(code[codePointer])]

		// depending on the type we will process commands differently
		switch t := newCommand.(type) {
		case loop:
			// push new loop into the top of stack
			currentLoop = append([]loop{t}, currentLoop...)
		case endLoop:
			// append the current loop into the innerLoop of the outer loop
			// if there's no outer loop, it will be appended to the main loop
			currentLoop[1].innerLoop = append(currentLoop[1].innerLoop, currentLoop[0])
			// pop current loop
			currentLoop = currentLoop[1:]
		default:
			// append a new command into the innerLoop of the current loop
			currentLoop[0].innerLoop = append(currentLoop[0].innerLoop, newCommand)
		}
	}

	compile()
}

// Function compile execute all commands in the main loop.
func compile() {
	for _, com := range currentLoop[0].innerLoop {
		com.execute(&memmorySet)
	}
}
