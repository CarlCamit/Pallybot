package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const pongMessage = "PONG :tmi.twitch.tv\r\n"

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
				conn.WriteMessage(1, []byte(pongMessage))
				fmt.Printf("%s", pongMessage)
				// Move on to next message after responding
				continue
			}

			// Check if the message is a chat message
			if strings.Contains(rawMessage, "PRIVMSG") == false {
				// Move on to next message if not a chat message
				continue
			}

			rawMessageSlice := strings.Split(strings.TrimSpace(rawMessage), " ")
			chatMessageSlice := strings.Split(rawMessageSlice[3], " ")
			// Check if message is a command
			if strings.HasPrefix(chatMessageSlice[0], ":!") {
				switch chatMessageSlice[0] {
				case ":!twitter":
					conn.WriteMessage(1, []byte("PRIVMSG #bewitchedpixels :/me Follow Pix on Twitter to get updates for when she goes live! twitter.com/BewitchedPixels\r\n"))
					continue
				case ":!tumblr":
					conn.WriteMessage(1, []byte("PRIVMSG #bewitchedpixels :/me Check out Pix's Tumblr here: bewitchedpixels.tumblr.com\r\n"))
					continue
				case ":!insta":
					conn.WriteMessage(1, []byte("PRIVMSG #bewitchedpixels :/me Check out Pix's Instagram here: instagram.com/bewitchedpixels\r\n"))
					continue
				case ":!social":
					conn.WriteMessage(1, []byte("PRIVMSG #bewitchedpixels :/me ALL THE THINGS: instagram.com/bewitchedpixels || bewitchedpixels.tumblr.com || twitter.com/BewitchedPixels\r\n"))
					continue
				case ":!donate":
					conn.WriteMessage(1, []byte("PRIVMSG #bewitchedpixels :/me Donations very much appreciated, but are of course never required, and are non-refundable https://twitch.streamlabs.com/bewitchedpixels\r\n"))
					continue
				case ":!animals":
					conn.WriteMessage(1, []byte("PRIVMSG #bewitchedpixels :/me Pix has 2 dogs. Eduardo is a 10 year old Black long-haired Chi/Pom. Pixel is a 11 year old Tawny short-haired Chihuahua.\r\n"))
					continue
				case ":!discord":
					conn.WriteMessage(1, []byte("PRIVMSG #bewitchedpixels :/me Come join our cult of awesome witches and wizards: https://discord.gg/aFtpg7R\r\n"))
					continue
					// case ":!blind":
					// 	conn.WriteMessage(1, []byte("PRIVMSG #bewitchedpixels :/me Hey guys! This is my first playthrough so please avoid providing spoilers which includes adding emotional context to an upcoming part, and no tips or tricks. If I need any help I will let you guys know, thank you and enjoy!!!\r\n"))
					// 	continue
				}
			}
		}
	}()

	wg.Wait()
}
