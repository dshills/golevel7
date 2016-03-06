package golevel7

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Decoder reades hl7 messages from a stream
type Decoder struct {
	r io.Reader
}

// NewDecoder returns a new Decoder that reades from from stream r
// Assumes one message per stream
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

/*
	\x0b MESSAGE \x1c\x0d
	SEG\x0d
*/

// Split will split a set of HL7 messages
func (d *Decoder) Split() [][]byte {
	msgSep := []byte{'\x1c', '\x0d'}
	by, err := ioutil.ReadAll(d.r)
	by = bytes.Trim(by, "\x00")
	if len(by) == 0 || err != nil {
		return [][]byte{}
	}
	msgs := bytes.Split(by, msgSep)
	vmsgs := [][]byte{}
	for _, msg := range msgs {
		if len(msg) < 4 {
			continue
		}
		msg = bytes.TrimLeft(msg, "\x0b")
		vmsgs = append(vmsgs, msg)
	}
	return vmsgs
}

/*
func (d *Decoder) split() [][]byte {
	by, err := ioutil.ReadAll(d.r)
	if len(by) == 0 || err != nil {
		return [][]byte{}
	}
	by = bytes.Trim(by, "\x00")
	by = bytes.TrimLeft(by, "\x0b")
	by = bytes.TrimSuffix(by, []byte{'\x1c', '\x0d'})

	return bytes.Split(by, []byte("\x0a"))
}
*/

// Messages returns a new Message slice parsed from stream r
func (d *Decoder) Messages() ([]*Message, error) {
	z := []*Message{}
	bufs := d.Split()
	for _, buf := range bufs {
		msg := NewMessage(buf)
		if err := msg.parse(); err != nil {
			return nil, err
		}
		z = append(z, msg)
	}
	return z, nil
}
