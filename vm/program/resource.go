package program

import (
	"fmt"

	"github.com/formancehq/machine/core"
)

type Resource interface {
	isResource()
	GetType() core.Type
}

type Constant struct {
	Inner core.Value
}

func (Constant) isResource()          {}
func (c Constant) GetType() core.Type { return c.Inner.GetType() }
func (c Constant) String() string     { return fmt.Sprintf("%v", c.Inner) }

type Parameter struct {
	Typ  core.Type
	Name string
}

func (Parameter) isResource()          {}
func (p Parameter) GetType() core.Type { return p.Typ }
func (p Parameter) String() string     { return fmt.Sprintf("<%v %v>", p.Typ, p.Name) }

type Metadata struct {
	SourceAccount core.Address
	Key           string
	Typ           core.Type
}

func (Metadata) isResource()          {}
func (m Metadata) GetType() core.Type { return m.Typ }
func (m Metadata) String() string {
	return fmt.Sprintf("<%v meta(%v, %v)>", m.Typ, m.SourceAccount, m.Key)
}
