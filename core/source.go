package core

type Source struct {
	inner ISource
}

func NewSource(src ISource) Source {
	return Source{inner: src}
}

type ISource interface {
	isISource()
}

func (Account) isISource()         {}
func (SourceInOrder) isISource()   {}
func (SourceAllotment) isISource() {}
func (SourceMaxed) isISource()     {}

type SourceInOrder []Source

type SourceAllotment struct {
	allot   Allotment
	sources []Source
}

type SourceMaxed struct {
	max    Monetary
	source Source
}
