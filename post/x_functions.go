package post

import (
	"streamline-post/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
)


func XPost(post structs.PostMessage) error {
    api_url := "https://api.twitter.com/2/tweets"
    payload := map[string]string{
        "text": post.Body,
    }

    json_payload, err := json.Marshal(payload)
    if err != nil {
        return fmt.Errorf("failed to marshal payload: %w", err)
    }

    req, err := http.NewRequest("POST", api_url, bytes.NewBuffer(json_payload))
    if err != nil {
        return fmt.Errorf("failed to create post request: %w", err)
    }

    req.Header.Set("Authorization", "Bearer "+post.User.AccessToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    return nil
}
