package main

import (
	"os"
	"testing"
)

func TestParsePatchFile(t *testing.T) {
	err := parsePatchFile("tests/test.patch")
	if err != nil {
		t.Fatalf("Error parsing patch file: %v", err)
	}

	expectedFiles := []string{
		".test.yml.patch",
		"CHANGELOG.md.patch",
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Fatalf("Expected " + file + " to be created, but it was not")
		}
		defer os.Remove(file)
	}
}
