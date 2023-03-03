package mitake

import (
	"fmt"
	"github.com/go-packagist/sms/gateway"
	"golang.org/x/text/encoding/traditionalchinese"
	"net/url"
)

type Config struct {
	ApiUrl   string
	Username string
	Password string
}

// MiTake 三竹短信
// @see https://sms.mitake.com.tw/
type MiTake struct {
	*gateway.Base

	config *Config
}

var (
	_ gateway.Config  = (*Config)(nil)
	_ gateway.Gateway = (*MiTake)(nil)
)

// New a new MiTake gateway
func New(config *Config) gateway.Gateway {
	return &MiTake{
		Base:   new(gateway.Base),
		config: config,
	}
}

// Send a message
func (m *MiTake) Send(phone, message interface{}) (*gateway.Response, error) {
	// parse phone and message
	ph, msg, err := m.Parse(phone, message)
	if err != nil {
		return nil, err
	}

	// convert to big5
	smbody, err := traditionalchinese.Big5.NewEncoder().String(msg.Content)
	if err != nil {
		return nil, err
	}

	// send request
	resp, err := m.Http().R().
		// EnableTrace().
		Get(m.config.ApiUrl + "?" +
			fmt.Sprintf("username=%s&password=%s&type=now&encoding=big5&dstaddr=%s&smbody=%s",
				m.config.Username, m.config.Password, ph.GetPhoneNumber(), url.PathEscape(smbody))) // fix: use `url.PathEscape` fix to url encode about `space`

	if err != nil {
		return nil, err
	}

	// spew.Dump(resp.Request.TraceInfo(), resp.Body(), err)

	return gateway.NewResponse(resp), nil
}
