package golevel7

import (
	"fmt"
	"time"
)

// MsgInfo describes the basic message fields
type MsgInfo struct {
	SendingApp        string `hl7:"MSH.3"`
	SendingFacility   string `hl7:"MSH.4"`
	ReceivingApp      string `hl7:"MSH.5"`
	ReceivingFacility string `hl7:"MSH.6"`
	MsgDate           string `hl7:"MSH.7"`  // if blank will generate
	MessageType       string `hl7:"MSH.9"`  // Required example ORM^001
	ControlID         string `hl7:"MSH.10"` // if blank will generate
	ProcessingID      string `hl7:"MSH.11"` // default P
	VersionID         string `hl7:"MSH.12"` // default 2.4
}

// StartMessage returns a Message with an MSH segment based on the MsgInfo struct
func StartMessage(info MsgInfo) (*Message, error) {
	if info.MessageType == "" {
		return nil, fmt.Errorf("Message Type is required")
	}
	now := time.Now()
	t := now.Format("20060102150405")
	if info.MsgDate == "" {
		info.MsgDate = t
	}
	if info.ControlID == "" {
		info.ControlID = fmt.Sprintf("MSGID%s%d", t, now.Nanosecond())
	}
	if info.ProcessingID == "" {
		info.ProcessingID = "P"
	}
	if info.VersionID == "" {
		info.VersionID = "2.4"
	}
	msg := NewMessage([]byte{})
	Marshal(msg, &info)
	return msg, nil
}
