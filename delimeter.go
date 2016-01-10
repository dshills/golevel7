package golevel7

const eof = rune(0)
const endMsg = '\x0A'
const segTerm = '\x0D'

// Delimeters holds the list of hl7 message delimeters
type Delimeters struct {
	DelimeterField string
	Field          rune
	Component      rune
	Repetition     rune
	Escape         rune
	SubComponent   rune
	Truncate       rune
	LFTermMsg      bool
}

// NewDelimeters returns the default set of HL7 delimeters
func NewDelimeters() *Delimeters {
	return &Delimeters{
		DelimeterField: "^~\\&",
		Field:          '|',
		Component:      '^',
		Repetition:     '~',
		Escape:         '\\',
		SubComponent:   '&',
		Truncate:       '#',
	}
}
