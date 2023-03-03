# sms

## Installation

```bash
go get github.com/go-packagist/sms
```

## Usage

```go
package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/go-packagist/sms"
	"github.com/go-packagist/sms/gateway"
	"github.com/go-packagist/sms/gateway/log"
	"github.com/go-packagist/sms/gateway/mitake"
)

func main() {
	Example1()
	Example2()
}

func Example1() {
	cfg := &sms.Config{
		Default: "log",
		Gateways: map[string]gateway.Config{
			"mitake": &mitake.Config{
				ApiUrl: "http://sms.mitake.com.tw/path",
				Username: "test",
				Password: "test",
			},
			"log": &log.Config{},
		},
	}

	s := sms.New(cfg)
	resp1, _ := s.Gateway("mitake").Send("13312341234", "Hello World")
	spew.Dump(resp1.IsSuccessful())

	resp2, _ := s.Gateway("log").Send("13312341234", "Hello World")
	spew.Dump(resp2.IsSuccessful())
}

func Example2() {
	l := log.New(&log.Config{})
	l.Send("13312341234", "hello")
	l.Send(gateway.NewPhone("13312341234", ""), gateway.NewMessage("hello "))

	m := mitake.New(&mitake.Config{
		ApiUrl: "http://sms.mitake.com.tw/path",
		Username: "test",
		Password: "test",
	})
	m.Send("13312341234", func(message *gateway.Message) *gateway.Message {
		return message.SetContent("hello")
	})
}
```

## License

MIT