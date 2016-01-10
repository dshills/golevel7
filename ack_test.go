package golevel7

import (
	"errors"
	"os"
	"testing"
)

func TestAcknowledge(t *testing.T) {
	fname := "./testdata/msg.hl7"
	file, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	msg, err := NewDecoder(file).Message()
	if err != nil {
		t.Error(err)
	}
	ack := Acknowledge(msg, nil)
	if ack == nil {
		t.Fatal("Expected ACK message got nil")
	}
	ack = Acknowledge(msg, errors.New("This is a test error"))
	if ack == nil {
		t.Fatal("Expected ACK message got nil")
	}
}
