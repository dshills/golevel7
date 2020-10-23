package golevel7

import "testing"

func TestFind(t *testing.T) {

	data, err := readFile("./testdata/msg3.hl7")
	if err != nil {
		t.Fatal(err)
	}
	msg := &Message{Value: []rune(string(data))}
	msg.parse()
	if err != nil {
		t.Fatal(err)
	}

	val, err := msg.Find("OBR.4.3")
	if err != nil {
		t.Error(err)
	}
	if val != "CPT-4" {
		t.Errorf("Expected CPT-4 got %s\n", val)
	}

	val, err = msg.Find("OBX.3.3")
	if err != nil {
		t.Error(err)
	}
	if val != "LOINC" {
		t.Errorf("Expected LOINC got %s\n", val)
	}
}

func TestFindAll(t *testing.T) {

	data, err := readFile("./testdata/msg3.hl7")
	if err != nil {
		t.Fatal(err)
	}
	msg := &Message{Value: []rune(string(data))}
	msg.parse()
	if err != nil {
		t.Fatal(err)
	}

	vals, err := msg.FindAll("OBX.1")
	if err != nil {
		t.Error(err)
	}

	if len(vals) != 4 {
		t.Fatalf("Expected 4 got %d\n", len(vals))
	}

	if vals[0] != "1" {
		t.Errorf("Expected 1 got %s\n", vals[0])
	}
	if vals[1] != "2" {
		t.Errorf("Expected 2 got %s\n", vals[1])
	}
	if vals[2] != "3" {
		t.Errorf("Expected 3 got %s\n", vals[2])
	}
	if vals[3] != "4" {
		t.Errorf("Expected 4 got %s\n", vals[3])
	}
}

func TestRepFields(t *testing.T) {
	data, err := readFile("./testdata/msg.hl7")
	if err != nil {
		t.Fatal(err)
	}
	msg := &Message{Value: []rune(string(data))}
	msg.parse()
	if err != nil {
		t.Fatal(err)
	}

	vals, err := msg.FindAll("PID.11.3")
	if err != nil {
		t.Error(err)
	}
	if len(vals) != 2 {
		t.Errorf("Expected 2 got %d\n", len(vals))
	}
}
