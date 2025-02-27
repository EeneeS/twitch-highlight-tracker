package irc

import (
	"strings"
)

type Message struct {
  Raw string
  Text string
}

func ParseMessage(message string) *Message {
  msg := Message{Raw: message}
  parts := strings.Split(message, " :")
  if len(parts) > 1 {
    msg.Text = parts[1]
  }
  return &msg
}
