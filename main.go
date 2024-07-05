package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	crawler "gocrawler/crawler/wiki"
	"gocrawler/util"
	"os"
)

func main() {
	wikiList := crawler.CrawlWiki()

	file, err := os.Open("./kanji.csv")

	if err != nil {
		return
	}

	reader := csv.NewReader(bufio.NewReader(file))
	rows, err := reader.ReadAll()

	if err != nil {
		return
	}

	defer file.Close()

	result := []crawler.TableData{}

	table := util.CsvToCollection(rows)

	for _, wikiItem := range wikiList {
		wikiItem.MEANING = table[wikiItem.KANJI]
		result = append(result, wikiItem)
	}

	fmt.Println(result)
}
