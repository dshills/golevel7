package golevel7

import (
	"bufio"
	"bytes"
	"io"
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

const bufCap = 1024 * 100

func readBuf(reader io.Reader) ([]byte, error) {
	r := bufio.NewReader(reader)
	buf := make([]byte, 0, bufCap)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		switch {
		case err == io.EOF:
			return buf, err
		case n < bufCap:
			return buf, nil
		case err != nil:
			return nil, err
		}
	}
}

// Split will split a set of HL7 messages
//	\x0b MESSAGE \x1c\x0d
func Split(buf []byte) [][]byte {
	msgSep := []byte{'\x1c', '\x0d'}
	msgs := bytes.Split(buf, msgSep)
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

// Messages returns a new Message slice parsed from stream r
func (d *Decoder) Messages() ([]*Message, error) {
	buf, err := readBuf(d.r)
	if err != nil {
		return nil, err
	}
	bufs := Split(buf)
	z := []*Message{}
	for _, buf := range bufs {
		msg := NewMessage(buf)
		if err := msg.parse(); err != nil {
			return nil, err
		}
		z = append(z, msg)
	}
	return z, nil
}
