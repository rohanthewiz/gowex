// Set up Monaco loader first
require.config({ paths: { 'vs': 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.52.2/min/vs' } });

// Then load the editor
require(['vs/editor/editor.main'], function() {
    // Default sample code
    const defaultCode = 'package main\n\nimport "fmt"\n\n//your code here\nfunc main() {fmt.Println("Hello, Web Go Executor!")}';

    // Create the editor once Monaco is fully loaded
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
            executionTime.textContent =` + "`Execution time: ${result.executionMs} ms`;" + `
            
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
});