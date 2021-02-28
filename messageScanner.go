package golevel7

import (
	"bufio"
	"golevel7/commons"
	"io"
)

type MessageScanner struct {
	r       io.Reader
	b       *bufio.Scanner
	thisMsg *Message
	err     error
}

func NewMessageScanner(r io.Reader) {
	ms := &MessageScanner{
		r: r,
		b: bufio.NewScanner(r),
	}
	ms.b.Split(commons.CrLfSplit)
}

func (ms *MessageScanner) Scan() (gotOne bool) {
	if ms.b.Scan() {
		if ms.err = ms.b.Err(); ms.err != nil {
			if ms.b.Bytes() != nil {
				gotOne = true
			}
		}
		if gotOne {
			ms.thisMsg = NewMessage(ms.b.Bytes())
		} else {
			ms.thisMsg = nil
		}
		return gotOne
	}
	return gotOne
}

func (ms *MessageScanner) Message() *Message {
	return ms.thisMsg
}

func (ms *MessageScanner) Err() error {
	return ms.err
}
