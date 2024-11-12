// main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/syumai/workers"
)

func main() {
	url := "https://my-db-humamalhusaini.turso.io/v2/pipeline"
	authToken := "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MzEyNzcwODUsImlkIjoiNzMwZDZiZDAtYzlmNS00MjFkLWE5MTQtNGNmYjZkOTlhNjU0In0.zkSvvLcXOAgvRNqJrY-Bi8bZjynZiMF3EHhKSpwIWeNWznhlR17mmSWsKNZLaCw_3wvuYOYqXSBYOJrPd248DQ"

	// Define the request payload
	payload := map[string]interface{}{
		"requests": []map[string]interface{}{
			{"type": "execute", "stmt": map[string]string{"sql": "SELECT 1"}},
			{"type": "close"},
		},
	}

	// Encode payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	workers.Serve(nil)
}
