package compiler

import (
	"encoding/binary"

	"github.com/numary/machine/core"
	"github.com/numary/machine/vm/program"
)

func (p *parseVisitor) AppendInstruction(instructions byte) {
	p.instructions = append(p.instructions, instructions)
}

func (p *parseVisitor) PushAddress(addr core.Address) {
	p.instructions = append(p.instructions, program.OP_APUSH)
	bytes := addr.ToBytes()
	p.instructions = append(p.instructions, bytes...)
}

func (p *parseVisitor) PushInteger(val core.Number) {
	p.instructions = append(p.instructions, program.OP_IPUSH)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(val))
	p.instructions = append(p.instructions, bytes...)
}

func (p *parseVisitor) Bump(n uint) {
	p.PushInteger(core.Number(n))
	p.instructions = append(p.instructions, program.OP_BUMP)
}
