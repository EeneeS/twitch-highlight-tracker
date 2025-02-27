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

func (t *Tracker) ReadIncomming() {
  for {
    raw, err := t.client.ReadMessage()
    if err != nil {
      fmt.Println("failed to read message")
      return 
    }
    message := irc.ParseMessage(raw)
    if strings.HasPrefix(raw, "PING") {
      t.client.SendData(strings.Replace(raw, "PING", "PONG", 1))
    }
    fmt.Println(message)
  }
}

func (t *Tracker) Run() {
  t.ReadIncomming()
}
