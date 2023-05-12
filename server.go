package main

import (
	"encoding/json"
	"fmt"
	"github.com/r3labs/sse/v2"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Id         int   `json:"id"`
	Percentage int64 `json:"percentage"`
}

func main() {
	server := sse.New()
	server.CreateStream("events")

	mux := http.NewServeMux()
	mux.HandleFunc("/events", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")

		if server.StreamExists(request.URL.Query().Get("stream")) {
			server.ServeHTTP(writer, request)
			return
		}

		writer.WriteHeader(http.StatusGone)
		return
	})

	go func() {
		defer server.RemoveStream("events")
		for i := 0; i <= 100; i++ {
			msg := &Message{
				Id:         i,
				Percentage: int64(i),
			}
			message, _ := json.Marshal(msg)
			server.Publish("events", &sse.Event{ID: []byte(strconv.Itoa(msg.Id)), Data: message})
			time.Sleep(time.Second)
		}
	}()

	port := 8000
	if err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), mux); err != nil {
		panic(err)
	}
}
