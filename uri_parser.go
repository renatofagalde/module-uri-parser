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

// numericRegex casa um segmento inteiramente numérico (ex.: id serial do banco,
// como user_profile.id). Sem isto, só UUIDs viravam placeholder e rotas com id
// numérico (ex.: DELETE /cam/users/{user_id}/profiles/11) não casavam o grant,
// resultando em 403 mesmo com a permissão concedida.
var numericRegex = regexp.MustCompile(`^\d+$`)

// isID decide se um segmento do path é um identificador (UUID ou numérico) que
// deve virar o placeholder {resource_id}.
func isID(part string) bool {
	return uuidRegex.MatchString(part) || numericRegex.MatchString(part)
}

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
		if i > 0 && isID(part) {
			resource := singularize(parts[i-1])
			parts[i] = "{" + resource + "_id}"
		}
	}

	return Result{
		Action: method,
		URI:    strings.Join(parts, "/"),
	}
}
