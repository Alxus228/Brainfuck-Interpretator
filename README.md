# Brainfuck Interpreter
_Brainfuck-Interpretator_ is a project that can _execute_ any **brainfuck** code you want.

Moreover, you can _add your own commands_, guiding by initial commands in **commands.go** file.

If you want to execute your own code instead of the '_Hello World!!_', clone this repository and change the code in **main.go** in a line that calls **brainfuck.Brainfuck()** function.

![code example](https://user-images.githubusercontent.com/40440883/159159417-af8bf1dd-8264-45fc-85ea-97aa12e4a4e9.jpg)

### Installation
To install brainfuck package and succesfully use it in your project, you should:
* **Clone it** into a _new folder_.
* **Cut out** **_brainfuck_** folder, and **paste** it into your project root.
* **Add lines** `github.com/Alxus228/Brainfuck-Interpretator/brainfuck` in import in `.go` file, where you're going to use it, and `module github.com/Alxus228/Brainfuck-Interpretator` into `go.mod` file

If you want to use the project in this repository, you should just:
* **Clone it** into a folder with an appropriate name.

### Testing
If you changed the brainfuck package, and want to check that you haven't broken anything:
Run `go test ./...` in terminal.
To check test coverage, run `go test ./... --cover`, respectively.
