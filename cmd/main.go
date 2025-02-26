package main

import (
	"fmt"

	"github.com/eenees/twitch-highlight-tracker/internal/irc"
)

var (
  server = "irc.chat.twitch.tv:6667"
  channel = "ohnepixel"
)

func main() {
  client, err:= irc.NewClient(server)
  if err != nil {
    fmt.Println("failed to create client", err)
    return
  }

  err = client.Logon()
  if err != nil {
    fmt.Println("failed to logon", err)
    return
  }

  err = client.JoinChannel(channel)
  if err != nil {
    fmt.Println("failed to join channel", err)
    return
  }

  client.ReadMessages()
}
