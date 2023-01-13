package log

import (
	"github.com/go-packagist/sms/gateway"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSend(t *testing.T) {
	l := New(&Config{})
	err := l.Send(gateway.NewPhone("13312341234", "+86"), gateway.NewMessage("hello"))

	assert.Nil(t, err)
}
