package golevel7

import "reflect"

// Decode will parse an hl7 byte slice and return a Message
// if error is not nil the message can by queried for a list errors
// occuring during parsing
func Decode(buf []byte) (*Message, error) {
	msg := NewMessage(buf)
	err := msg.parse()
	return msg, err
}

// Unmarshal fills a structure from an HL7 message
// It will panic if interface{} is not a pointer to a struct
// Unmarshal will decode the entire message before trying to set values
// it will set the first matching segment / first matching field
// repeating segments and fields is not well suited to this
// for the moment all structure fields must be strings
func Unmarshal(data []byte, it interface{}) error {
	msg, err := Decode(data)
	if err != nil {
		return err
	}

	st := reflect.ValueOf(it).Elem()
	stt := st.Type()
	for i := 0; i < st.NumField(); i++ {
		fld := stt.Field(i)
		r := fld.Tag.Get("hl7")
		if r != "" {
			if val, _ := Retrieve(msg, r); val != "" {
				if st.Field(i).CanSet() {
					// TODO support fields other than string
					//fldT := st.Field(i).Type()
					st.Field(i).SetString(val)
				}
			}
		}
	}

	return nil
}
