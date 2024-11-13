package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/syumai/workers"
)

func main() {
	http.HandleFunc("/fetchData", fetchDataHandler)
	log.Println("Server started on :8787")
	log.Fatal(http.ListenAndServe(":8787", nil))

	workers.Serve(nil)
}

func fetchDataHandler(w http.ResponseWriter, r *http.Request) {
	// Define the URL and authorization token
	url := "https://URL.turso.io/v2/pipeline"
	authToken := "TOKEN"

	// Create the request body as a map
	requestBody := map[string]interface{}{
		"requests": []map[string]interface{}{
			{"type": "execute", "stmt": map[string]string{"sql": "SELECT * FROM users"}},
			{"type": "close"},
		},
	}

	// Marshal the body into JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Add headers
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// Initialize the HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to make request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Write the response to the client
	if resp.StatusCode == http.StatusOK {
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	} else {
		http.Error(w, fmt.Sprintf("Request failed with status %d: %s", resp.StatusCode, body), resp.StatusCode)
	}

}
