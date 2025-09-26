package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TurnstileResponse struct {
	Success     bool     `json:"success"`
	ChallangeTs string   `json:"challange_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

// be gapeduli bentuk challange nya, yang penting nerima token turnstile nya
func VerifyTurnstile(token string) (bool, error) {
	secret := os.Getenv("TURNSTILE_SECRET_KEY")
	url := "https://challenges.cloudflare.com/turnstile/v0/siteverify"
	data := fmt.Sprintf("secret=%s&response=%s", secret, token)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var tsResp TurnstileResponse
	if err := json.Unmarshal(body, &tsResp); err != nil {
		return false, err
	}

	return tsResp.Success, nil
}
