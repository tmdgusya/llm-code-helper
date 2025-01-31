package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func setupTestEnvironment() {
	_ = os.MkdirAll("src/dist", 0755)
	_ = os.MkdirAll("src", 0755)

	_ = os.WriteFile("config.json", []byte(`{
		"dir" : "src/",
		"file_extension" : "*.kt|*.ts|*.js",
		"ignore_files": "dist/**"
	}`), 0644)

	_ = os.WriteFile("instruct.prompt", []byte("Test Instruction"), 0644)
	_ = os.WriteFile("src/test.js", []byte("console.log('Hello JS');"), 0644)
	_ = os.WriteFile("src/test.kt", []byte("fun main() { println(\"Hello Kotlin\") }"), 0644)
	_ = os.WriteFile("src/test.ts", []byte("console.log('Hello TS');"), 0644)
	_ = os.WriteFile("src/dist/test2.js", []byte("console.log('Ignore this JS');"), 0644)
}

func teardownTestEnvironment() {
	_ = os.RemoveAll("src")
	_ = os.Remove("config.json")
	_ = os.Remove("instruct.prompt")
	_ = os.Remove("output.prompt")
}

func TestLoadConfig(t *testing.T) {
	setupTestEnvironment()
	defer teardownTestEnvironment()

	config := loadConfig()
	if config.Dir != "src/" {
		t.Fatalf("Expected dir to be 'src/', got %s", config.Dir)
	}
	if config.FilePattern != "*.kt|*.ts|*.js" {
		t.Fatalf("Expected file pattern to be '*.kt|*.ts|*.js', got %s", config.FilePattern)
	}
	if config.IgnoreFiles != "dist/**" {
		t.Fatalf("Expected ignore files to be 'dist/**', got %s", config.IgnoreFiles)
	}
}

func TestCrawlFiles(t *testing.T) {
	setupTestEnvironment()
	defer teardownTestEnvironment()

	config := &Config{
		Dir:         "src/",
		FilePattern: "*.kt|*.ts|*.js",
		IgnoreFiles: "dist/**",
	}

	files := crawlFiles(config)

	expectedFiles := []string{
		"src/test.js",
		"src/test.kt",
		"src/test.ts",
	}

	if len(files) != len(expectedFiles) {
		t.Fatalf("Expected %d files, got %d files: %v", len(expectedFiles), len(files), files)
	}

	for _, expected := range expectedFiles {
		found := false
		for _, file := range files {
			if strings.Contains(file, expected) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s not found in result: %v", expected, files)
		}
	}
}

func TestGeneratePrompt(t *testing.T) {
	setupTestEnvironment()
	defer teardownTestEnvironment()

	files := []string{
		"src/test.js",
		"src/test.kt",
		"src/test.ts",
	}
	generatePrompt(files)

	content, err := os.ReadFile("output.prompt")
	if err != nil {
		t.Fatalf("Error reading output file: %v", err)
	}

	if !strings.Contains(string(content), "Test Instruction") {
		t.Errorf("Expected 'Test Instruction' in output.prompt")
	}
	for _, file := range files {
		if !strings.Contains(string(content), filepath.Base(file)) {
			t.Errorf("Expected file content for %s not found in output.prompt", file)
		}
	}
}
