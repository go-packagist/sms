package gateway

import "strings"

type Phone struct {
	IddCode     string
	PhoneNumber string
}

func NewPhone(phoneNumber, iddcode string) *Phone {
	return &Phone{
		IddCode:     strings.TrimLeft(iddcode, "+0"),
		PhoneNumber: phoneNumber,
	}
}

func (p *Phone) GetIddCode() string {
	return p.IddCode
}

func (p *Phone) GetPhoneNumber() string {
	return p.PhoneNumber
}

func (p *Phone) GetFullNumber() string {
	if p.IddCode == "" {
		return p.PhoneNumber
	}

	return "+0" + p.IddCode + p.PhoneNumber
}
