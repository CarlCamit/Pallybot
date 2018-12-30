package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Printf("[%s] Connecting to Twitch IRC...\n", time.Now())

	var (
		wg     = sync.WaitGroup{}
		dialer = websocket.Dialer{
			ReadBufferSize:   1024,
			WriteBufferSize:  1024,
			HandshakeTimeout: 30 * time.Second,
		}
	)

	conn, _, err := dialer.Dial("wss://irc-ws.chat.twitch.tv:443", nil)
	if err != nil {
		fmt.Printf("[%s] Cannot connect to Twitch IRC.\r\n", time.Now())
		return
	}

	fmt.Printf("[%s] Connected to Twitch IRC!\r\n", time.Now())

	conn.WriteMessage(1, []byte("PASS <password>\r\n"))
	conn.WriteMessage(1, []byte("NICK <name>\r\n"))
	conn.WriteMessage(1, []byte("JOIN <channel>\r\n"))

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

			fmt.Printf("%s", messageBytes)

			rawMessage := bytes.NewBuffer(messageBytes).String()
			// Check if message is a heartbeat
			if rawMessage == "PING :tmi.twitch.tv\r\n" {
				chatResponse := "PONG :tmi.twitch.tv\r\n"
				fmt.Printf("%s", chatResponse)
				conn.WriteMessage(1, []byte(chatResponse))
				// Move on to next message after responding
				continue
			}

			rawMessageSlice := strings.Split(rawMessage, " ")
			// Check if the message is a chat message
			if rawMessageSlice[1] != "PRIVMSG" {
				// Move on to next message if not a chat message
				continue
			}

			chatMessageSlice := strings.Split(rawMessageSlice[3], " ")
			// Check if message is a command
			if strings.HasPrefix(chatMessageSlice[0], ":!") {
				command := strings.TrimPrefix(chatMessageSlice[0], ":!")
				fmt.Printf("%s", command)
			}
		}
	}()

	wg.Wait()
}
