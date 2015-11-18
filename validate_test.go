package golevel7

import (
	"os"
	"testing"
)

func TestValid(t *testing.T) {

	fname := "./testdata/msg.txt"
	file, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 1024)
	if _, err = file.Read(data); err != nil {
		t.Fatal(err)
	}

	msg, err := Decode(data)
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 5 {
		t.Errorf("Expected 5 segments got %d\n", len(msg.Segments))
	}

	valid, failures := IsValid(msg, NewValidMSH24())
	if valid == false {
		t.Error("Expected valid MSH got invalid. Failures:")
		for i, f := range failures {
			t.Errorf("%d %+v\n", i, f)
		}
	}

	valid, failures = IsValid(msg, NewValidPID24())
	if valid == false {
		t.Error("Expected valid PID got invalid. Failures:")
		for i, f := range failures {
			t.Errorf("%d %+v\n", i, f)
		}
	}
}
