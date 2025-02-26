package main

import (
	"fmt"

	"github.com/eenees/twitch-highlight-tracker/internal/config"
	"github.com/eenees/twitch-highlight-tracker/internal/irc"
)

func main() {

  cfg := config.LoadConfig()

  client, err:= irc.NewClient(cfg.Server)
  if err != nil {
    fmt.Println("failed to create client", err)
    return
  }
  defer client.Close()


  err = client.Logon()
  if err != nil {
    fmt.Println("failed to logon", err)
    return
  }

  err = client.JoinChannel(cfg.Channel)
  if err != nil {
    fmt.Println("failed to join channel", err)
    return
  }

  client.ReadMessages()
}
