package golevel7

import (
	"bytes"
	"fmt"
)

// Field is an HL7 field
type Field struct {
	SeqNum     int
	Components []Component
	Value      []byte
}

func (f *Field) String() string {
	var str string
	for _, c := range f.Components {
		str += "Field Component: " + string(c.Value) + "\n"
		str += c.String()
	}
	return str
}

func (f *Field) parse(seps *Separators) error {
	r := bytes.NewReader(f.Value)
	i := 0
	ii := 0
	// special case for seperators in MSH seq 2
	if string(f.Value) == seps.SepField {
		cmp := Component{Value: f.Value}
		cmp.SubComponents = append(cmp.SubComponents, SubComponent{Value: f.Value})
		f.Components = append(f.Components, cmp)
		return nil
	}
	// special case for field separator in MSH seq 1
	if string(f.Value) == string(seps.FieldSep) {
		cmp := Component{Value: f.Value}
		cmp.SubComponents = append(cmp.SubComponents, SubComponent{Value: f.Value})
		f.Components = append(f.Components, cmp)
		return nil
	}
	for {
		ch, _, _ := r.ReadRune()
		ii++
		switch {
		case ch == seps.ComSep:
			cmp := Component{Value: f.Value[i : ii-1]}
			cmp.parse(seps)
			f.Components = append(f.Components, cmp)
			i = ii
		case ch == seps.EscSep:
			ii++
			r.ReadRune()
		case ch == eof:
			if ii > i {
				cmp := Component{Value: f.Value[i : ii-1]}
				cmp.parse(seps)
				f.Components = append(f.Components, cmp)
			}
			return nil
		}
	}
}

// Component returns the component i
func (f *Field) Component(i int) (*Component, error) {
	if i >= len(f.Components) {
		return nil, fmt.Errorf("Component out of range")
	}
	return &f.Components[i], nil
}
