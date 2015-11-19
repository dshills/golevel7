package golevel7

import (
	"bytes"
	"fmt"
)

const segmentSep = '\x0d'
const segmentSep2 = '\x0a'
const eof = rune(0)
const truncSep = '#'

// Separators holds the list of variable hl7 separators for a message
type Separators struct {
	FieldSep  rune
	ComSep    rune
	RepSep    rune
	EscSep    rune
	SubComSep rune
	SepField  string
}

// Message is an HL7 message
type Message struct {
	Segments   []Segment
	Value      []byte
	Separators Separators
}

// NewMessage returns a new message with the v byte value
func NewMessage(v []byte) *Message {
	return &Message{Value: v}
}

func (m *Message) String() string {
	var str string
	for _, s := range m.Segments {
		str += "Message Segment: " + string(s.Value) + "\n"
		str += s.String()
	}
	return str
}

// Segment returns the first matching segmane with name s
func (m *Message) Segment(s string) (*Segment, error) {
	for i, seg := range m.Segments {
		fld, err := seg.Field(0)
		if err != nil {
			continue
		}
		if string(fld.Value) == s {
			return &m.Segments[i], nil
		}
	}
	return nil, fmt.Errorf("Segment not found")
}

// AllSegments returns the first matching segmane with name s
func (m *Message) AllSegments(s string) ([]*Segment, error) {
	segs := []*Segment{}
	for i, seg := range m.Segments {
		fld, err := seg.Field(0)
		if err != nil {
			continue
		}
		if string(fld.Value) == s {
			segs = append(segs, &m.Segments[i])
		}
	}
	if len(segs) == 0 {
		return segs, fmt.Errorf("Segment not found")
	}
	return segs, nil
}

// Get returns the first value specified by the Location
func (m *Message) Get(l *Location) (string, error) {
	if l.Segment == "" {
		return string(m.Value), nil
	}
	seg, err := m.Segment(l.Segment)
	if err != nil {
		return "", err
	}
	return seg.Get(l)
}

// GetAll returns all values specified by the Location
func (m *Message) GetAll(l *Location) ([]string, error) {
	vals := []string{}
	if l.Segment == "" {
		vals = append(vals, string(m.Value))
		return vals, nil
	}
	segs, err := m.AllSegments(l.Segment)
	if err != nil {
		return vals, err
	}
	for _, s := range segs {
		vs, err := s.GetAll(l)
		if err != nil {
			return vals, err
		}
		vals = append(vals, vs...)
	}
	return vals, nil
}

func (m *Message) parse() error {
	if err := m.parseSep(); err != nil {
		return err
	}
	r := bytes.NewReader(m.Value)
	i := 0
	ii := 0
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			ch = eof
		}
		ii++
		switch {
		case ch == segmentSep || ch == segmentSep2:
			seg := Segment{Value: m.Value[i:ii]}
			seg.parse(&m.Separators)
			m.Segments = append(m.Segments, seg)
			i = ii
			ch, _, _ := r.ReadRune()
			if ch == segmentSep || ch == segmentSep2 {
				i++
				ii++
			} else {
				r.UnreadRune()
			}
		case ch == m.Separators.EscSep:
			ii++
			r.ReadRune()
		case ch == eof:
			v := m.Value[i:ii]
			if len(v) > 4 { // seg name + field sep
				seg := Segment{Value: v}
				seg.parse(&m.Separators)
				m.Segments = append(m.Segments, seg)
			}
			return nil
		}
	}
}

func (m *Message) parseSep() error {
	if string(m.Value[:3]) != "MSH" {
		return fmt.Errorf("Invalid message: Missing MSH segment")
	}

	r := bytes.NewReader(m.Value)
	for i := 0; i < 8; i++ {
		ch, _, _ := r.ReadRune()
		if ch == eof {
			return fmt.Errorf("Invalid message: eof while parsing MSH")
		}
		switch i {
		case 3:
			m.Separators.FieldSep = ch
		case 4:
			m.Separators.SepField = string(ch)
			m.Separators.ComSep = ch
		case 5:
			m.Separators.SepField += string(ch)
			m.Separators.RepSep = ch
		case 6:
			m.Separators.SepField += string(ch)
			m.Separators.EscSep = ch
		case 7:
			m.Separators.SepField += string(ch)
			m.Separators.SubComSep = ch
		}
	}
	return nil
}
