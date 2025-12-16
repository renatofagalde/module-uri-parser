# module-uri-parser

Módulo Go para normalizar URIs, substituindo UUIDs por placeholders baseados no recurso.

## Exemplo

```go
out := uriparser.Parse("GET", "/pessoas/4152bd3d-9121-4889-b683-6ee72b8e3719")
fmt.Println(out.URI)
// /pessoas/id_pessoa
