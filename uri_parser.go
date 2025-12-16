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

func Parse(method, uri string) Result {
	parts := strings.Split(uri, "/")

	for i, part := range parts {
		if uuidRegex.MatchString(part) && i > 0 {
			resource := strings.TrimSuffix(parts[i-1], "s")
			parts[i] = "id_" + resource
		}
	}

	return Result{
		Action: method,
		URI:    strings.Join(parts, "/"),
	}
}
