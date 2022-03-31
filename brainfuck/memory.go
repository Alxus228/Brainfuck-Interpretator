package brainfuck

const cellsSize int = 200

// Type memory mimics an actual memory's functional but has much fewer capacity.
//
// You can set any capacity to the cells array you want if initial size isn't enough for your needs.
type memory struct {
	cells   [cellsSize]byte
	pointer int
}
