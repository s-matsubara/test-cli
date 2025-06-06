# AGENTS Instructions for test-cli

## Commit Messages
- Use Conventional Commits style: `<type>: <description>`
  - Example: `feat: add custom flag for greet`
  - Types: feat, fix, docs, test, chore, etc.
- Keep the summary imperative, under 50 characters and start in lowercase

## Coding Style
- Format Go code with `go fmt` before committing.
- Run `golangci-lint run ./...` for lint checks.

## Testing
- Run all unit tests via `go test ./... -v`.
- Integration tests are gated by `INTEGRATION_TEST=true`.
- Ensure `mise run go-test-all` passes if you use mise.

## Pull Request Notes
- Summarize important changes in the PR body.
- Include test results or mention if tests could not run.

## Directory Overview
- `cmd/`    – subcommands (`greet`, `version`, etc.)
- `main.go` – entry point calling `cmd.Execute()`
- `mise.toml` – task definitions (build, lint, test)
