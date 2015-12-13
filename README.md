# Go Level 7 [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/dshills/golevel7)

## Overview

	Go Level 7 is a decoder / encoder for HL7.
> 	Health Level-7 or HL7 refers to a set of international standards for transfer of clinical and administrative data between software applications used by various healthcare providers. These standards focus on the application layer, which is "layer 7" in the OSI model. -Wikipedia

## Features

* Decode HL7 messages
* Unmarshal into Go structs
* Simple query syntax
* Message validation
* Encode messages
* Build ACK messages

## Installation
	go get github.com/dshills/golevel7

## Usage

###	Data Location Syntax

	segment-name.field-sequence-number.component.subcomponent
	Segments are specified using the three letter name (MSH)
	Fields are specifified by the sequence number of the HL7 specification (Field 0 is always the segment name)
	Components and Subcomponents are 0 based indexes
	"" returns the message
	"PID" returns the PID segment
	"PID.5" returns the 5th field of the PID segment
	"PID.5.1" returns the 1st component of the 5th field of the PID segment
	"PID.5.1.2" returns the 2nd subcomponent of the 1st component of the 5th field of the PID

###	Data Extraction / Unmarshal

```go
data := []byte(...) // raw message
type my7 struct {
	FirstName string `hl7:"PID.5.1"`
	LastName  string `hl7:"PID.5.0"`
}
st := my7{}

err := golevel7.Unmarshal(data, &st)

// from an io.Reader

golevel7.NewDecoder(reader).Decode(&st)
```

### Message Query

```go
msg, err := golevel7.NewDecoder(reader).Message()

// First matching value
val, err := msg.Find("PID.5.1")

// All matching values
vals, err := msg.FindAll("PID.11.1")
```

### Message building

```go
	type aMsg struct {
		FirstName string `hl7:"PID.5.1"`
		LastName  string `hl7:"PID.5.0"`
	}
	mi := golevel7.MsgInfo{
		SendingApp:        "MyApp",
		SendingFacility:   "MyPlace",
		ReceivingApp:      "EMR",
		ReceivingFacility: "MedicalPlace",
		MessageType:       "ORM^001",
	}
	msg, err := golevel7.StartMessage(mi)
	am := aMsg{FirstName: "Davin", LastName: "Hills"}
	bstr, err = golevel7.Marshal(msg, &am)

	// Manually

	type MyHL7Message struct {
		SendingApp        string `hl7:"MSH.3"`
		SendingFacility   string `hl7:"MSH.4"`
		ReceivingApp      string `hl7:"MSH.5"`
		ReceivingFacility string `hl7:"MSH.6"`
		MsgDate           string `hl7:"MSH.7"`
		MessageType       string `hl7:"MSH.9"`
		ControlID         string `hl7:"MSH.10"`
		ProcessingID      string `hl7:"MSH.11"`
		VersionID         string `hl7:"MSH.12"`
		FirstName         string `hl7:"PID.5.1"`
		LastName          string `hl7:"PID.5.0"`
	}

	my := MyHL7Message{
		SendingApp:        "MyApp",
		SendingFacility:   "MyPlace",
		ReceivingApp:      "EMR",
		ReceivingFacility: "MedicalPlace",
		MessageType:       "ORM^001",
		MsgDate:           "20151209154606",
		ControlID:         "MSGID1",
		ProcessingID:      "P",
		VersionID:         "2.4",
		FirstName:         "Davin",
		LastName:          "Hills",
	}
	err := golevel7.NewEncoder(writer).Encode(&my)
```

### Message Validation

Message validation is accomplished using the IsValid function. Create a slice of Validation structs and pass them, with the message, to the IsValid function. The first return value is a pass / fail bool. The second return value returns the Validation structs that failed.

A number of validation slices are already defined and can be combined to build custom validation criteria. The NewValidMSH24() function is one example. It returns a set of validations for the MSH segment for version 2.4 of the HL7 specification.

```go
val := []golevel7.Validation{
	Validation{Location: "MSH.0", VCheck: SpecificValue, Value: "MSH"},
	Validation{Location: "MSH.1", VCheck: HasValue},
	Validation{Location: "MSH.2", VCheck: HasValue},
}
msg, err := golevel7.NewDecoder(reader).Message()
valid, failures := msg.IsValid(val)
```

## To Do

* Better handling of repeating fields for marshal and unmarshal
* Better handling of repeating segments for marshal and unmarshal

## Alternatives

* [gohl7](https://github.com/yehezkel/gohl7)

## License
Copyright 2015, 2016 Davin Hills. All rights reserved.
MIT license. License details can be found in the LICENSE file.


