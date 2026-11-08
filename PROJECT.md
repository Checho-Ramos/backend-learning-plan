# Project brief

Stack: Go 1.2x, PostgreSQL, Redis, Docker Compose.
Commands: `go test ./...`, `go test -race ./...`, `go vet ./...`, `staticcheck ./...`, `docker compose up --build`.
Layout: cmd/api, internal/<domain>, migrations, docs/adr, notes.
Rules: stdlib first; no ORM; SQL by hand; wrap errors; context everywhere; table-driven tests.
Definition of done: runs from clean checkout, tests fail when logic is broken, README updated.
Never: add a dependency without asking, rewrite files I did not ask about, or write code I cannot explain.
