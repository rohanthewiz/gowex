package main

import (
    "github.com/rohanthewiz/element"
)

type cssContent struct{}

func (cs cssContent) Render(b *element.Builder) (x any) {
    t := b.Text

    // Create CSS directory and file
    b.Style().R(t(
        `html, body {
    height: 100%;
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

#format-button {
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
}`,
    ))
    return
}
