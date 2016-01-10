package golevel7

import "io"

// Decoder reades hl7 messages from a stream
type Decoder struct {
	r io.Reader
}

// NewDecoder returns a new Decoder that reades from from stream r
// Assumes one message per stream
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Message returns a new Message struct parsed from stream r
func (d *Decoder) Message() (*Message, error) {
	p := make([]byte, 1000000)
	i, err := d.r.Read(p)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, i+1)
	copy(buf, p)
	msg := NewMessage(buf)
	if err = msg.parse(); err != nil {
		return nil, err
	}
	return msg, nil
}

// Decode reads from r into interface it
func (d *Decoder) Decode(it interface{}) error {
	msg, err := d.Message()
	if err != nil {
		return err
	}
	return msg.Unmarshal(it)
}

// Unmarshal fills a structure from an HL7 message
// It will panic if interface{} is not a pointer to a struct
// Unmarshal will decode the entire message before trying to set values
// it will set the first matching segment / first matching field
// repeating segments and fields is not well suited to this
// for the moment all unmarshal target fields must be strings
func Unmarshal(data []byte, it interface{}) error {
	msg := NewMessage(data)
	err := msg.parse()
	if err != nil {
		return err
	}
	return msg.Unmarshal(it)
}
