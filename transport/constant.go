package transport

import (
	"time"

	"github.com/nats-io/nats.go"
)

var Nc *nats.Conn
var Domain string
var Timeout = 20 * time.Second

func Constructor(url, domain string) {
	Nc = ClientConstructor(url)
	Domain = domain
}
