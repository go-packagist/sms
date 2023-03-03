package log

import (
	"github.com/go-packagist/sms/gateway"
	"github.com/go-resty/resty/v2"
	_log "log"
	"net/http"
)

type Config struct{}

type Log struct {
	*gateway.Base

	config *Config
}

var (
	_ gateway.Config  = (*Config)(nil)
	_ gateway.Gateway = (*Log)(nil)
)

func New(config *Config) gateway.Gateway {
	return &Log{
		config: config,
	}
}

func (l *Log) Send(phone, message interface{}) (*gateway.Response, error) {
	ph, msg, err := l.Parse(phone, message)
	if err != nil {
		return nil, err
	}

	_log.Printf("%s %s", ph.GetFullNumber(), msg.GetContent())

	return gateway.NewResponse(&resty.Response{
		Request: &resty.Request{},
		RawResponse: &http.Response{
			StatusCode: http.StatusOK,
			Status:     http.StatusText(http.StatusOK),
			Proto:      "HTTP/1.1",
			Body:       http.NoBody,
		},
	}), nil
}
