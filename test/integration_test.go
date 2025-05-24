package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

const baseURL = "http://localhost:8080"

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

func TestUserRegistrationAndLogin(t *testing.T) {
	// Generate unique email for test
	email := fmt.Sprintf("test%d@example.com", time.Now().Unix())

	// Test Registration
	regReq := RegisterRequest{
		Name:     "Test User",
		Email:    email,
		Password: "password123",
	}

	regJSON, _ := json.Marshal(regReq)
	resp, err := http.Post(baseURL+"/auth/register", "application/json", bytes.NewBuffer(regJSON))
	if err != nil {
		t.Fatalf("Registration request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	// Test Login
	loginReq := LoginRequest{
		Email:    email,
		Password: "password123",
	}

	loginJSON, _ := json.Marshal(loginReq)
	resp, err = http.Post(baseURL+"/auth/login", "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		t.Fatalf("Login request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var authResp AuthResponse
	json.NewDecoder(resp.Body).Decode(&authResp)

	if authResp.Token == "" {
		t.Error("Expected token in response")
	}
}
