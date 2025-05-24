package domain

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	// Test creating a new user
	user := NewUser("John Doe", "john@example.com", "password123")

	// Check fields were set correctly
	if user.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got %s", user.Name)
	}

	if user.Email != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got %s", user.Email)
	}

	// Check that timestamps were set
	if user.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}

	if user.UpdatedAt.IsZero() {
		t.Error("UpdatedAt should not be zero")
	}
}

func TestUpdateTimestamp(t *testing.T) {
	user := NewUser("Jane", "jane@test.com", "secret")

	// Wait a tiny bit and update timestamp
	time.Sleep(time.Millisecond)
	user.UpdateTimestamp()

	// UpdatedAt should be newer than CreatedAt
	if !user.UpdatedAt.After(user.CreatedAt) {
		t.Error("UpdatedAt should be after CreatedAt")
	}
}
