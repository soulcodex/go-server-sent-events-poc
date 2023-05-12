package main

import (
	"fmt"
	"github.com/r3labs/sse/v2"
)

func main() {
	client := sse.NewClient("http://localhost:8000/events")

	_ = client.Subscribe("events", func(msg *sse.Event) {
		message := string(msg.Data)
		fmt.Println(message)
	})
}
