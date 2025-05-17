package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// indexHandler serves the main page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

// Todo add GoImports
func main() {
	// Ensure the templates directory exists
	if err := os.MkdirAll("templates", 0755); err != nil {
		log.Fatalf("Failed to create templates directory: %v", err)
	}

	// Create the index.html template ** always **
	if err := createIndexTemplate(); err != nil {
		log.Fatalf("Failed to create index template: %v", err)
	}

	// Create the static directory and files
	if err := os.MkdirAll("static/css", 0755); err != nil {
		log.Fatalf("Failed to create static directory: %v", err)
	}

	if err := createStaticFiles(); err != nil {
		log.Fatalf("Failed to create static files: %v", err)
	}

	// Set up HTTP routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/execute", executeHandler)
	// Add this line to the main function after the existing routes
	http.HandleFunc("/api/format", formatHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	port := ":8080"
	log.Printf("Starting server on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// createIndexTemplate creates the HTML template file
func createIndexTemplate() error {
	indexHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Code Executor</title>
    <link rel="stylesheet" href="/static/css/style.css">
    <!-- Monaco Editor -->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.52.2/min/vs/loader.min.js"></script>
</head>
<body>
    <div class="app-container">
        <header>
            <h1>Go Code Executor</h1>
        </header>
        
        <main>
            <div class="editor-container">
                <div id="editor"></div>
			<div class="button-container">
				<button id="format-button">Format</button>
				<button id="run-button">Run (Ctrl+Enter)</button>
            </div>
            
            <div class="output-container">
                <div class="output-header">
                    <h2>Execution Results</h2>
                    <div id="execution-status">Ready</div>
                </div>
                
                <div class="output-content">
                    <div class="output-section">
                        <h3>Standard Output</h3>
                        <pre id="stdout-output" class="output-area"></pre>
                    </div>
                    
                    <div class="output-section">
                        <h3>Standard Error</h3>
                        <pre id="stderr-output" class="output-area error"></pre>
                    </div>
                    
                    <div class="execution-info">
                        <div id="execution-time"></div>
                        <div id="execution-result"></div>
                    </div>
                </div>
            </div>
        </main>
        
        <footer>
            <p>Go Code Executor - A web-based Go code execution environment</p>
        </footer>
    </div>

    <script src="/static/js/app.js"></script>
</body>
</html>`

	return os.WriteFile("templates/index.html", []byte(indexHTML), 0644)
}

// createStaticFiles creates CSS and JS files for the web interface
func createStaticFiles() error {
	// Create CSS directory and file
	cssContent := `body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    margin: 0;
    padding: 0;
    background-color: #808080;
    color: #333;
}

.app-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    max-width: 1200px;
    margin: 0 auto;
    background-color: #bbb;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

header {
    background-color: #874BFD;
    color: white;
    padding: 1rem;
    text-align: center;
}

h1 {
    margin: 0;
}

main {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    padding: 1rem;
}

.editor-container {
    flex: 1;
    min-height: 300px;
    margin-bottom: 1rem;
    border: 1px solid #ddd;
    border-radius: 5px;
    overflow: hidden;
    display: flex;
    flex-direction: column;
}

#editor {
    height: 300px;
    width: 100%;
}

.button-container {
    padding: 0.5rem;
    background-color: #f5f5f5;
    display: flex;
    justify-content: flex-end;
}

#run-button {
    background-color: #874BFD;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
}

#run-button:hover {
    background-color: #7038e0;
}

format-button {
    background-color: #4CAF50;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
    margin-right: 10px;
}

#format-button:hover {
    background-color: #3e8e41;
}

.output-container {
    flex: 1;
    border: 1px solid #ddd;
    border-radius: 5px;
    overflow: hidden;
    display: flex;
    flex-direction: column;
}

.output-header {
    background-color: #f5f5f5;
    padding: 0.5rem 1rem;
    border-bottom: 1px solid #ddd;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.output-header h2 {
    margin: 0;
    font-size: 1.2rem;
}

#execution-status {
    font-size: 0.9rem;
}

.output-content {
    flex: 1;
    padding: 1rem;
    overflow-y: auto;
}

.output-section {
    margin-bottom: 1rem;
}

.output-section h3 {
    margin-top: 0;
    font-size: 1rem;
    color: #666;
}

.output-area {
    background-color: #f8f8f8;
    padding: 0.5rem;
    border-radius: 4px;
    max-height: 200px;
    overflow-y: auto;
    white-space: pre-wrap;
    margin: 0;
}

.error {
    color: #e53935;
}

.success {
    color: #43a047;
}

.execution-info {
    margin-top: 1rem;
    display: flex;
    justify-content: space-between;
    font-size: 0.9rem;
    color: #666;
}

footer {
    background-color: #f5f5f5;
    padding: 1rem;
    text-align: center;
    font-size: 0.9rem;
    color: #666;
    border-top: 1px solid #ddd;
}

@media (min-width: 768px) {
    main {
        flex-direction: row;
    }
    
    .editor-container {
        flex: 1;
        margin-right: 1rem;
        margin-bottom: 0;
    }
    
    .output-container {
        flex: 1;
    }
}`

	if err := os.WriteFile("static/css/style.css", []byte(cssContent), 0644); err != nil {
		return err
	}

	// Create JS directory and file
	if err := os.MkdirAll("static/js", 0755); err != nil {
		return err
	}

	jsContent := `// Initialize Monaco Editor
require.config({ paths: { 'vs': 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.52.2/min/vs' } });

require(['vs/editor/editor.main'], function() {
    // Default sample code
    const defaultCode = 'package main\n\nimport "fmt"\n\n//your code here\nfunc main() {fmt.Println("Hello, Web Go Executor!")}';

    // Create the editor
    const editor = monaco.editor.create(document.getElementById('editor'), {
        value: defaultCode,
        language: 'go',
        theme: 'vs',
        automaticLayout: true,
        minimap: {
            enabled: false
        },
        scrollBeyondLastLine: false,
        lineNumbers: 'on',
        renderLineHighlight: 'all',
        tabSize: 4,
        insertSpaces: false
    });


    // Set up run button and keyboard shortcut
    const runButton = document.getElementById('run-button');
    
    function executeCode() {
        const code = editor.getValue();
        runCode(code);
    }
    
    runButton.addEventListener('click', executeCode);
    
    // Add Ctrl+Enter shortcut
    editor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.Enter, executeCode);

const formatButton = document.getElementById('format-button');
    
    function formatCode() {
        const code = editor.getValue();
        formatGoCode(code);
    }
    
    formatButton.addEventListener('click', formatCode);

    // Add keyboard shortcut (Ctrl+Shift+F) for formatting
    editor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyMod.Shift | monaco.KeyCode.KeyF, formatCode);


    
    // Elements for displaying results
    const executionStatus = document.getElementById('execution-status');
    const stdoutOutput = document.getElementById('stdout-output');
    const stderrOutput = document.getElementById('stderr-output');
    const executionTime = document.getElementById('execution-time');
    const executionResult = document.getElementById('execution-result');


    async function formatGoCode(code) {
		formatButton.disabled = true;

		try {
			const response = await fetch('/api/format', {
			method: 'POST',
			headers: {
			'Content-Type': 'application/json'
		},
			body: JSON.stringify({ code })
		});

			if (!response.ok) {
			throw new Error('Failed to format code: ' + response.statusText);
		}

			const result = await response.json();

			if (result.success) {
			// Update editor with formatted code
			editor.setValue(result.formattedCode);
		} else {
			console.error('Format error:', result.error);
			alert('Failed to format code: ' + result.error);
		}
		} catch (error) {
			console.error('Error formatting code:', error);
			alert('Error formatting code: ' + error.message);
		} finally {
			formatButton.disabled = false;
		}
	}
    
    async function runCode(code) {
        // Update UI to show code is executing
        executionStatus.textContent = 'Executing...';
        executionStatus.className = '';
        runButton.disabled = true;
        
        // Clear previous results
        stdoutOutput.textContent = '';
        stderrOutput.textContent = '';
        executionTime.textContent = '';
        executionResult.textContent = '';
        executionResult.className = '';
        
        try {
            const response = await fetch('/api/execute', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ code })
            });
            
            if (!response.ok) {
                throw new Error('Failed to execute code: ' + response.statusText);
            }
            
            const result = await response.json();
            
            // Update UI with results
            stdoutOutput.textContent = result.stdout || 'No output';
            stderrOutput.textContent = result.stderr || 'No errors';
            executionTime.textContent = ` + "`Execution time: ${result.executionMs} ms`;" + `            
            if (result.success) {
                executionStatus.textContent = 'Completed';
                executionStatus.className = 'success';
                executionResult.textContent = 'Execution successful!';
                executionResult.className = 'success';
            } else {
                executionStatus.textContent = 'Failed';
                executionStatus.className = 'error';
                executionResult.textContent = 'Error: ' + (result.error || 'Unknown error');
                executionResult.className = 'error';
            }
        } catch (error) {
            executionStatus.textContent = 'Error';
            executionStatus.className = 'error';
            executionResult.textContent = 'Error: ' + error.message;
            executionResult.className = 'error';
            console.error('Error executing code:', error);
        } finally {
            runButton.disabled = false;
        }
    }
});`

	return os.WriteFile("static/js/app.js", []byte(jsContent), 0644)
}
