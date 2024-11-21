package nlp

import (
	"testing"
)

func TestCleanData(t *testing.T) {
	comments := []string{
		"COMMENT###This is a comment",
		"COMMENT###Another comment",
	}
	expected := []string{"This is a comment", "Another comment"}

	result := CleanData(comments)
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %s but got %s", expected[i], v)
		}
	}
}

func TestAnalyzeComments(t *testing.T) {
	comments := []string{"This is good", "This is bad", "Awesome product"}
	counts, _ := AnalyzeComments(comments)

	if counts["Positive"] == 0 {
		t.Errorf("Expected positive count > 0, got %d", counts["Positive"])
	}
	if counts["Negative"] == 0 {
		t.Errorf("Expected negative count > 0, got %d", counts["Negative"])
	}
}
