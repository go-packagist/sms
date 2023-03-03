package gateway

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type Gateway interface {
	Send(interface{}, interface{}) (*Response, error)
}

type Config interface {
}

type Base struct {
	httpClient *resty.Client
}

func (b *Base) ParsePhone(phone interface{}) (*Phone, error) {
	switch phone.(type) {
	case string:
		return NewPhone(phone.(string), ""), nil
	case int:
		return NewPhone(strconv.Itoa(phone.(int)), ""), nil
	case *Phone:
		return phone.(*Phone), nil
	default:
		return nil, errors.New("invalid phone type")
	}
}

func (b *Base) ParseMessage(message interface{}) (*Message, error) {
	switch message.(type) {
	case string:
		return NewMessage(message.(string)), nil
	case *Message:
		return message.(*Message), nil
	case MessageFunc:
		return message.(MessageFunc)(&Message{}), nil
	default:
		return nil, errors.New("invalid message type")
	}
}

func (b *Base) Parse(phone, message interface{}) (*Phone, *Message, error) {
	ph, err := b.ParsePhone(phone)
	if err != nil {
		return nil, nil, err
	}

	msg, err := b.ParseMessage(message)
	if err != nil {
		return ph, nil, err
	}

	return ph, msg, nil
}

// Http returns a resty client
func (b *Base) Http() *resty.Client {
	if b.httpClient == nil {
		b.httpClient = resty.New()
	}

	return b.httpClient
}
