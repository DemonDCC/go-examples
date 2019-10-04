package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/zhangce1999/go-examples/nats/utils"
	nats "github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect(utils.GetURL())
	if err != nil {
		log.Fatal(err)
	}

	sub, err := conn.Subscribe("demon", defaultHandleFunc)
	if err != nil {
		log.Println(err)
	}

	sub.AutoUnsubscribe(5)

	runtime.Goexit()
}

func defaultHandleFunc(m *nats.Msg) {
	fmt.Printf("topic: %s, payload: %s\n", m.Subject, string(m.Data))
}
