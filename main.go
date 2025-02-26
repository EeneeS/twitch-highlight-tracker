package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var (
  server = "irc.chat.twitch.tv:6667"
  channel = "theprimeagen"
)

func main() {
  conn, err := net.Dial("tcp", server)
  if err != nil {
    fmt.Println("failed connection", err)
    return
  }
  defer conn.Close()

  logon(conn)
  joinChannel(conn, channel)

  readLoop(conn)

}

func joinChannel(conn net.Conn, channel string) {
	sendData(conn, fmt.Sprintf("JOIN #%s", channel))
}

func logon(conn net.Conn) {
  sendData(conn, "PASS oath:justinfan1234")
  sendData(conn, "NICK justinfan1234")
}

func sendData(conn net.Conn, message string) {
  fmt.Fprintf(conn, "%s\r\n", message)
}

func readLoop(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}
		fmt.Print(message)

		// Respond to PING messages to stay connected
		if strings.HasPrefix(message, "PING") {
			sendData(conn, strings.Replace(message, "PING", "PONG", 1))
		}
	}
}
