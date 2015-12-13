package golevel7

// ACK is struct for an ack message
type ACK struct {
	Code         string `hl7:"MSA.1"`
	OrgControlID string `hl7:"MSA.2"`
	ErrMsg       string `hl7:"MSA.3"`
}

// Acknowledge generates an ACK message based on the passed in message
// st can be nil for success or to send an AE code
func Acknowledge(msg *Message, st error) *Message {
	mi := MsgInfo{}
	err := msg.Unmarshal(&mi)
	if err != nil {
		return nil
	}
	amsg, _ := StartMessage(*NewMsgInfoAck(&mi))
	ack := ACK{}
	ack.Code = "AA"
	ack.OrgControlID = mi.ControlID
	if st != nil {
		ack.Code = "AE"
		ack.ErrMsg = st.Error()
	}

	Marshal(amsg, &ack)
	return amsg
}
