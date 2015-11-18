package golevel7

import "testing"

func TestCompParse(t *testing.T) {
	val := []byte("v1&v2&v3&&v5")
	seps := Separators{
		FieldSep:  '|',
		ComSep:    '^',
		RepSep:    '~',
		EscSep:    '\\',
		SubComSep: '&',
		SepField:  "^~\\&",
	}
	cmp := &Component{Value: val}
	cmp.parse(&seps)
	if len(cmp.SubComponents) != 5 {
		t.Errorf("Expected 5 subcomponents got %d\n", len(cmp.SubComponents))
	}
}
