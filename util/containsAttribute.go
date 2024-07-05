package util

import "golang.org/x/net/html"

func ContainsAttribute(attrs []html.Attribute, attr html.Attribute) bool {
	for _, a := range attrs {
		if a.Key == attr.Key && a.Val == attr.Val {
			return true
		}
	}
	return false
}
