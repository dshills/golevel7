package golevel7

import (
	"bytes"
	"fmt"
)

// Segment is an HL7 segment
type Segment struct {
	Fields []Field
	Value  []byte
}

func (s *Segment) String() string {
	var str string
	for _, f := range s.Fields {
		str += fmt.Sprintf("Segment Field: Seq: %d Value: %s\n", f.SeqNum, f.Value)
		str += f.String()
	}
	return str
}

func (s *Segment) parse(seps *Separators) error {
	isMSH := false
	if string(s.Value[:3]) == "MSH" {
		isMSH = true
	}
	r := bytes.NewReader(s.Value)
	i := 0
	ii := 0
	seq := 0
	for {
		ch, _, _ := r.ReadRune()
		ii++
		switch {
		case isMSH && seq == 2 && ch == seps.RepSep:
			// ignore repeat separator in separator definition
		case isMSH && seq == 2 && ch == seps.EscSep:
			// ignore escape separator in separator definition
		case ch == seps.FieldSep:
			if isMSH && seq == 2 {
				// the separator list is a field in MSH seq 2
				s.forceField([]byte(s.Value[i:ii-1]), seq)
			} else {
				fld := Field{Value: s.Value[i : ii-1], SeqNum: seq}
				fld.parse(seps)
				s.Fields = append(s.Fields, fld)
			}
			i = ii
			seq++
			if isMSH && seq == 1 {
				// The field separator is itself a field for MSH seq 1
				s.forceField([]byte(string(seps.FieldSep)), seq)
				seq++
			}
		case ch == seps.RepSep:
			fld := Field{Value: s.Value[i : ii-1], SeqNum: seq}
			fld.parse(seps)
			s.Fields = append(s.Fields, fld)
			i = ii
		case ch == seps.EscSep:
			ii++
			r.ReadRune()
		case ch == eof:
			if ii > i {
				fld := Field{Value: s.Value[i : ii-1], SeqNum: seq}
				fld.parse(seps)
				s.Fields = append(s.Fields, fld)
			}
			return nil
		}
	}
}

// forceField will force the creation of a field / component / subcomponent
// This is used for separator defines in the MSH segemnt
// ...and the name forceField is cool ;)
func (s *Segment) forceField(val []byte, seq int) {
	fld := Field{Value: val, SeqNum: seq}
	cmp := Component{Value: val}
	cmp.SubComponents = append(cmp.SubComponents, SubComponent{Value: val})
	fld.Components = append(fld.Components, cmp)
	s.Fields = append(s.Fields, fld)
}

// Field returns the field with sequence number i
func (s *Segment) Field(i int) (*Field, error) {
	for idx, fld := range s.Fields {
		if fld.SeqNum == i {
			return &s.Fields[idx], nil
		}
	}
	return nil, fmt.Errorf("Field not found")
}

// AllFields returns all fields with sequence number i
func (s *Segment) AllFields(i int) ([]*Field, error) {
	flds := []*Field{}
	for idx, fld := range s.Fields {
		if fld.SeqNum == i {
			flds = append(flds, &s.Fields[idx])
		}
	}
	if len(flds) == 0 {
		return flds, fmt.Errorf("Field not found")
	}
	return flds, nil
}

// Get returns the first value specified by the Location
func (s *Segment) Get(l *Location) (string, error) {
	if l.FieldSeq == -1 {
		return string(s.Value), nil
	}
	fld, err := s.Field(l.FieldSeq)
	if err != nil {
		return "", err
	}
	return fld.Get(l)
}

// GetAll returns all values specified by the Location
func (s *Segment) GetAll(l *Location) ([]string, error) {
	vals := []string{}
	if l.FieldSeq == -1 {
		vals = append(vals, string(s.Value))
		return vals, nil
	}
	flds, err := s.AllFields(l.FieldSeq)
	if err != nil {
		return vals, err
	}
	for _, f := range flds {
		v, err := f.Get(l)
		if err != nil {
			return vals, err
		}
		vals = append(vals, v)
	}
	return vals, nil
}
