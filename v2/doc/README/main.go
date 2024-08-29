package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/stevegt/goadapt"
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

	// - the table of contents is a markdown table with two columns: link and heading
	// - the link is a markdown link to the file
	sb.WriteString("| File | Subject |\n")
	sb.WriteString("| ---- | ------- |\n")

	for _, file := range files {
		// Extract the first heading from the markdown file
		// This will be used as the section title in the table of contents
		heading, err := extractFirstHeading(file.Path)
		if err != nil {
			fmt.Println("Error extracting first heading:", err)
			os.Exit(1)
		}
		// create the title and link for the markdown file
		relativePath := strings.TrimPrefix(file.Path, "./")
		link := fmt.Sprintf("[%s](./%s)", file.Name, relativePath)
		line := fmt.Sprintf("| %s | %s |\n", link, heading)
		sb.WriteString(line)
	}

	sb.WriteString("\n")

	return sb.String()
}

// Function to extract the first heading from a Markdown file
func extractFirstHeading(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file line by line until we find a heading
	var heading string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// get next line
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			// remove any leading whitespace and hash symbols
			// - there may be more than one hash symbol; we want to
			// remove all of them
			heading = strings.TrimLeft(line, " #")
			break
		}
	}
	err = scanner.Err()
	Ck(err)

	if heading == "" {
		heading = path
	}

	return heading, nil
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
	directory := os.Args[1]

	// Scan the directory for Markdown files
	files, err := scanDirectory(directory)
	if err != nil {
		fmt.Println("Error scanning directory:", err)
		os.Exit(1)
	}

	// Generate the README content
	readmeContent := generateReadme(files)

	// Write the generated content to stdout
	fmt.Print(readmeContent)
}
