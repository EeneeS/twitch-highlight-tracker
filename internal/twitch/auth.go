package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TwitchResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func GetAppAccessToken(clientId string, clientSecret string) (string, error) {
	url := "https://id.twitch.tv/oauth2/token"

	data := struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
	}{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		GrantType:    "client_credentials",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error marshaling: ", err)
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute the request: %v", err)
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with statuscode %v: %v", resp.StatusCode, string(rawBody))
	}

	var twitchResponse TwitchResponse
	err = json.Unmarshal(rawBody, &twitchResponse)

	return twitchResponse.AccessToken, nil
}
