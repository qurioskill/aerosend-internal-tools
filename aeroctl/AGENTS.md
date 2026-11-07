# aeroctl

## Overview

The Aerosend ops CLI, in Go with no third-party dependencies. `main.go` parses the two-word subcommand and its flags, then hands off to a package under `internal/`, whose `Run` result becomes the process exit code.

## Interface

- Module: `github.com/aerosend/internal-tools/aeroctl` (`go.mod`, Go 1.22).
- Subpackages: [./internal/AGENTS.md](./internal/AGENTS.md).

| Command | Flags | Handler |
|---|---|---|
| `aeroctl dns check <domain>` | `--selector` (default `aero1`) | `dnscheck.Run` |
| `aeroctl ip rotate` | `--pool` (required), `--below` (default `70`), `--dry-run`, `--core` | `rotate.Run` |

- `--core` defaults to the `AEROSEND_CORE_URL` environment variable, then to `http://localhost:8081` (`envOr`).

## Invariants

- Exit codes: `0` for success, `1` for a failed check or rotation, `2` for usage errors (`usage`, `flag.ExitOnError`).
- Flags must come before the domain: `aeroctl dns check --selector aero1 getacme.com`. Go's `flag` package stops parsing at the first positional argument, so the order shown in the `usage` text, `<domain> [--selector aero1]`, fails the `fs.NArg() != 1` check and exits `2`.
- An empty `--pool` exits through `usage`.
- `main` only parses flags and dispatches through its `switch`.

## Side Effects

- `main.go` reads `AEROSEND_CORE_URL` and otherwise only delegates. DNS and HTTP effects are in the [internal packages](./internal/AGENTS.md).
- `go build -o aeroctl .` writes the `aeroctl` binary, which `.gitignore` excludes.

## Verification

- Run: `go test ./...` from `repos/internal-tools/aeroctl`
- The module has no `*_test.go` files, so the command runs no tests.
