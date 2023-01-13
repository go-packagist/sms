package mitake

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/go-packagist/sms/gateway"
	"github.com/go-packagist/sms/support/utils"
)

const ApiUri = "https://smsapi.mitake.com.tw"
const ApiSmsUrl = "/api/mtk/SmSend"

type Config struct {
	Username string
	Password string
}

// MiTake 三竹短信
// @see https://sms.mitake.com.tw/common/index.jsp?t=1673245880983#
type MiTake struct {
	*gateway.Base

	config *Config
}

var _ gateway.Config = (*Config)(nil)
var _ gateway.Gateway = (*MiTake)(nil)

// New a new MiTake gateway
func New(config *Config) gateway.Gateway {
	return &MiTake{
		Base:   new(gateway.Base),
		config: config,
	}
}

// Send a message
func (m *MiTake) Send(phone, message interface{}) error {
	ph, msg, err := m.Parse(phone, message)
	if err != nil {
		return err
	}

	body, err := utils.Utf8ToBig5([]byte(msg.Content))

	params := map[string]string{
		"CharsetURL": "UTF8",
		"username":   m.config.Username,
		"password":   m.config.Password,
		"dstaddr":    ph.GetPhoneNumber(),
		"smbody":     string(body),
		"type":       "now",
	}

	response, err := m.Http().R().
		EnableTrace().
		SetQueryParams(params).
		Get(ApiUri + ApiSmsUrl)
	// // response, err := m.Http().R().SetQueryParams(params).Get(ApiUri + ApiSmsUrl)
	// // responseBody, _ := utils.Big5ToUtf8(response.Body())
	spew.Dump(ApiUri+ApiSmsUrl, response.Request.TraceInfo(), response.String(), err)
	// if err != nil {
	// 	return err
	// }

	return nil
}
