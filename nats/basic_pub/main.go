package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nats-io/nats.go"
)

var (
	// TOPICS -
	TOPICS = os.Getenv("TOPICS")
)

func main() {
	var (
		seq = 0
	)

	topics := strings.FieldsFunc(TOPICS, func(r rune) bool {
		return r == ','
	})

	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, topic := range topics {
		seq++
		conn.Publish(topic, []byte(fmt.Sprintf("[%d:%s]: hello", seq, topic)))
	}

	conn.Flush()
}
