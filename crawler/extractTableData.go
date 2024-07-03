package crawler

import (
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func ExtractText(n *html.Node) string {
	var text string
	if n.Type == html.TextNode {
		text = n.Data
	} else if n.Type == html.ElementNode && n.Data == "br" {
		text = "\n"
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childText := ExtractText(c)
		text += childText
	}

	return strings.TrimSpace(text)
}

func FilterRow(row *[]string) []string {
	excludedHeader := []string{"구자", "삭제연도"}

	res := []string{}

	for _, rowItem := range *row {
		if !slices.Contains(excludedHeader, rowItem) {
			res = append(res, rowItem)
		}
	}

	return res
}

func ExtractTableData(doc *html.Node) ([][]string, error) {
	var tableData [][]string
	var extract func(*html.Node)

	extract = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			var header = []string{"아이디", "한자", "구자", "부수", "총획", "학년", "추가연도", "삭제연도", "발음"}
			var rows [][]string

			for child := n.FirstChild; child != nil; child = child.NextSibling {
				if child.Type == html.ElementNode && child.Data == "tbody" {
					for tr := child.FirstChild; tr != nil; tr = tr.NextSibling {
						if tr.Type == html.ElementNode && tr.Data == "tr" {
							var row []string
							for td := tr.FirstChild; td != nil; td = td.NextSibling {
								if td.Type == html.ElementNode && td.Data == "td" {
									row = append(row, ExtractText(td))
								}
							}

							if len(row) > 0 {
								rows = append(rows, FilterRow(&row))
							}
						}
					}
					break
				}
			}

			tableData = append(tableData, FilterRow(&header))
			tableData = append(tableData, rows...)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if len(tableData) == 0 {
				extract(c)
			}
		}
	}

	extract(doc)
	return tableData, nil
}
