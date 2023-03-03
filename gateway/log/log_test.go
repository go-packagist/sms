package log

import (
	"github.com/go-packagist/sms/gateway"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSend(t *testing.T) {
	l := New(&Config{})
	resp, err := l.Send(gateway.NewPhone("13312341234", "+86"), gateway.NewMessage("hello"))

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.GetStatusCode())
	assert.Equal(t, "OK", resp.GetStatus())
	assert.Equal(t, "HTTP/1.1", resp.GetProto())
	assert.Equal(t, "", resp.GetBody())
	assert.True(t, resp.IsSuccessful())
}
