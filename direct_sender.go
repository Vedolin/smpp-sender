package smppsender

import smpp "github.com/CodeMonkeyKevin/smpp34"

const timeout = 5

// DirectSender ...
type DirectSender struct {
	Addr     string
	Port     int
	Login    string
	Password string
	Params   map[string]interface{}
}

// Send ...
func (ds DirectSender) Send(phone, text string) error {
	if ds.Params == nil {
		ds.Params = map[string]interface{}{}
	}
	ds.Params["system_id"] = ds.Login
	ds.Params["password"] = ds.Password

	tx, err := smpp.NewTransmitter(
		ds.Addr,
		ds.Port,
		timeout,
		ds.Params,
	)
	if err != nil {
		return err
	}

	// Send SubmitSm
	if _, err := tx.SubmitSm("", phone, text, &smpp.Params{}); err != nil {
		return err
	}

	return nil
}
