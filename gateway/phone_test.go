package gateway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhone(t *testing.T) {
	p1 := NewPhone("13312341234", "+86")
	assert.Equal(t, "13312341234", p1.GetPhoneNumber())
	assert.Equal(t, "86", p1.GetIddCode())
	assert.Equal(t, "+08613312341234", p1.GetFullNumber())

	p2 := NewPhone("13312341234", "86")
	assert.Equal(t, "13312341234", p2.GetPhoneNumber())
	assert.Equal(t, "86", p2.GetIddCode())
	assert.Equal(t, "+08613312341234", p2.GetFullNumber())

	p3 := NewPhone("13312341234", "0086")
	assert.Equal(t, "13312341234", p3.GetPhoneNumber())
	assert.Equal(t, "86", p3.GetIddCode())
	assert.Equal(t, "+08613312341234", p3.GetFullNumber())
}
