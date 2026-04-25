package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func CallAnthropic(text string) (map[string]interface{}, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")

	url := "https://api.anthropic.com/v1/messages"

	payload := map[string]interface{}{
		"model": "claude-3-5-sonnet-latest",
		"max_tokens": 300,
		"messages": []map[string]string{
			{
				"role": "user",
				"content": "Summarize this text and give key points:\n" + text,
			},
		},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))

	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}