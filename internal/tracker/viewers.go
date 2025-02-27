package tracker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type TwitchStreamsResponse struct {
	Data []struct {
		ViewerCount int `json:"viewer_count"`
	} `json:"data"`
}

func (t *Tracker) UpdateViewerCount() {
	for {
		viewerCount, err := t.fetchViewerCountFromAPI()
		if err != nil {
			fmt.Println("failed to update viewercount: ", err) // maybe fall back to the previous fetched value (if any or >0?)
			return
		}
		t.viewerLock.Lock()
		t.viewerCount = viewerCount
		t.viewerLock.Unlock()
		time.Sleep(time.Minute * 10)
	}
}

func (t *Tracker) fetchViewerCountFromAPI() (int, error) {
	url := fmt.Sprintf("https://api.twitch.tv/helix/streams?user_login=%v", t.channel)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		return 0, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", t.accessToken))
	req.Header.Set("Client-Id", os.Getenv("TWITCH_CLIENT_ID"))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("request failed with statuscode %v: %v", resp.StatusCode, string(rawBody))
	}

	var twitchStreamsResponse TwitchStreamsResponse
	err = json.Unmarshal(rawBody, &twitchStreamsResponse)

	return twitchStreamsResponse.Data[0].ViewerCount, nil
}

func (t *Tracker) GetViewerCount() int {
	t.viewerLock.RLock()
	defer t.viewerLock.RUnlock()
	return t.viewerCount
}
