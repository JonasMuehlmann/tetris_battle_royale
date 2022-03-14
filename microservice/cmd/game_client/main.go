package main

import (
	"log"
	common "microservice/internal"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	*log.Default() = *common.NewDefaultLogger()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: ":8080", Path: "/ws"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return

			case <-interrupt:
				log.Println("interrupt")

				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)

					return
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}

				return
			}
		}
	}()

	uuid := uuid.NewString()

	log.Println(uuid)

	err = c.WriteMessage(websocket.TextMessage, []byte(`{"userID": "`+uuid+`"}`))
	if err != nil {
		log.Println("write:", err)

		return
	}

	defer close(done)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)

			return
		}

		log.Printf("recv: %s", message)
	}
}
