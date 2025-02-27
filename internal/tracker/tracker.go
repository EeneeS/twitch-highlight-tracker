package tracker

import (
	"fmt"
	"slices"
	"strings"

	"github.com/eenees/twitch-highlight-tracker/internal/irc"
)

type Tracker struct {
  client *irc.Client
  keywords []string
  // this should store the timestamps
}

func NewTracker(client *irc.Client, keywords []string) *Tracker {
  return &Tracker{
    client: client,
    keywords: keywords,
  }
}

func (t *Tracker) ReadIncomming() {
  for {
    raw, err := t.client.ReadMessage()
    if err != nil {
      fmt.Println("failed to read message")
      return 
    }

    if strings.HasPrefix(raw, "PING") {
      t.client.SendData(strings.Replace(raw, "PING", "PONG", 1))
    }

    if strings.Contains(raw, "PRIVMSG") {
      message := irc.ParseMessage(raw)
      if slices.Contains(t.keywords, message.Text) {
        fmt.Println(message.Text)
      }
    }
  }
}

func (t *Tracker) Run() {
  t.ReadIncomming()
}
