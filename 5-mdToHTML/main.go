package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/russross/blackfriday/v2"
)

func main() {
    // Check if a Markdown file is provided as a command-line argument
    if len(os.Args) != 2 {
        log.Fatal("Usage: go run main.go sample.md")
    }

    // Read the input Markdown file
    inputFileName := os.Args[1]
    markdownData, err := os.ReadFile(inputFileName)
    if err != nil {
        log.Fatal(err)
    }

    // Get the base name of the input file (without the extension)
    baseName := filepath.Base(inputFileName)

    // Remove the ".md" extension from the base name
    baseNameWithoutExtension := baseName[:len(baseName)-len(filepath.Ext(baseName))]

    // Change the file extension to ".html" for the output file
    outputFileName := fmt.Sprintf("%s.html", baseNameWithoutExtension)

    // Convert Markdown to HTML
    htmlData := blackfriday.Run(markdownData)

    // Define CSS styles
    cssStyles := `
        <style>
            /* Include the contents of styles.css here */
            body {
                font-family: Arial, sans-serif;
                background-color: #f2f2f2;
                margin: 20px;
            }

            h1, h2, h3, h4 {
                color: #007bff;
            }

            p {
                color: #333;
                line-height: 1.5;
            }

            a {
                color: #007bff;
                text-decoration: none;
            }

            a:hover {
                text-decoration: underline;
            }

            img {
                max-width: 100%;
                height: auto;
                border: 2px solid #ccc;
                border-radius: 5px;
                margin: 10px 0;
            }

            code {
                background-color: #f5f5f5;
                padding: 2px 5px;
                border-radius: 3px;
                font-family: Consolas, monospace;
            }

            blockquote {
                background-color: #f9f9f9;
                border-left: 5px solid #ccc;
                padding: 10px;
                margin: 10px 0;
            }

            hr {
                border: 1px solid #ccc;
            }

            table {
                width: 100%;
                border-collapse: collapse;
            }

            th, td {
                border: 1px solid #ddd;
                padding: 8px;
                text-align: left;
            }

            th {
                background-color: #f2f2f2;
            }
        </style>
    `

    // Combine CSS styles and HTML content
    finalHTML := fmt.Sprintf("<html><head>%s</head><body>%s</body></html>", cssStyles, htmlData)

    // Write the final HTML to the output file
    if err := os.WriteFile(outputFileName, []byte(finalHTML), 0644); err != nil {
        log.Fatal(err)
    }

    log.Printf("Markdown converted to HTML and saved as %s\n", outputFileName)
}
