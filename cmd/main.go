package main

import (
	"fmt"
	"os"

	"github.com/eenees/twitch-highlight-tracker/internal/config"
	"github.com/eenees/twitch-highlight-tracker/internal/irc"
	"github.com/eenees/twitch-highlight-tracker/internal/tracker"
	"github.com/eenees/twitch-highlight-tracker/internal/twitch"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	cfg := config.LoadConfig()

	accessToken, err := twitch.GetAppAccessToken(os.Getenv("TWITCH_CLIENT_ID"), os.Getenv("TWITCH_CLIENT_SECRET"))
	if err != nil {
		fmt.Println("failed to get app access token", err)
		return
	}
	cfg.AccessToken = accessToken

	client, err := irc.NewClient(cfg.Server)
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

	tracker := tracker.NewTracker(client, cfg)

	tracker.Run()
}
