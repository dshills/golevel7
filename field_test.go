package golevel7

import "testing"

func TestFieldParse(t *testing.T) {
	val := []byte("520 51st Street^^Denver^CO^80020^USA")
	seps := Separators{
		FieldSep:  '|',
		ComSep:    '^',
		RepSep:    '~',
		EscSep:    '\\',
		SubComSep: '&',
		SepField:  "^~\\&",
	}
	fld := &Field{Value: val}
	fld.parse(&seps)
	if len(fld.Components) != 6 {
		t.Errorf("Expected 6 components got %d\n", len(fld.Components))
	}
}
