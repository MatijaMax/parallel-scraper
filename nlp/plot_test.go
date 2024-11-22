package nlp

import (
	"testing"
)

func TestCreatePieChart(t *testing.T) {
	pos, neg := 70.0, 30.0
	name := "test-pie-chart.png"
	err := CreatePieChart(pos, neg, name)
	if err != nil {
		t.Errorf("CreatePieChart returned an error: %v", err)
	}
	/*
		filePath := "../data/test-pie-chart.png"
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist", filePath)
		} else {
			os.Remove(filePath)
		} */
}
