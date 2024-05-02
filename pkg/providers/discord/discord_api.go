package discord

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/projectdiscovery/gologger"
)

func (options *Options) SendThreaded(message string) error {

	payload := APIRequest{
		Content:   message,
		Username:  options.DiscordWebHookUsername,
		AvatarURL: options.DiscordWebHookAvatarURL,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", options.DiscordWebHookURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	// Send the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		gologger.Error().Msgf("Error sending request: %v", err.Error())
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		gologger.Error().Msgf("Request failed with status code: %s", strconv.Itoa(resp.StatusCode))
	}

	return nil
}
