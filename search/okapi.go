package search

import (
	"math"
	"os"
	"strings"
	"x134-search/index"
)

const WHITESPACE = " "
const K1 = 1.5
const B = 0.75

func idf(keyword string) float64 {
	N := float64(GetTotalDocuments())
	n := float64(GetKeywordContainingDocuments(keyword))

	rawScore := (N-n+0.5)/(n+0.5) + 1
	return math.Log(rawScore)
}

func queryDocuments(query string) map[string]float64 {
	keywords := strings.Split(query, WHITESPACE)

	noKeywords := len(keywords)
	idfScores := make([]float64, noKeywords)
	for i, keyword := range keywords {
		idfScores[i] = idf(keyword)
	}

	files, _ := os.ReadDir(index.PAGES_DIR)
	fileScores := make(map[string]float64)

	for i, keyword := range keywords {
		for _, file := range files {
			name := file.Name()
			f := GetKeywordCountInDoc(name, keyword)

			nonIDFScoreNumerator := float64(f) * (K1 + 1)
			nonIDFScoreDenominator := float64(f) + K1*(1-B+B*(float64(GetDocumentLength(name))/GetAverageDocumentLength()))
			nonIDFScore := nonIDFScoreNumerator / nonIDFScoreDenominator

			finalScore := idfScores[i] * nonIDFScore
			fileScores[name] += finalScore
		}
	}

	return fileScores
}
