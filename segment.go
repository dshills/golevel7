package golevel7

import (
	"fmt"
	"errors"
	"strings"
)

// Segment is an HL7 segment
type Segment struct {
	Fields []Field
	Value  []rune
	maxSeq int
}

func (s *Segment) String() string {
	var str string
	for _, f := range s.Fields {
		str += fmt.Sprintf("Segment Field: Seq: %d Value: %s\n", f.SeqNum, string(f.Value))
		str += f.String()
	}
	return str
}

func (s *Segment) isMSH() bool {
	var toCheck []rune
	if len(s.Value)>=3 {
		toCheck = s.Value[:3]
	} else if len(s.Fields)!=0{
		f, err := s.Field(0)
		if err != nil {
			return false
		}
		toCheck=f.Value[:3]
	} else {
		return false
	}
	if string(toCheck) == "MSH" {
		return true
	} else {
		return false
	}
}

func (s *Segment) parse(seps *Delimeters) error {
	if len(s.Value) < 3 {
		return fmt.Errorf("Invalid segment. Length %v", len(s.Value))
	}
	isMSH := s.isMSH()

	r := strings.NewReader(string(s.Value))
	i := 0
	ii := 0
	seq := 0
	for {
		ch, _, _ := r.ReadRune()
		ii++
		switch {
		case ch == eof || (ch == endMsg && seps.LFTermMsg):
			if ii > i {
				fld := Field{Value: s.Value[i : ii-1], SeqNum: seq}
				fld.parse(seps)
				s.Fields = append(s.Fields, fld)
			}
			return nil
		case isMSH && seq == 2 && ch == seps.Repetition:
			// ignore repeat separator in separator definition
		case isMSH && seq == 2 && ch == seps.Escape:
			// ignore escape separator in separator definition
		case ch == seps.Field:
			if isMSH && seq == 2 {
				// the separator list is a field in MSH seq 2
				s.forceField(s.Value[i:ii-1], seq)
			} else {
				fld := Field{Value: s.Value[i : ii-1], SeqNum: seq}
				fld.parse(seps)
				s.Fields = append(s.Fields, fld)
			}
			i = ii
			seq++
			if isMSH && seq == 1 {
				// The field separator is itself a field for MSH seq 1
				s.forceField([]rune(string(seps.Field)), seq)
				seq++
			}
		case ch == seps.Repetition:
			fld := Field{Value: s.Value[i : ii-1], SeqNum: seq}
			fld.parse(seps)
			s.Fields = append(s.Fields, fld)
			i = ii
		case ch == seps.Escape:
			ii++
			r.ReadRune()
		}
	}
}

// forceField will force the creation of a field / component / subcomponent
// This is used for separator defines in the MSH segemnt
// ...and the name forceField is cool ;)
func (s *Segment) forceField(val []rune, seq int) {
	if seq > s.maxSeq {
		s.maxSeq = seq
	}
	fld := Field{Value: val, SeqNum: seq}
	cmp := Component{Value: val}
	cmp.SubComponents = append(cmp.SubComponents, SubComponent{Value: val})
	fld.Components = append(fld.Components, cmp)
	s.Fields = append(s.Fields, fld)
}

func (s *Segment) encode(seps *Delimeters) []rune {
	buf := []string{}
	for _, f := range s.Fields {
		buf = append(buf, string(f.Value))
	}
	if s.isMSH() {
		firstFields := strings.Join(buf[0:3], "")
		otherFields := strings.Join(buf[3:], string(seps.Field))
		return []rune(strings.Join([]string{firstFields, otherFields}, string(seps.Field)))
	} else {
		return []rune(strings.Join(buf, string(seps.Field)))
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

// Set will insert a value into a message at Location
func (s *Segment) Set(l *Location, val string, seps *Delimeters) error {
	if l.FieldSeq == -1 {
		return errors.New("Field is required")
	}
	if s.maxSeq < l.FieldSeq {
		for i := s.maxSeq + 1; i <= l.FieldSeq; i++ {
			s.forceField([]rune(""), i)
		}
	}
	fld, err := s.Field(l.FieldSeq)
	if err != nil {
		return err
	}
	err = fld.Set(l, val, seps)
	if err != nil {
		return err
	}
	s.Value = s.encode(seps)
	return nil
}
