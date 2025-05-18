package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/rohanthewiz/rweb"
)

func exeHandler(s *rweb.Server) {
	s.Post("/api/execute", func(ctx rweb.Context) error {
		codeBytes := ctx.Request().Body()
		req := CodeRequest{}

		err := json.Unmarshal(codeBytes, &req)
		if err != nil {
			fmt.Println("error", err)
			return err
		}

		return ctx.WriteJSON(executeGoCode(req.Code))
	})
}

// ExecutionRequest represents the incoming request with code to execute
type CodeRequest struct {
	Code string `json:"code"`
}

// ExecutionResult represents the results of executing Go code
type ExecutionResult struct {
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
	ExecutionMs int64  `json:"executionMs"`
	Error       string `json:"error,omitempty"`
	Success     bool   `json:"success"`
}

// executeGoCode creates a temporary file with the provided code, executes it, and returns the results
func executeGoCode(code string) ExecutionResult {
	start := time.Now()

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "goexec")
	if err != nil {
		return ExecutionResult{
			Error:   fmt.Sprintf("Failed to create temp directory: %v", err),
			Success: false,
		}
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(tempDir)

	// Create a temporary file for the Go code
	tempFile := tempDir + "/main.go"
	if err := os.WriteFile(tempFile, []byte(code), 0644); err != nil {
		return ExecutionResult{
			Error:   fmt.Sprintf("Failed to write temp file: %v", err),
			Success: false,
		}
	}

	// Run the code with 'go run'
	cmd := exec.Command("go", "run", tempFile)

	// Capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	duration := time.Since(start)

	result := ExecutionResult{
		Stdout:      stdout.String(),
		Stderr:      stderr.String(),
		ExecutionMs: duration.Milliseconds(),
		Success:     err == nil,
	}

	if err != nil {
		result.Error = err.Error()
	}

	return result
}

// executeHandler handles the API endpoint for executing code
func executeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CodeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := executeGoCode(req.Code)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
