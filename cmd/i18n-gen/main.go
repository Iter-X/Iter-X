//go:generate go run main.go -O ../../i18n/ -P ../../proto/xerr/errors.proto
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/emicklei/proto"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Define flags
	protoFile := flag.String("P", "internal/common/xerr/errors.proto", "Path to the .proto file")
	outputDir := flag.String("O", "./i18n/", "Path to the output directory")
	flag.Parse()

	// Parse the .proto file
	entries, err := parseProto(*protoFile)
	if err != nil {
		log.Printf("Failed to parse proto file: %v\n", err)
		return
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Printf("Failed to create output directory: %v\n", err)
		return
	}

	// Generate or update TOML files
	for _, v := range []string{"en", "zh"} {
		tomlPath := fmt.Sprintf("%s/%s.toml", *outputDir, v)
		if err := generateTOML(entries, tomlPath); err != nil {
			log.Printf("Failed to generate %s.toml: %v\n", v, err)
			return
		}
		log.Printf("%s.toml generated/updated successfully.", v)
	}
}

// parseProto reads the .proto file and extracts enum names as keys in order.
func parseProto(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open proto file: %w", err)
	}
	defer file.Close()

	var entries []string
	reader := bufio.NewReader(file)
	parser := proto.NewParser(reader)

	definition, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("parse proto: %w", err)
	}

	proto.Walk(definition,
		proto.WithEnum(func(e *proto.Enum) {
			for _, elem := range e.Elements {
				if field, ok := elem.(*proto.EnumField); ok {
					entries = append(entries, field.Name)
				}
			}
		}),
	)

	return entries, nil
}

// generateTOML updates or creates a TOML file based on the provided entries and maintains order.
func generateTOML(entries []string, filePath string) error {
	existingEntries, err := loadExistingTOML(filePath)
	if err != nil {
		return fmt.Errorf("load existing TOML: %w", err)
	}

	// Merge existing entries while maintaining order
	entryMap := make(map[string]string)
	for _, entry := range entries {
		if val, exists := existingEntries[entry]; exists {
			entryMap[entry] = val
		} else {
			entryMap[entry] = ""
		}
	}

	// Generate TOML content
	var buffer bytes.Buffer
	for _, entry := range entries {
		buffer.WriteString(fmt.Sprintf("[%s]\nother = \"%s\"\n\n", entry, entryMap[entry]))
	}

	// Write the updated content to the file
	if err := os.WriteFile(filePath, buffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("write TOML file: %w", err)
	}

	return nil
}

// loadExistingTOML parses an existing TOML file into a map of keys with their values.
func loadExistingTOML(filePath string) (map[string]string, error) {
	entries := make(map[string]string)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return entries, nil // File does not exist, return empty map
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open TOML file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentKey string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentKey = line[1 : len(line)-1]
		} else if strings.HasPrefix(line, "other = ") {
			if currentKey != "" {
				entries[currentKey] = strings.Trim(line[len("other = "):], "\"")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read TOML file: %w", err)
	}

	return entries, nil
}

// snakeToCamelCase converts snake_case to CamelCase.
func snakeToCamelCase(input string) string {
	words := strings.Split(input, "_")
	caser := cases.Title(language.English) // Create a caser for proper title casing
	for i := range words {
		words[i] = caser.String(strings.ToLower(words[i]))
	}
	return strings.Join(words, "")
}
