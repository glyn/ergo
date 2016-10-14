package main

import (
	"testing"
)

// example test code using built in testing package

func TestValidGetUserId(t *testing.T) {
	userId, err := getUserId("root")

	if err != nil {
		t.Errorf("Failed to obtain UID: %s", err)
	}

	if userId != 0 {
		t.Errorf("Expected UID to be 0, but was %d", userId)
	}
}

func TestInvalidGetUserId(t *testing.T) {
	_, err := getUserId("someuserthatdoesnotexist")

	if err != nil {
		t.Errorf("Failed to obtain UID: %s", err)
	}

	if err == nil {
		t.Errorf("Expected not to find user: someuserthatdoesnotexist but was found")
	}
}
