package golevel7

import (
	"os"
	"testing"
)

func TestValid(t *testing.T) {

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

	valid, failures := msg.IsValid(NewValidMSH24())
	if valid == false {
		t.Error("Expected valid MSH got invalid. Failures:")
		for i, f := range failures {
			t.Errorf("%d %+v\n", i, f)
		}
	}

	valid, failures = msg.IsValid(NewValidPID24())
	if valid == false {
		t.Error("Expected valid PID got invalid. Failures:")
		for i, f := range failures {
			t.Errorf("%d %+v\n", i, f)
		}
	}
}
