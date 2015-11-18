package golevel7

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
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
	/*
		for i, seg := range msg.Segments {
			t.Errorf("Seg %d %s\n", i, seg.Value)
		}
	*/
}

func TestUnmarshal(t *testing.T) {
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

	type my7 struct {
		FirstName string `hl7:"PID.5.1"`
		LastName  string `hl7:"PID.5.0"`
	}
	st := my7{}

	Unmarshal(data, &st)

	if st.FirstName != "John" {
		t.Errorf("Expected John got %s\n", st.FirstName)
	}
	if st.LastName != "Jones" {
		t.Errorf("Expected Jones got %s\n", st.LastName)
	}
}
