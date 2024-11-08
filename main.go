package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/syumai/workers"
)

func main() {
	url := "https://example.turso.io/v2/pipeline"
	authToken := "TOKEN"

	// Define the request body
	jsonData := []byte(`{"requests":[{"type":"execute", "stmt": { "sql": "SELECT 1" }}, {"type": "close"}]}`)

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Print the response
	fmt.Println("Response:", string(body))

	workers.Serve(nil)
}
