package main

import (
	"bytes"
	"fmt"
	"sync"
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
		wg = sync.WaitGroup{}
	)

	conn, _, err := dialer.Dial("wss://irc-ws.chat.twitch.tv:443", nil)
	if err != nil {
		fmt.Printf("[%s] Cannot connect to Twitch IRC.\r\n", time.Now())
		return
	}

	fmt.Printf("[%s] Connected to Twitch IRC!\r\n", time.Now())

	conn.WriteMessage(1, []byte("PASS oauth:9zej7iro2o2m8ytan980m0sj9p3t4a\r\n"))
	conn.WriteMessage(1, []byte("NICK ProvidenceBot\r\n"))
	conn.WriteMessage(1, []byte("JOIN #paladinight\r\n"))

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			_, messageBytes, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("[%s] Connection to Twitch IRC lost, err: %s\r\n", time.Now(), err)
				conn.Close()
				return
			}

			fmt.Printf("%s", bytes.NewBuffer(messageBytes).String())
		}
	}()

	wg.Wait()
}
