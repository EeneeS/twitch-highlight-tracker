package irc

import (
	"fmt"
	"strings"
)

type Message struct {
  Raw string
  Text string
}

func ParseMessage(message string) *Message {
  msg := Message{Raw: message}
  if strings.Contains(msg.Raw, "PRIVMSG") {
      parts := strings.Split(message, " :")
      if len(parts) > 1 {
        msg.Text = parts[1]
        fmt.Println(msg.Text)
      }
  }
  return &msg
}
