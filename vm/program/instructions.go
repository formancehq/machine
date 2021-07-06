package program

const (
	OP_PUSH2 = byte(iota + 1)
	OP_PUSH8
	OP_IADD
	OP_ISUB
	OP_PRINT
	OP_FAIL
	OP_SEND
)
