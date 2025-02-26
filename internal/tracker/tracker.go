package tracker

import (
	"fmt"
	"strings"

	"github.com/eenees/twitch-highlight-tracker/internal/irc"
)

type Tracker struct {
  client *irc.Client
  // this should store the timestamps
}

func NewTracker(client *irc.Client) *Tracker {
  return &Tracker{
    client: client,
  }
}

func (t *Tracker) Run() {
  for {
    rawMessage, err := t.client.ReadMessage()
    if err != nil {
      fmt.Println("failed to read message", err)
      return
    }

    message := irc.ParseMessage(rawMessage)

    if strings.HasPrefix(rawMessage, "PING") {
      t.client.SendData(strings.Replace(rawMessage, "PING", "PONG", 1))
    }

    fmt.Println(message)
  }
}
