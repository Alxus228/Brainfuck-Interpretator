// Package brainfuck implements executing code, written in the same named language.
// To use it, you should simply call function Interpret, with code that has to be
// executed as parameter, in format of a string.
//
// If you want to assure yourself that the package works correctly, you can use
// online brainfuck compiler: https://copy.sh/brainfuck/ or anything else you want.
//
// Specification of the language can be found on Wiki: https://en.wikipedia.org/wiki/Brainfuck.
package brainfuck

// Interpret is an API function that receives brainfuck code as a string argument, and executes it.
func Interpret(code string) {
	// This slice of loops we're going to use as a stack for
	// loop commands in which we will append commands inside "[]".
	//
	// After the interpretation we will execute it, but not with the
	// currentLoop[0].execute(memorySet), because we don't need to repeat it
	var currentLoop = []loop{
		{}, // this is the main loop
	}

	// this is a cycle, in which we, depending on the symbol from 'code' string, push commands into the currentLoop stack
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

	compile(currentLoop[0].innerLoop)
}

// Function compile execute all commands in the main loop.
func compile(mainLoop []command) {
	// The memory set that we're going to pass to our commands.
	var memorySet memory

	for _, com := range mainLoop {
		com.execute(&memorySet)
	}
}

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
