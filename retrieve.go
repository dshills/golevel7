package golevel7

import (
	"strconv"
	"strings"
)

// Retrieve gets a value from a message
// loc uses the format segment.field.component.subcomponent
// loc == "" returns the message
// loc == "MSH" returns the MSH segment
// loc == "MSH.2" returns the second field of the MSH segment
// etc
// if the loc is not valid an error is returned
func Retrieve(m *Message, loc string) (string, error) {
	l := strings.Split(loc, ".")
	switch len(l) {
	case 0: // message
		return string(m.Value), nil
	case 1: // segment
		seg, err := m.Segment(l[0])
		if err != nil {
			return "", err
		}
		return string(seg.Value), nil
	case 2: // field
		seg, err := m.Segment(l[0])
		if err != nil {
			return "", err
		}
		i, err := strconv.Atoi(l[1])
		if err != nil {
			return "", err
		}
		fld, err := seg.Field(i)
		if err != nil {
			return "", err
		}
		return string(fld.Value), nil
	case 3: // component
		seg, err := m.Segment(l[0])
		if err != nil {
			return "", err
		}
		i, err := strconv.Atoi(l[1])
		if err != nil {
			return "", err
		}
		fld, err := seg.Field(i)
		if err != nil {
			return "", err
		}
		i, err = strconv.Atoi(l[2])
		if err != nil {
			return "", err
		}
		com, err := fld.Component(i)
		if err != nil {
			return "", err
		}
		return string(com.Value), nil
	default: // subcomponent
		seg, err := m.Segment(l[0])
		if err != nil {
			return "", err
		}
		i, err := strconv.Atoi(l[1])
		if err != nil {
			return "", err
		}
		fld, err := seg.Field(i)
		if err != nil {
			return "", err
		}
		i, err = strconv.Atoi(l[2])
		if err != nil {
			return "", err
		}
		com, err := fld.Component(i)
		if err != nil {
			return "", err
		}
		i, err = strconv.Atoi(l[3])
		if err != nil {
			return "", err
		}
		sc, err := com.SubComponent(i)
		if err != nil {
			return "", err
		}
		return string(sc.Value), nil
	}
}
