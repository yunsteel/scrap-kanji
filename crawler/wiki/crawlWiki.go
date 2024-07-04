package crawler

import (
	"fmt"
	"gocrawler/util"
)

func CrawlWiki() {
	url := "https://ja.wikipedia.org/wiki/%E5%B8%B8%E7%94%A8%E6%BC%A2%E5%AD%97%E4%B8%80%E8%A6%A7#%E6%9C%AC%E8%A1%A8"
	filepath := "table.csv"

	doc, err := util.FetchPage(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	tableData, err := ExtractTableData(doc)

	if err != nil {
		fmt.Println(err)
		return
	}

	util.SaveToCSV(tableData, filepath)
}
