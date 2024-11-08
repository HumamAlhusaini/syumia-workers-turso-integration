package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Define the Stmt struct
type Stmt struct {
	SQL string `json:"sql,omitempty"`
}

// Define the Request struct
type Request struct {
	Type string `json:"type"`
	Stmt *Stmt  `json:"stmt,omitempty"`
}

// Define the PipelineRequest struct
type PipelineRequest struct {
	Requests []Request `json:"requests"`
}

func main() {
	e := echo.New()

	e.POST("/sendRequest", func(c echo.Context) error {
		url := "https://example.turso.io/v2/pipeline"
		authToken := "TOKEN"

		// Construct the request payload
		requestPayload := PipelineRequest{
			Requests: []Request{
				{Type: "execute", Stmt: &Stmt{SQL: "SELECT 1"}},
				{Type: "close"},
			},
		}

		// Convert the payload to JSON
		jsonData, err := json.Marshal(requestPayload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to encode JSON"})
		}

		// Prepare the HTTP request
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create request"})
		}

		// Add headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+authToken)

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send request"})
		}
		defer resp.Body.Close()

		// Read and return the response
		var responseBody map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read response"})
		}

		return c.JSON(http.StatusOK, responseBody)
	})

	e.Start(":8080")
}
