package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

// FormatResult represents the result of formatting Go code
type FormatResult struct {
	FormattedCode string `json:"formattedCode"`
	Error         string `json:"error,omitempty"`
	Success       bool   `json:"success"`
}

// formatGoCode formats the provided Go code
func formatGoCode(code string) FormatResult {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "gofmt")
	if err != nil {
		return FormatResult{
			Error:   fmt.Sprintf("Failed to create temp directory: %v", err),
			Success: false,
		}
	}
	defer os.RemoveAll(tempDir)

	// Create a temporary file for the Go code
	tempFile := tempDir + "/main.go"
	if err := os.WriteFile(tempFile, []byte(code), 0644); err != nil {
		return FormatResult{
			Error:   fmt.Sprintf("Failed to write temp file: %v", err),
			Success: false,
		}
	}

	// Run go fmt on the file
	cmd := exec.Command("gofmt", "-s", tempFile)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return FormatResult{
			Error:   fmt.Sprintf("Failed to format code: %v", err),
			Success: false,
		}
	}

	// Read the formatted code
	formattedCode := stdout.String()
	if formattedCode == "" {
		// If gofmt doesn't output anything, read the original file
		// (gofmt only outputs to stdout if the code is piped in)
		content, err := os.ReadFile(tempFile)
		if err != nil {
			return FormatResult{
				Error:   fmt.Sprintf("Failed to read formatted file: %v", err),
				Success: false,
			}
		}
		formattedCode = string(content)
	}

	return FormatResult{
		FormattedCode: formattedCode,
		Success:       true,
	}
}

// formatHandler handles the API endpoint for formatting code
func formatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ExecutionRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := formatGoCode(req.Code)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
