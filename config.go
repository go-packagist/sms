package sms

import (
	"github.com/go-packagist/sms/gateway"
)

type Config struct {
	Default  string
	Gateways map[string]gateway.Config
}
