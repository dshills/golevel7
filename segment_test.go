package golevel7

import "testing"

func TestSegParse(t *testing.T) {
	val := []rune(`PID|||12001||Jones^John^^^Mr.||19670824|M|||123 West St.^^Denver^CO^80020^USA~520 51st Street^^Denver^CO^80020^USA|||||||
`)
	seps := NewDelimeters()
	seg := &Segment{Value: val}
	seg.parse(seps)
	if len(seg.Fields) != 20 {
		t.Errorf("Expected 20 fields got %d\n", len(seg.Fields))
	}
}

func TestSegSet(t *testing.T) {
	seps := NewDelimeters()
	loc := "ZZZ.10"
	l := NewLocation(loc)
	seg := &Segment{}
	err := seg.Set(l, "TEST", seps)
	if err != nil {
		t.Error(seg)
	}
	str, err := seg.Get(l)
	if err != nil {
		t.Error(err)
	}
	if str != "TEST" {
		t.Errorf("Expected TEST got %s\n", str)
	}
}
