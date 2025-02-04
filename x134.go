package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"x134-search/index"
	"x134-search/search"
)

const (
	INDEX_MODE  = "index"
	SEARCH_MODE = "search"
	USAGE_MSG   = "Usage: ./x134 <mode> [optional args]"
)

func parseArgs() (string, string) {
	sites := flag.String("sites", "sites.txt", "Provide the sites.txt file")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println(USAGE_MSG)
		os.Exit(1)
	}

	return args[0], *sites
}

func searchMode() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter search query: ")
	query, _ := reader.ReadString('\n')
	query = query[:len(query)-1]

	results := search.Search(query)

	fmt.Println("\nSearch Results:")
	for i, url := range results {
		fmt.Printf("%d. %s\n", i+1, url)
	}
}

func executeMode(mode string, sites string) {
	switch mode {
	case INDEX_MODE:
		index.Index(sites)
	case SEARCH_MODE:
		searchMode()
	default:
		fmt.Println("Mode not recognized:", mode)
		os.Exit(1)
	}
}

func main() {
	mode, sites := parseArgs()
	executeMode(mode, sites)
}
