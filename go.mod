module github.com/muzyk0/go-shortener-links

go 1.22.10

require (
	github.com/go-chi/chi/v5 v5.2.0
	github.com/stretchr/testify v1.10.0
	go-shorener-links/config v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go-shorener-links/config => ./cmd/shortener/config
