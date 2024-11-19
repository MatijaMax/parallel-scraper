package nlp

import (
	"fmt"
	"strings"

	"github.com/cdipaolo/sentiment"
)

func CleanData(comments []string) []string {
	var newComments []string
	fmt.Printf("DUZINA %d", len(comments))
	for i := 0; i < len(comments); i++ {
		comment := comments[i]
		line := strings.SplitN(comment, "###", 2)
		fmt.Printf("Cleaned comment: %s\n", line[1])
		newComments = append(newComments, line[1])
	}
	return newComments
}

func AnalyzeComments(comments []string) (map[string]int, map[string]float64) {

	model, err := sentiment.Restore()
	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}
	var positive, negative int
	percentages := map[string]float64{
		"Positive": 0.0,
		"Negative": 0.0,
	}

	for _, comment := range comments {
		analysis := model.SentimentAnalysis(comment, sentiment.English)
		if analysis.Score > 0 {
			positive++
		} else {
			negative++
		}
	}

	totalComments := len(comments)
	if totalComments > 0 {
		percentages["Positive"] = float64(positive) / float64(totalComments) * 100
		percentages["Negative"] = float64(negative) / float64(totalComments) * 100
	}

	counts := map[string]int{
		"Positive": positive,
		"Negative": negative,
	}

	return counts, percentages
}
