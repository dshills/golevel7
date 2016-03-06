package golevel7

// ACK is struct for an ack message
type ACK struct {
	Code         string `hl7:"MSA.1"`
	OrgControlID string `hl7:"MSA.2"`
	ErrMsg       string `hl7:"MSA.3"`
}

// Acknowledge generates an ACK message based on the MsgInfo struct
// st can be nil for success or to send an AE code
func Acknowledge(mi MsgInfo, st error) *Message {
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
