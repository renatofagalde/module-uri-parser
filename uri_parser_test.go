package uriparser

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		method string
		uri    string
		want   string
	}{
		{"GET", "/pessoas/4152bd3d-9121-4889-b683-6ee72b8e3719", "/pessoas/{pessoa_id}"},
		{"POST", "/pessoas/4152bd3d-9121-4889-b683-6ee72b8e3719/telefones", "/pessoas/{pessoa_id}/telefones"},
		{"GET", "/pessoas/4152bd3d-9121-4889-b683-6ee72b8e3719/enderecos/65020de3-8a6c-416f-97d2-9755b0111bdc", "/pessoas/{pessoa_id}/enderecos/{endereco_id}"},
		{"GET", "/companies/4152bd3d-9121-4889-b683-6ee72b8e3719/categories/65020de3-8a6c-416f-97d2-9755b0111bdc", "/companies/{company_id}/categories/{category_id}"},
		{"GET", "/cam/users/4152bd3d-9121-4889-b683-6ee72b8e3719/profiles", "/cam/users/{user_id}/profiles"},
		{"GET", "/cam/companies/4152bd3d-9121-4889-b683-6ee72b8e3719/modules", "/cam/companies/{company_id}/modules"},
		{"PATCH", "/cms/apolices/4152bd3d-9121-4889-b683-6ee72b8e3719/lives", "/cms/apolices/{apolice_id}/lives"},
	}

	for _, tt := range tests {
		out := Parse(tt.method, tt.uri)
		if out.URI != tt.want {
			t.Fatalf("got=%s want=%s", out.URI, tt.want)
		}
	}
}
