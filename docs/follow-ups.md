# Follow-ups

## Found 2026-11-11 (older than v0.4.0)

- `go vet`: `internal/infrastructure/gcp/ObjectProvider.go:86` and `:107` discard the cancel function from
  `context.WithTimeout` (a context leak).
- `gofmt -l` lists `test/unit/presentation/api/gcp/object/get/Operation_test.go` (format only).
