package sms

import (
	"fmt"
	"github.com/go-packagist/sms/gateway"
	"github.com/go-packagist/sms/gateway/log"
	"github.com/go-packagist/sms/gateway/mitake"
	"sync"
)

type Sms struct {
	config *Config
	mu     *sync.Mutex

	gateway map[string]gateway.Gateway // aleady resolved gateway
}

func New(cfg *Config) *Sms {
	return &Sms{
		config:  cfg,
		mu:      &sync.Mutex{},
		gateway: make(map[string]gateway.Gateway, 5),
	}
}

// Send a sms.
// Example: example/sms/sms.go
func (s *Sms) Send(phone, message interface{}) error {
	return s.Gateway().Send(phone, message)
}

// Gateway choose a gateway to send sms.
func (s *Sms) Gateway(names ...string) gateway.Gateway {
	name := s.getName(names...)

	if g, ok := s.gateway[name]; ok {
		return g
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.gateway[name] = s.resolveGateway(name)

	return s.gateway[name]
}

// getName get gateway name from mutile names.
func (s *Sms) getName(names ...string) string {
	if len(names) == 0 {
		return s.config.Default
	}

	return names[0]
}

// resolveGateway resolve a gateway by name.
func (s *Sms) resolveGateway(name string) gateway.Gateway {
	cfg, ok := s.config.Gateways[name]
	if !ok {
		panic(fmt.Sprintf("Gateway [%s] not found", name))
	}

	switch cfg.(type) {
	case *mitake.Config:
		return mitake.New(cfg.(*mitake.Config))
	case *log.Config:
		return log.New(cfg.(*log.Config))
	default:
		panic(fmt.Sprintf("Gateway [%s] not found", name))
	}
}
