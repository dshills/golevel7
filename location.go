package golevel7

import (
	"strconv"
	"strings"
)

/**
Location syntax

// loc uses the format segment.field.component.subcomponent
// loc == "" returns the message
// loc == "MSH" returns the MSH segment
// loc == "MSH.2" returns the second field of the MSH segment
// etc

**/

// Location specifies a value or values in an Message
type Location struct {
	Segment  string
	FieldSeq int
	Comp     int
	SubComp  int
}

// NewLocation creates a Location struct based on location string syntax
func NewLocation(l string) *Location {
	la := strings.Split(l, ".")
	loc := Location{FieldSeq: -1, Comp: -1, SubComp: -1}
	lenLA := len(la)
	if lenLA > 0 {
		loc.Segment = la[0]
	}
	if lenLA > 1 {
		if i, err := strconv.Atoi(la[1]); err == nil {
			loc.FieldSeq = i
		}
	}
	if lenLA > 2 {
		if i, err := strconv.Atoi(la[2]); err == nil {
			loc.Comp = i
		}
	}
	if lenLA > 3 {
		if i, err := strconv.Atoi(la[3]); err == nil {
			loc.SubComp = i
		}
	}

	return &loc
}

// mshOffset used just for building messages. Since the field seperator is used
// in the MSH seg 1 building messages gets confused about locations
func mshOffset(l *Location) *Location {
	if l.Segment == "MSH" && l.FieldSeq > 2 {
		l.FieldSeq--
	}
	return l
}
