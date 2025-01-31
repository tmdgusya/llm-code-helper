package main

import (
	"encoding/json"
	"flag"
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
	// Define command-line flag for config path
	projectPath := flag.String("projectPath", "", "Path to project folder")
	flag.Parse()

	config := loadConfig(*projectPath)
	files := crawlFiles(config, *projectPath)
	generatePrompt(config, files, *projectPath)
}

func loadConfig(projectPath string) *Config {
	configFile, err := os.ReadFile(filepath.Join(projectPath, "llm-helper-config.json"))
	if err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}

	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		panic(fmt.Errorf("invalid config format: %w", err))
	}

	return &config
}

func crawlFiles(config *Config, projectPath string) []string {
	patterns := strings.Split(config.FilePattern, "|")
	ignorePatterns := strings.Split(config.IgnoreFiles, "\n")

	gi := ignore.CompileIgnoreLines(ignorePatterns...)

	var files []string
	filepath.WalkDir(filepath.Join(projectPath, config.Dir), func(path string, d fs.DirEntry, err error) error {
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

func generatePrompt(config *Config, files []string, projectPath string) {
	instruction, _ := os.ReadFile(filepath.Join(projectPath, "instruct.prompt"))
	var output strings.Builder

	output.WriteString(string(instruction))
	output.WriteString("\n\n")

	for _, file := range files {
		content, _ := os.ReadFile(file)
		relPath, _ := filepath.Rel(".", file)
		output.WriteString(fmt.Sprintf("``````\n%s\n%s\n``````\n", relPath, string(content)))
	}

	os.WriteFile(filepath.Join(projectPath, "output.prompt"), []byte(output.String()), 0644)
}
