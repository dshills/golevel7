package golevel7

import "testing"

func TestBuildMessage(t *testing.T) {

	mi := MsgInfo{
		SendingApp:        "BettrLife",
		SendingFacility:   "UnivIa",
		ReceivingApp:      "Epic",
		ReceivingFacility: "UnivIa",
		MessageType:       "ORM^001",
	}
	msg, err := StartMessage(mi)
	if err != nil {
		t.Fatal(err)
	}
	v, err := msg.Find("MSH.3")
	if err != nil {
		t.Error(err)
	}
	if v != "BettrLife" {
		t.Errorf("Expected BettrLife got %s\n", v)
	}
	v, err = msg.Find("MSH.4")
	if err != nil {
		t.Error(err)
	}
	if v != "UnivIa" {
		t.Errorf("Expected UnivIa got %s\n", v)
	}
	v, err = msg.Find("MSH.5")
	if err != nil {
		t.Error(err)
	}
	if v != "Epic" {
		t.Errorf("Expected Epic got %s\n", v)
	}
	v, err = msg.Find("MSH.6")
	if err != nil {
		t.Error(err)
	}
	if v != "UnivIa" {
		t.Errorf("Expected UnivIa got %s\n", v)
	}
	v, err = msg.Find("MSH.9")
	if err != nil {
		t.Error(err)
	}
	if v != "ORM^001" {
		t.Errorf("Expected ORM^001 got %s\n", v)
	}
	v, err = msg.Find("MSH.7")
	if err != nil {
		t.Error(err)
	}
	if v == "" {
		t.Error("Expected value got none")
	}
	v, err = msg.Find("MSH.10")
	if err != nil {
		t.Error(err)
	}
	if v == "" {
		t.Error("Expected value got none")
	}
	v, err = msg.Find("MSH.11")
	if err != nil {
		t.Error(err)
	}
	if v != "P" {
		t.Errorf("Expected P got %s\n", v)
	}
	v, err = msg.Find("MSH.12")
	if err != nil {
		t.Error(err)
	}
	if v != "2.4" {
		t.Errorf("Expected 2.4 got %s\n", v)
	}
}

type aMsg struct {
	FirstName string `hl7:"PID.5.1"`
	LastName  string `hl7:"PID.5.0"`
}

func TestMessageBuilding(t *testing.T) {
	mi := MsgInfo{
		SendingApp:        "BettrLife",
		SendingFacility:   "UnivIa",
		ReceivingApp:      "Epic",
		ReceivingFacility: "UnivIa",
		MessageType:       "ORM^001",
	}
	msg, err := StartMessage(mi)
	if err != nil {
		t.Fatal(err)
	}
	am := aMsg{FirstName: "Davin", LastName: "Hills"}
	_, err = Marshal(msg, &am)
	if err != nil {
		t.Error(err)
	}
}
