package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Printf("[%s] Connecting to Twitch IRC...\n", time.Now())

	var (
		err    error
		dialer = websocket.Dialer{
			ReadBufferSize:   1024,
			WriteBufferSize:  1024,
			HandshakeTimeout: 30 * time.Second,
		}
	)

	_, _, err = dialer.Dial("wss://irc-ws.chat.twitch.tv:443", nil)
	if err != nil {
		fmt.Printf("[%s] Cannot connect to Twitch IRC.\n", time.Now())
		return
	}

	fmt.Printf("[%s] Connected to Twitch IRC!\n", time.Now())
}
