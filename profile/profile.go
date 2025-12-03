package profile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Property struct {
	Name      string  `json:"name"`
	Value     string  `json:"value"`
	Signature *string `json:"signature,omitempty"`
}

type PlayerProfileRaw struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}

func GetUUID(username string) (string, error) {
	url := fmt.Sprintf("https://api.mojang.com/users/profiles/minecraft/%s", username)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching UUID: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	var result struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("error parsing JSON: %w", err)
	}

	return result.ID, nil
}

func FetchProfileRaw(uuid string, unsigned bool) (*PlayerProfileRaw, error) {
	url := fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/profile/%s", uuid)
	if !unsigned {
		url += "?unsigned=false"
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching profile: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var profile PlayerProfileRaw
	if err := json.Unmarshal(body, &profile); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &profile, nil
}
