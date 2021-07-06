package core

type ValueType = byte

const (
	TYPE_ADDRESS = ValueType(iota + 1)
	TYPE_ASSET
	TYPE_NUMBER
	TYPE_MONETARY
)

type Value interface {
	GetType() ValueType
}

type Address string

func (*Address) GetType() ValueType { return TYPE_ADDRESS }

type Asset string

func (*Asset) GetType() ValueType { return TYPE_ASSET }

type Number uint64

func (*Number) GetType() ValueType { return TYPE_NUMBER }

type Monetary struct {
	Asset  string
	Number uint64
}

func (*Monetary) GetType() ValueType { return TYPE_MONETARY }
