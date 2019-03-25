package golevel7

import (
	"errors"
	"io"
	"reflect"
)

// Encoder writes hl7 messages to a stream
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new Encoder that writes to stream w
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode writes the encoding of it to the stream
// It will panic if interface{} is not a pointer to a struct
func (e *Encoder) Encode(it interface{}) error {
	msg := &Message{}
	b, err := Marshal(msg, it)
	if err != nil {
		return err
	}
	i, err := e.w.Write(b)
	if err != nil {
		return err
	}
	if i < len(b) {
		return errors.New("Failed to write all bytes")
	}
	return nil
}

// Marshal will insert values into a message
// It will panic if interface{} is not a pointer to a struct
func Marshal(m *Message, it interface{}) ([]byte, error) {
	seg := Segment{Value: []rune("MSH" + string(m.Delimeters.Field) + m.Delimeters.DelimeterField)}
	seg.parse(&m.Delimeters)
	m.Segments = append(m.Segments, seg)
	st := reflect.ValueOf(it).Elem()
	stt := st.Type()
	for i := 0; i < st.NumField(); i++ {
		fld := stt.Field(i)
		r := fld.Tag.Get("hl7")
		if r != "" {
			l := NewLocation(r)
			val := st.Field(i).String()
			if err := m.Set(l, val); err != nil {
				return nil, err
			}
		}
	}
	return []byte(string(m.Value)), nil
}
