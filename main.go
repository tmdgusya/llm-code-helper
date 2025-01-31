package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	ignore "github.com/sabhiram/go-gitignore"
)

type Config struct {
	Dir         string `json:"dir"`
	FilePattern string `json:"file_extension"`
	IgnoreFiles string `json:"ignore_files"`
}

func main() {
	config := loadConfig()
	files := crawlFiles(config)
	generatePrompt(files)
}

func loadConfig() *Config {
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}

	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		panic(fmt.Errorf("invalid config format: %w", err))
	}

	return &config
}

func crawlFiles(config *Config) []string {
	patterns := strings.Split(config.FilePattern, "|")
	ignorePatterns := strings.Split(config.IgnoreFiles, "\n")

	gi := ignore.CompileIgnoreLines(ignorePatterns...)

	var files []string
	filepath.WalkDir(config.Dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(config.Dir, path)
		unixPath := filepath.ToSlash(relPath)

		// Check against gitignore patterns
		if gi.MatchesPath(unixPath) {
			return nil
		}

		// File extension matching
		for _, pattern := range patterns {
			if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
				files = append(files, path)
				break
			}
		}
		return nil
	})

	return files
}

func generatePrompt(files []string) {
	instruction, _ := os.ReadFile("instruct.prompt")
	var output strings.Builder

	output.WriteString(string(instruction))
	output.WriteString("\n\n")

	for _, file := range files {
		content, _ := os.ReadFile(file)
		relPath, _ := filepath.Rel(".", file)
		output.WriteString(fmt.Sprintf("``````\n%s\n%s\n``````\n", relPath, string(content)))
	}

	os.WriteFile("output.prompt", []byte(output.String()), 0644)
}
