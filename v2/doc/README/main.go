package main

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"
	"fmt"
)

// Structure to hold information about a Markdown file
type MarkdownFile struct {
	Name string
	Path string
}

// Function to generate a table of contents and links to all Markdown files
func generateReadme(files []MarkdownFile) string {
	var sb strings.Builder
	sb.WriteString("## PromiseGrid grid-cli v2 Documentation\n\n")
	sb.WriteString("### Table of Contents\n\n")

	for _, file := range files {
		// Create a link for each markdown file
		relativePath := strings.TrimPrefix(file.Path, "./")
		link := fmt.Sprintf("- [%s](./%s)\n", file.Name, relativePath)
		sb.WriteString(link)
	}

	sb.WriteString("\n")

	return sb.String()
}

// Function to recursively scan the directory and find all Markdown files
func scanDirectory(directory string) ([]MarkdownFile, error) {
	var files []MarkdownFile

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			files = append(files, MarkdownFile{
				Name: strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
				Path: path,
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func main() {
	// Specify the root documentation directory
	directory := "./"

	// Scan the directory for Markdown files
	files, err := scanDirectory(directory)
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		os.Exit(1)
	}

	// Generate the README content
	readmeContent := generateReadme(files)

	// Write the generated content to the README.md file
	err = ioutil.WriteFile("README.md", []byte(readmeContent), 0644)
	if err != nil {
		fmt.Println("Error writing README.md:", err)
		os.Exit(1)
	}

	fmt.Println("README.md successfully generated.")
}
