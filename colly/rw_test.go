package colly

import (
	"os"
	"strings"
	"testing"
)

func TestWriteComments(t *testing.T) {
	file := "../data/test.txt"
	comments := []string{"COMMENT### This is a comment", "COMMENT### Another comment"}
	err := WriteComments(file, comments)
	if err != nil {
		t.Fatalf("Failed to write comments: %v", err)
	}
	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	for _, c := range comments {
		if !contains(content, c) {
			t.Errorf("Expected to find comment: %s", c)
		}
	}
}

func TestReadComments(t *testing.T) {
	file := "../data/test.txt"
	comments := []string{"COMMENT### This is a comment", "COMMENT### Another comment"}
	writeComments := []string{"COMMENT### This is a comment\n", "COMMENT### Another comment\n"}
	err := WriteComments(file, writeComments)
	if err != nil {
		t.Fatalf("Failed to write comments: %v", err)
	}
	readComments, err := ReadComments(file)
	if err != nil {
		t.Fatalf("Failed to read comments: %v", err)
	}
	for i, c := range comments {
		if readComments[i] != c {
			t.Errorf("Expected %s but got %s", c, readComments[i])
		}
	}
}

func contains(content []byte, substr string) bool {
	return strings.Contains(string(content), substr)
}
