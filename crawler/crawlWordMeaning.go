package crawler

import (
	"gocrawler/util"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetWords(n *html.Node) string {
	res := ""
	if util.ContainsAttribute(n.Attr, html.Attribute{Key: "class", Val: "cleanword_type kujk_type"}) {
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			if (util.ContainsAttribute(child.Attr, html.Attribute{Key: "class", Val: "list_search"})) {
				for c := child.FirstChild; c != nil; c = c.NextSibling {
					if c.Data == "li" {
						res += " " + ExtractText(c)
					}
				}
			}
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		word := (GetWords(child))
		if len(word) > 0 {
			res += word
		}
	}

	return strings.TrimSpace(res)
}

func CrawlWordMeaning(kanji string) string {
	url := "https://dic.daum.net/search.do?q=" + url.QueryEscape(kanji) + "&dic=jp"

	doc, err := util.FetchPage(url)

	if err != nil {
		return ""
	}

	words := GetWords(doc)
	return words
}
