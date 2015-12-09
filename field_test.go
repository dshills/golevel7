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

func TestFieldSet(t *testing.T) {
	seps := Separators{
		FieldSep:  '|',
		ComSep:    '^',
		RepSep:    '~',
		EscSep:    '\\',
		SubComSep: '&',
		SepField:  "^~\\&",
	}
	fld := &Field{}
	loc := "ZZZ.1.10"
	l := NewLocation(loc)
	err := fld.Set(l, "TEST", &seps)
	if err != nil {
		t.Error(err)
	}
	if len(fld.Components) != 11 {
		t.Fatalf("Expected 11 got %d\n", len(fld.Components))
	}
	if string(fld.Components[10].SubComponents[0].Value) != "TEST" {
		t.Errorf("Expected TEST got %s\n", fld.Components[10].SubComponents[0].Value)
	}
}
