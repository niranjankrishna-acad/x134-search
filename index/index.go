package index

import (
	"log"
	"os"
	"strings"
)

const NEWLINE = "\n"

func getSites(sitesFile string) []string {
	file, error := os.ReadFile(sitesFile)

	errorExists := error != nil
	if errorExists {
		log.Fatal(error)
	}

	text := string(file)
	sites := strings.Split(text, NEWLINE)

	return sites

}
func Index(sitesFile string) {
	sites := getSites(sitesFile)
	crawl(sites)
}
