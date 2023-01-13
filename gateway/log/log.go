package log

import (
	"github.com/go-packagist/sms/gateway"
	_log "log"
)

type Config struct{}

type Log struct {
	*gateway.Base

	config *Config
}

var _ gateway.Config = (*Config)(nil)
var _ gateway.Gateway = (*Log)(nil)

func New(config *Config) gateway.Gateway {
	return &Log{
		config: config,
	}
}

func (l *Log) Send(phone, message interface{}) error {
	ph, msg, err := l.Parse(phone, message)
	if err != nil {
		return err
	}

	_log.Printf("%s %s",
		ph.GetFullNumber(), msg.GetContent(),
	)

	return nil
}
