package uriparser

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		method string
		uri    string
		want   string
	}{
		{"GET", "/pessoas/4152bd3d-9121-4889-b683-6ee72b8e3719", "/pessoas/id_pessoa"},
		{"POST", "/pessoas/4152bd3d-9121-4889-b683-6ee72b8e3719/telefones", "/pessoas/id_pessoa/telefones"},
		{"GET", "/pessoas/4152bd3d-9121-4889-b683-6ee72b8e3719/enderecos/65020de3-8a6c-416f-97d2-9755b0111bdc", "/pessoas/id_pessoa/enderecos/id_endereco"},
	}

	for _, tt := range tests {
		out := Parse(tt.method, tt.uri)
		if out.URI != tt.want {
			t.Fatalf("got=%s want=%s", out.URI, tt.want)
		}
	}
}
