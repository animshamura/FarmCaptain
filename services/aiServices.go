package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"farmcaptain/config"
)

// GetAIAdvice calls the AI service to get advice for a specific crop
func GetAIAdvice(cropID string) (string, error) {
	// Prepare request to the AI service
	url := config.GetEnv("AI_URL", "http://localhost:8000/getAIAdvice")
	requestData := map[string]string{"crop_id": cropID}
	jsonData, _ := json.Marshal(requestData)

	// Send the request to the AI service
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error contacting AI service: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding AI response: %v", err)
	}

	advice := result["advice"].(string)
	return advice, nil
}
