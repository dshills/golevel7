package golevel7

import (
	"bufio"
	"bytes"
	"io"
	//"log"
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
	superBuf := []byte{}
	var err error
	for {
		buf := make([]byte, 0, bufCap)
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		quit:= false
		switch {
		case err == io.EOF:
			//log.Printf("->err == io.EOF")
			superBuf = append(superBuf,buf...)
			quit = true
		case n < bufCap:
			//log.Printf("->n < bufCap")
			superBuf = append(superBuf,buf...)
			//quit = true
		case err != nil:
			//log.Printf("->err != nil")
			superBuf = nil
		default:
			//log.Printf("->default")
			superBuf = append(superBuf,buf...)
		}
		if quit {
			break
		}
	}
	return superBuf, err
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
