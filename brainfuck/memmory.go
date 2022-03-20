package brainfuck

// Type memmory mimics an actual memmory's functional but has much fewer capaticy.
//
// You can set any capaticy to the cells array you want if initial size isn't enough for your needs.
type memmory struct {
	cells   [200]byte
	pointer int
}
