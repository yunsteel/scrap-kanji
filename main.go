package main

import (
	"bufio"
	"encoding/csv"
	"gocrawler/crawler"
	"gocrawler/util"
	"os"
)

func main() {
	kanjiList := crawler.CrawlWiki()

	file, err := os.Open("./reference/kanji.csv")

	if err != nil {
		return
	}

	reader := csv.NewReader(bufio.NewReader(file))
	rows, err := reader.ReadAll()

	if err != nil {
		return
	}

	defer file.Close()

	table := [][]string{
		{"아이디",
			"한자",
			"구자",
			"부수",
			"총획",
			"학년",
			"추가연도",
			"삭제연도",
			"발음",
			"뜻"},
	}

	kanjiMap := util.CsvToCollection(rows)

	for _, kanjiItem := range kanjiList {
		kanjiItem.MEANING = kanjiMap[kanjiItem.KANJI]

		row := []string{kanjiItem.ID, kanjiItem.KANJI, kanjiItem.OLD_KANJI, kanjiItem.RADICAL, kanjiItem.STROKE, kanjiItem.GRADE, kanjiItem.YEAR_ADDED, kanjiItem.YEAR_DELETED, kanjiItem.PRONUNCIATION, kanjiItem.MEANING}
		table = append(table, row)
	}

	util.SaveToCSV(table, "./table.csv")
}
