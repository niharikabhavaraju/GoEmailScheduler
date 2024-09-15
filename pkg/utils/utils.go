package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// Mailgun  configuration
var (
	mailgunDomain string
	mailgunAPIKey string
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read environment variables
	mailgunDomain = os.Getenv("MAILGUN_DOMAIN")
	mailgunAPIKey = os.Getenv("MAILGUN_API_KEY")
}

func ParseBody(r *http.Request, i interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), i); err != nil {
			return
		}
	}
}

// Create a Basic Auth header value
func createBasicAuthHeader(username, password string) string {
	token := fmt.Sprintf("%s:%s", username, password)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))
	return "Basic " + encodedToken
}

// Send email using Mailgun API with Basic Auth
func SendEmail(from, to, subject, text string) error {
	// Create the API URL
	apiURL := fmt.Sprintf("https://api.mailgun.net/v3/%s/messages", mailgunDomain)

	// Create the data payload
	data := url.Values{}
	data.Set("from", from)
	data.Set("to", to)
	data.Set("subject", subject)
	data.Set("text", text)

	// Create the request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	// Set Basic Auth header
	req.Header.Set("Authorization", createBasicAuthHeader("api", mailgunAPIKey))

	// Set content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", body)
	}

	// Print the response   if required
	fmt.Println("Response:", string(body))

	return nil
}

// err := sendEmail(from, to, subject, text)
// if err != nil {
// 	log.Fatalf("Failed to send email: %v", err)
// }
