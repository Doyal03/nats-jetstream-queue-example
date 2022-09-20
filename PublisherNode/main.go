package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"math"
	// "os"
	"time"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL) // localhost:4222
	if err != nil {
		fmt.Println("1xx", err)
	}
	defer conn.Close()

	js, err := conn.JetStream(nats.PublishAsyncMaxPending(math.MaxInt))
	if err != nil {
		fmt.Println("2xx", err)
	}

	for i := 0; i < 100; i++ {
		_, err := js.PublishAsync("foo", []byte("Hullo People, how are youuuu.... "))
		if err != nil {
			fmt.Println("3xx", err)
			// fmt.Println(js.PublishAsyncPending())
			// os.Exit(1)
		}
	}

	select {
	case <-js.PublishAsyncComplete():
		fmt.Println("All received within timeout")
	case <-time.After((1 * time.Nanosecond)):
		fmt.Println("timeout")
	}

}
