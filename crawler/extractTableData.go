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
		if slices.Contains(c.Attr, html.Attribute{Key: "style", Val: "display:none"}) {
			continue
		}
		if c.Data == "sup" {
			continue
		}
		childText := ExtractText(c)
		text += childText
	}

	return strings.TrimSpace(text)
}

type TableData struct {
	ID            string
	KANJI         string
	OLD_KANJI     string
	RADICAL       string
	STROKE        string
	GRADE         string
	YEAR_ADDED    string
	YEAR_DELETED  string
	PRONUNCIATION string
	HUN_UM        string
	MEANING       string
}

func ExtractTableData(doc *html.Node) ([]TableData, error) {
	var tableData = []TableData{}
	var extract func(*html.Node)

	extract = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			for child := n.FirstChild; child != nil; child = child.NextSibling {
				if child.Type == html.ElementNode && child.Data == "tbody" {
					for tr := child.FirstChild; tr != nil; tr = tr.NextSibling {
						if tr.Type == html.ElementNode && tr.Data == "tr" {
							var row TableData

							i := 0
							for td := tr.FirstChild; td != nil; td = td.NextSibling {
								if td.Type == html.ElementNode && td.Data == "td" {
									value := ExtractText(td)
									switch i {
									case 0:
										row.ID = value
									case 1:
										row.KANJI = value
									case 2:
										row.OLD_KANJI = value
									case 3:
										row.RADICAL = value
									case 4:
										row.STROKE = value
									case 5:
										row.GRADE = value
									case 6:
										row.YEAR_ADDED = value
									case 7:
										row.YEAR_DELETED = value
									case 8:
										row.PRONUNCIATION = value
									}
									i++
								}
							}

							if len(row.ID) > 0 && row.ID != "0" {
								tableData = append(tableData, row)
							}
						}
					}
					break
				}
			}
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
