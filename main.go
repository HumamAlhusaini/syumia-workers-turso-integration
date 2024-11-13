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

	workers.Serve(nil)
}

func fetchDataHandler(w http.ResponseWriter, r *http.Request) {
	// Define the URL and authorization token
	url := "https://URL"
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
		http.Error(w, "Failed to marshal JSON"+err.Error(), http.StatusInternalServerError)
	}

	// Print the JSON data
	log.Println("This is the JSON Data" + string(jsonData))
	log.Println()

	// Create a new POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Failed to create request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Add headers
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	// print the request
	log.Println("This is the request", req)
	log.Println()

	// print request header
	log.Println("This is the header", req.Header)
	log.Println()

	// Initialize the HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to make request: "+err.Error(), http.StatusInternalServerError)
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
