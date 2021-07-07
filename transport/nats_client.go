package transport

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// ClientConstructor - ClientConstructor
func ClientConstructor(url string) *nats.Conn {
	// Initial
	connectionID := uuid.New().String()
	opts := []nats.Option{nats.Name("NATS_ExtensionLib_Requestor_" + connectionID)}
	opts = SetupConnOptions(opts)

	// Connect
	nc, err := nats.Connect(url, opts...)
	if err != nil {
		log.Println("CONNECT TO NATS SERVER FAILED:", err)
		os.Exit(-1)
	}

	return nc
}

// ClientDestructor - ClientDestructor
func ClientDestructor(nc *nats.Conn) {
	nc.Close()
}

// SetupConnOptions - SetupConnOptions
func SetupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 24 * 60 * 60
	reconnectDelay := 2 * time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(totalWait))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		fmt.Println("Got disconnected!\nError Detail:", err)
		fmt.Println("Reconnect attempt in", totalWait, "seconds")
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		fmt.Println("Got reconnected to", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		fmt.Println("Connection closed. Reason:", nc.LastError())
	}))
	return opts
}
