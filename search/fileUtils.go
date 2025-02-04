package search

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"x134-search/index"
)

var keywordFrequency = make(map[string]map[string]int)
var documentLengths = make(map[string]int)
var totalDocuments int
var averageDocumentLength float64
var keywordDocumentCount = make(map[string]int)

func AnalyzeFiles(keywords []string) {
	files, err := os.ReadDir(index.PAGES_DIR)
	if err != nil {
		return
	}

	totalLength := 0
	totalDocuments = len(files)

	for _, file := range files {
		filePath := filepath.Join(index.PAGES_DIR, file.Name())
		length := countWordsInFile(filePath)
		documentLengths[file.Name()] = length
		totalLength += length

		keywordFrequency[file.Name()] = make(map[string]int)
		for _, keyword := range keywords {
			count := countKeywordOccurrencesInFile(filePath, keyword)
			keywordFrequency[file.Name()][keyword] = count
			if count > 0 {
				keywordDocumentCount[keyword]++
			}
		}
	}

	if totalDocuments > 0 {
		averageDocumentLength = float64(totalLength) / float64(totalDocuments)
	}
}

func countWordsInFile(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}

	return count
}

func countKeywordOccurrencesInFile(filePath, keyword string) int {
	file, err := os.Open(filePath)
	if err != nil {
		return 0
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count += strings.Count(scanner.Text(), keyword)
	}

	return count
}

func GetTotalDocuments() int {
	return totalDocuments
}

func GetDocumentLength(docName string) int {
	return documentLengths[docName]
}

func GetAverageDocumentLength() float64 {
	return averageDocumentLength
}

func GetKeywordCountInDoc(docName, keyword string) int {
	return keywordFrequency[docName][keyword]
}

func GetKeywordContainingDocuments(keyword string) int {
	return keywordDocumentCount[keyword]
}
