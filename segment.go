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
		case ch == seps.FieldSep:
			fld := Field{Value: s.Value[i : ii-1], SeqNum: seq}
			fld.parse(seps)
			s.Fields = append(s.Fields, fld)
			i = ii
			seq++
		case ch == seps.RepSep:
			if isMSH && seq == 1 {
				// ignore repeat in separator definition
				continue
			}
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

// Field returns the field with sequence number i
func (s *Segment) Field(i int) (*Field, error) {
	for idx, fld := range s.Fields {
		if fld.SeqNum == i {
			return &s.Fields[idx], nil
		}
	}
	return nil, fmt.Errorf("Field not found")
}
