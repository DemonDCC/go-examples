package main

import (
	"log"
	"runtime"

	"github.com/zhangce1999/go-examples/nats/utils"
	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect(utils.GetURL())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		go conn.Publish("demon", []byte("hello"))
	}

	runtime.Goexit()
}
