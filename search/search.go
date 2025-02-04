package search

import (
	"sort"
	"strings"
)

const TOPN = 5

func GetTopNFiles(fileScores map[string]float64, topN int) []string {
	if topN <= 0 {
		topN = 5
	}

	sortedFiles := make([]string, 0, len(fileScores))
	for file := range fileScores {
		sortedFiles = append(sortedFiles, file)
	}

	sort.Slice(sortedFiles, func(i, j int) bool {
		return fileScores[sortedFiles[i]] > fileScores[sortedFiles[j]]
	})

	topFiles := sortedFiles
	if len(sortedFiles) > topN {
		topFiles = sortedFiles[:topN]
	}

	for i, file := range topFiles {
		topFiles[i] = ReverseFileNameMapping(file)
	}

	return topFiles
}

func ReverseFileNameMapping(file string) string {
	return strings.ReplaceAll(file, "_", "/")
}

func Search(query string) []string {
	fileScores := queryDocuments(query)
	topNPages := GetTopNFiles(fileScores, TOPN)

	return topNPages
}
