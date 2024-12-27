package post

import (
	"streamline-post/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
)

func LinkedInPost(post structs.PostMessage) error {
    api_url := "https://api.linkedin.com/v2/ugcPosts"
    payload := map[string]interface{}{
        "author": post.User.LinkedInURN,
        "lifecycleState": "PUBLISHED",
        "specificContent": map[string]interface{}{
            "com.linkedin.ugc.ShareContent": map[string]interface{}{
                "shareCommentary": map[string]interface{}{
                    "text": post.Body,
                },
                "shareMediaCategory": "NONE",
            },
        },
        "visibility": map[string]string{
            "com.linkedin.ugc.MemberNetworkVisibility": post.Visibility,
        },
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

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    return nil
}
