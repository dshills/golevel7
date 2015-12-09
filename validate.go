package golevel7

// VCheck is the type validity check to be done
type VCheck int

// VCheck values
const (
	HasValue = iota
	SpecificValue
)

// Validation contains information to validate a message value
type Validation struct {
	Location string // Query syntax
	VCheck   VCheck // What to check
	Value    string // Matching value for SpecificValue
	err      error  // error to use
}

// NewValidORMDietaryOrder24 is an example of validating a ORM^001 message
func NewValidORMDietaryOrder24() []Validation {
	v := []Validation{
		Validation{Location: "MSH.9.0", VCheck: SpecificValue, Value: "ORM"},
		Validation{Location: "MSH.9.1", VCheck: SpecificValue, Value: "001"},
	}
	v = append(v, NewValidMSH24()...)
	v = append(v, NewValidPID24()...)
	v = append(v, NewValidPV124()...)
	v = append(v, NewValidORC24()...)

	return v
}

// NewValidMSH24 is the validation for the MSH segment for version 2.4
func NewValidMSH24() []Validation {
	return []Validation{
		Validation{Location: "MSH.0", VCheck: SpecificValue, Value: "MSH"},
		Validation{Location: "MSH.1", VCheck: HasValue},
		Validation{Location: "MSH.2", VCheck: HasValue},
		Validation{Location: "MSH.9", VCheck: HasValue},
		Validation{Location: "MSH.10", VCheck: HasValue},
		Validation{Location: "MSH.11", VCheck: HasValue},
		Validation{Location: "MSH.12", VCheck: HasValue},
	}
}

// NewValidPID24 is the validation for the PID segment for version 2.4
func NewValidPID24() []Validation {
	return []Validation{
		Validation{Location: "PID.0", VCheck: SpecificValue, Value: "PID"},
		Validation{Location: "PID.3", VCheck: HasValue},
		Validation{Location: "PID.5", VCheck: HasValue},
	}
}

// NewValidPV124 is the validation for the PV1 segment for version 2.4
func NewValidPV124() []Validation {
	return []Validation{
		Validation{Location: "PV1.0", VCheck: SpecificValue, Value: "PV1"},
	}
}

// NewValidORC24 is the validation for the ORC segment for version 2.4
func NewValidORC24() []Validation {
	return []Validation{
		Validation{Location: "ORC.0", VCheck: SpecificValue, Value: "ORC"},
		Validation{Location: "ORC.1", VCheck: HasValue},
	}
}

// NewValidODS24 is the validation for the ODS segment for version 2.4
func NewValidODS24() []Validation {
	return []Validation{
		Validation{Location: "ODS.0", VCheck: SpecificValue, Value: "ODS"},
		Validation{Location: "ODS.1", VCheck: HasValue},
		Validation{Location: "ODS.2", VCheck: HasValue},
	}
}
