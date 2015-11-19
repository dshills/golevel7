package golevel7

import (
	"bytes"
	"fmt"
)

//Component is an HL7 component
type Component struct {
	SubComponents []SubComponent
	Value         []byte
}

func (c *Component) String() string {
	var str string
	for _, s := range c.SubComponents {
		str += "Component SubComponent: " + string(s.Value) + "\n"
	}
	return str
}

func (c *Component) parse(seps *Separators) error {
	r := bytes.NewReader(c.Value)
	i := 0
	ii := 0
	for {
		ch, _, _ := r.ReadRune()
		ii++
		switch {
		case ch == seps.SubComSep:
			scmp := SubComponent{Value: c.Value[i : ii-1]}
			c.SubComponents = append(c.SubComponents, scmp)
			i = ii
		case ch == seps.EscSep:
			ii++
			r.ReadRune()
		case ch == eof:
			if ii > i {
				scmp := SubComponent{Value: c.Value[i : ii-1]}
				c.SubComponents = append(c.SubComponents, scmp)
			}
			return nil
		}
	}
}

// SubComponent returns the subcomponent i
func (c *Component) SubComponent(i int) (*SubComponent, error) {
	if i >= len(c.SubComponents) {
		return nil, fmt.Errorf("SubComponent out of range")
	}
	return &c.SubComponents[i], nil
}

func (c *Component) encode(seps *Separators) []byte {
	buf := [][]byte{}
	for _, sc := range c.SubComponents {
		buf = append(buf, sc.Value)
	}
	return bytes.Join(buf, []byte(string(seps.SubComSep)))
}

// Get returns the value specified by the Location
func (c *Component) Get(l *Location) (string, error) {
	if l.SubComp == -1 {
		return string(c.Value), nil
	}
	sc, err := c.SubComponent(l.SubComp)
	if err != nil {
		return "", err
	}
	return string(sc.Value), nil
}
