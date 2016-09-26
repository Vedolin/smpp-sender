package smppsender

import (
	"errors"
	"strings"
)

var ErrPhoneFormat = errors.New("Phone number in incorrect format")
var ErrNoDefaultSender = errors.New("Default sender not set")

const defaultSender = "999"
const maxCodeLength = 3

// Sender ...
type Sender struct {
	Directs map[string]DirectSender
}

// Send message with corresponding DirectSender based on country code in phone number
func (s Sender) Send(phone, text string) error {
	// Just random value, it must be clarified
	if len(phone) < 5 {
		return ErrPhoneFormat
	}

	if strings.HasPrefix(phone, "+") {
		phone = strings.TrimPrefix(phone, "+")
	}

	for n := 1; n <= maxCodeLength; n++ {
		code := phone[:n]

		ds, ok := s.Directs[code]
		if ok {
			return ds.Send(phone, text)
		}
	}

	ds, ok := s.Directs[defaultSender]
	if !ok {
		return ErrNoDefaultSender
	}

	return ds.Send(phone, text)
}
