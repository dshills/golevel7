package golevel7

import (
	"fmt"
	"strings"
)

// Field is an HL7 field
type Field struct {
	SeqNum     int
	Components []Component
	Value      []rune
}

func (f *Field) NumSubFields() int {
	return len(f.Components)
}

func (f *Field) String() string {
	var str string
	for _, c := range f.Components {
		str += "Field Component: " + string(c.Value) + "\n"
		str += c.String()
	}
	return str
}

func (f *Field) parse(seps *Delimeters) error {
	r := strings.NewReader(string(f.Value))
	i := 0
	ii := 0
	for {
		ch, _, _ := r.ReadRune()
		ii++
		switch {
		case ch == eof || (ch == endMsg && seps.LFTermMsg):
			if ii > i {
				cmp := Component{Value: f.Value[i : ii-1]}
				cmp.parse(seps)
				f.Components = append(f.Components, cmp)
			}
			return nil
		case ch == seps.Component:
			cmp := Component{Value: f.Value[i : ii-1]}
			cmp.parse(seps)
			f.Components = append(f.Components, cmp)
			i = ii
		case ch == seps.Escape:
			ii++
			r.ReadRune()
		}
	}
}

func (f *Field) encode(seps *Delimeters) []rune {
	buf := []string{}
	for _, c := range f.Components {
		buf = append(buf, string(c.Value))
	}
	return []rune(string(strings.Join(buf, string(seps.Component))))
}

// Component returns the component i
func (f *Field) Component(i int) (*Component, error) {
	if i >= len(f.Components) {
		return nil, fmt.Errorf("Component out of range")
	}
	return &f.Components[i], nil
}

// Get returns the value specified by the Location
func (f *Field) Get(l *Location) (string, error) {
	if l.Comp == -1 {
		return string(f.Value), nil
	}
	comp, err := f.Component(l.Comp)
	if err != nil {
		return "", err
	}
	return comp.Get(l)
}

// Set will insert a value into a message at Location
func (f *Field) Set(l *Location, val string, seps *Delimeters) error {
	loc := l.Comp
	if loc < 0 {
		loc = 0
	}
	if x := loc - len(f.Components) + 1; x > 0 {
		f.Components = append(f.Components, make([]Component, x)...)
	}
	err := f.Components[loc].Set(l, val, seps)
	if err != nil {
		return err
	}
	f.Value = f.encode(seps)
	return nil
}
