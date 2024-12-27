package test

import (
	"streamline-post/post"
	"streamline-post/structs"
	"testing"
)

func TestLinkedInPost(t *testing.T) {
	postMessage := structs.PostMessage{
		User: structs.User{
			LinkedInURN: "urn:li:person:123456",
			AccessToken: "test_token",
		},
		Body:       "Test LinkedIn Post",
		Visibility: "PUBLIC",
		Platform:   "linkedin",
	}

	// Mock LinkedInPost function
	err := post.LinkedInPost(postMessage)
	if err != nil {
		t.Errorf("LinkedInPost failed: %s", err)
	} else {
		t.Log("LinkedInPost executed successfully")
	}
}

func TestXPost(t *testing.T) {
	postMessage := structs.PostMessage{
		User: structs.User{
			TwitterID:   "123456",
			AccessToken: "test_token",
		},
		Body:       "Test X Post",
		Visibility: "PUBLIC",
		Platform:   "x",
	}

	// Mock XPost function
	err := post.XPost(postMessage)
	if err != nil {
		t.Errorf("XPost failed: %s", err)
	} else {
		t.Log("XPost executed successfully")
	}
}