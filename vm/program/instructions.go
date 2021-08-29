package program

const (
	OP_APUSH = byte(iota + 1)
	OP_IPUSH
	OP_SWAP           // <left: any> <right: any> => <right:any> <left:any>
	OP_IADD           // <number> <number> => <number>
	OP_ISUB           // <number> <number> => <number>
	OP_PRINT          // <any>
	OP_FAIL           //
	OP_ASSET          // <asset | monetary | funding> => <asset>
	OP_MAKE_ALLOTMENT // <portion>*N <int N> => <allotment(N)>
	OP_TAKE_ALL       // <account> <asset> => <funding>
	OP_TAKE           // <funding> <monetary> => <remaining: funding> <taken: funding>
	OP_TAKE_MAX       // <funding> <monetary> => <remaining: funding> <taken: funding>
	OP_TAKE_SPLIT     // <funding>*N <monetary> <allotment(N)> => <funding>*N
	OP_ASSEMBLE       // <fundings>*N <int N> => <funding>
	OP_REPAY          // <funding>
	OP_ALLOC          // <funding> <allotment(N)> => <funding>*N
	OP_SEND           // <funding> <account>
)
