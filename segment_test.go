package golevel7

import "testing"

func TestSegParse(t *testing.T) {
	val := []byte("PID|||12001||Jones^John^^^Mr.||19670824|M|||123 West St.^^Denver^CO^80020^USA~520 51st Street^^Denver^CO^80020^USA|||||||")
	seps := Separators{
		FieldSep:  '|',
		ComSep:    '^',
		RepSep:    '~',
		EscSep:    '\\',
		SubComSep: '&',
		SepField:  "^~\\&",
	}
	seg := &Segment{Value: val}
	seg.parse(&seps)
	if len(seg.Fields) != 20 {
		t.Errorf("Expected 20 fields got %d\n", len(seg.Fields))
	}
}
