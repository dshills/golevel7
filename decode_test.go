package golevel7

import (
	"os"
	"testing"
)

type my7 struct {
	FirstName string `hl7:"PID.5.1"`
	LastName  string `hl7:"PID.5.0"`
}

func TestDecode(t *testing.T) {
	fname := "./testdata/msg.txt"
	file, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	st := my7{}
	err = NewDecoder(file).Decode(&st)
	if err != nil {
		t.Error(err)
	}
	if st.FirstName != "John" {
		t.Errorf("Expected John got %s\n", st.FirstName)
	}
	if st.LastName != "Jones" {
		t.Errorf("Expected Jones got %s\n", st.LastName)
	}
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

	st := my7{}

	Unmarshal(data, &st)

	if st.FirstName != "John" {
		t.Errorf("Expected John got %s\n", st.FirstName)
	}
	if st.LastName != "Jones" {
		t.Errorf("Expected Jones got %s\n", st.LastName)
	}
}
