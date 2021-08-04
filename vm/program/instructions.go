package program

const (
	OP_APUSH = byte(iota + 1)
	OP_IPUSH
	OP_IADD
	OP_ISUB
	OP_PRINT
	OP_FAIL
	OP_SEND
	OP_MAKE_ALLOTMENT
	OP_MAKE_SOURCE_IN_ORDER  // srcs(Source*N) N
	OP_MAKE_SOURCE_ALLOTMENT // srcs(Source*N) allot(allotment(N))
	OP_MAKE_SOURCE_MAXED     // src(Source) max(Monetary)
	OP_SOURCE
	OP_ALLOC
)
