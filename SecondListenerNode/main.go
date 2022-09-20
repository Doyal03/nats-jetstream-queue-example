package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL) // localhost:4222
	if err != nil {
		panic(err)
	}
	count := 1

	conn.QueueSubscribe("foo", "hellos", func(msg *nats.Msg) {
		fmt.Printf("[%v] %v\n", count, string(msg.Data))
		count++
		conn.Publish(msg.Reply, []byte("Message delivered"))
	})

	runtime.Goexit()
	defer conn.Close()
}
