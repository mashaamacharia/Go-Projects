package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type AIResponse struct {
	Summary   string   `json:"summary"`
	KeyPoints []string `json:"key_points"`
}

func CallAnthropic(text string) (AIResponse, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")

	url := "https://api.anthropic.com/v1/messages"

	prompt := "Return ONLY valid JSON. No explanation.\n\nFormat:\n{\n  \"summary\": \"...\",\n  \"key_points\": [\"...\", \"...\"]\n}\n\nText:\n" + text

	payload := map[string]interface{}{
		"model":      "claude-sonnet-4-20250514",
		"max_tokens": 300,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return AIResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return AIResponse{}, err
	}

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return AIResponse{}, err
	}
	defer resp.Body.Close()

	var raw map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return AIResponse{}, err
	}

	// Extract text safely
	contentArr, ok := raw["content"].([]interface{})
	if !ok || len(contentArr) == 0 {
		return AIResponse{}, errors.New("invalid AI response structure")
	}

	textBlock, ok := contentArr[0].(map[string]interface{})
	if !ok {
		return AIResponse{}, errors.New("invalid content format")
	}

	textOutput, ok := textBlock["text"].(string)
	if !ok {
		return AIResponse{}, errors.New("missing text field")
	}

	// Parse JSON from AI
	var parsed AIResponse
	if err := json.Unmarshal([]byte(textOutput), &parsed); err != nil {
		return AIResponse{}, errors.New("failed to parse AI JSON")
	}

	return parsed, nil
}