# Go Level 7 [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/dshills/golevel7)

## Overview

	Go Level 7 is a decoder / encoder for HL7.
> 	Health Level-7 or HL7 refers to a set of international standards for transfer of clinical and administrative data between software applications used by various healthcare providers. These standards focus on the application layer, which is "layer 7" in the OSI model. -Wikipedia

## Features

* Decode HL7 messages
* Unmarshal into Go structs
* Simple query syntax
* Message validation
* Encode - Coming soon!

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

err := Unmarshal(data, &st)
```

### Parse / Decode

```go
data := []byte(...) // raw message
msg, err := Decode(data)
```

### Message Query

```go
data := []byte(...) // raw message
msg, err := Decode(data)

// First matching value
val, err := Retrieve(msg, "PID.5.1")

// All matching values
vals, err := RetrieveAll(msg, "PID.11.1")
```

### Message Validation

Message validation is accomplished using the IsValid function. Create a slice of Validation structs and pass them, with the message, to the IsValid function. The first return value is a pass / fail bool. The second return value returns the Validation structs that failed.

A number of validation slices are already defined and can be combined to build custom validation criteria. The NewValidMSH24() function is one example. It returns a set of validations for the MSH segment for version 2.4 of the HL7 specification.

```go
val := []Validation{
	Validation{Location: "MSH.0", VCheck: SpecificValue, Value: "MSH"},
	Validation{Location: "MSH.1", VCheck: HasValue},
	Validation{Location: "MSH.2", VCheck: HasValue},
}
data := []byte(...) // raw message
msg, err := Decode(data)
valid, failures := IsValid(msg, val)
```

## To Do

* message encoding
* ACK building

## Alternatives

* [gohl7](https://github.com/yehezkel/gohl7)

## License
Copyright 2015, 2016 Davin Hills. All rights reserved.
MIT license. License details can be found in the LICENSE file.


