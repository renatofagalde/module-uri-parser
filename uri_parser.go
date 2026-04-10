package uriparser

import (
	"regexp"
	"strings"
)

type Result struct {
	Action string `json:"action"`
	URI    string `json:"uri"`
}

var uuidRegex = regexp.MustCompile(
	`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`,
)

var irregularPlurals = map[string]string{
	"companies":  "company",
	"categories": "category",
}

func singularize(word string) string {
	if singular, ok := irregularPlurals[word]; ok {
		return singular
	}
	return strings.TrimSuffix(word, "s")
}

func Parse(method, uri string) Result {
	parts := strings.Split(uri, "/")

	for i, part := range parts {
		if uuidRegex.MatchString(part) && i > 0 {
			resource := singularize(parts[i-1])
			parts[i] = "{" + resource + "_id}"
		}
	}

	return Result{
		Action: method,
		URI:    strings.Join(parts, "/"),
	}
}
