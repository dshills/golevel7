package golevel7

import (
	"os"
	"testing"
)

func readFile(fname string) ([]byte, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]byte, 1024)
	if _, err = file.Read(data); err != nil {
		return nil, err
	}
	return data, nil
}

func TestMessage(t *testing.T) {
	data, err := readFile("./testdata/msg.txt")
	if err != nil {
		t.Fatal(err)
	}

	msg := &Message{Value: data}
	msg.parse()
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 5 {
		t.Errorf("Expected 5 segments got %d\n", len(msg.Segments))
	}

	data, err = readFile("./testdata/msg2.txt")
	if err != nil {
		t.Fatal(err)
	}
	msg = &Message{Value: data}
	msg.parse()
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 5 {
		t.Errorf("Expected 5 segments got %d\n", len(msg.Segments))
	}

	data, err = readFile("./testdata/msg3.txt")
	if err != nil {
		t.Fatal(err)
	}
	msg = &Message{Value: data}
	msg.parse()
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 9 {
		t.Errorf("Expected 9 segments got %d\n", len(msg.Segments))
	}

	data, err = readFile("./testdata/msg4.txt")
	if err != nil {
		t.Fatal(err)
	}
	msg = &Message{Value: data}
	msg.parse()
	if err != nil {
		t.Error(err)
	}
	if len(msg.Segments) != 9 {
		t.Errorf("Expected 9 segments got %d\n", len(msg.Segments))
	}
}

func TestMsgUnmarshal(t *testing.T) {
	fname := "./testdata/msg.txt"
	file, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	msg, err := NewDecoder(file).Message()
	if err != nil {
		t.Fatal(err)
	}
	st := my7{}

	msg.Unmarshal(&st)

	if st.FirstName != "John" {
		t.Errorf("Expected John got %s\n", st.FirstName)
	}
	if st.LastName != "Jones" {
		t.Errorf("Expected Jones got %s\n", st.LastName)
	}
}
