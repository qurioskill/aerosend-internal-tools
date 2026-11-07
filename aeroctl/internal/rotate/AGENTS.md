# rotate

## Overview

Backs `aeroctl ip rotate`. It lists the IPs in a pool through core's REST API and asks core to rotate each one below a health threshold.

## Interface

- `Run(coreURL, pool string, below float64, dryRun bool) int` (`rotate.go`): returns a process exit code.
- Core API calls, both consumed from `repos/core`:
  - `GET {coreURL}/api/v1/ip-pools/{pool}/ips`, decoded as a JSON array of `ip`.
  - `POST {coreURL}/api/v1/ips/{id}/rotate`, sent with no body.
- `ip` (unexported) maps the JSON fields `id`, `address` and `healthScore`.

## Invariants

- An IP is rotated only when `HealthScore < below`. A score equal to `below` is kept.
- `dryRun` skips every `POST`, but the list `GET` still runs.
- The first failed rotation returns `1` and stops the loop. IPs already rotated stay rotated, because nothing is rolled back.
- A rotate response with status `>= 300` counts as a failure. The error line then prints `<nil>`, because `err` is unset in that case.
- The list response status is not checked. Only a transport or JSON decode error returns `1`.
- Both requests use the default `http` client, so they have no timeout. `pool` is not URL-escaped.
- As of this writing, `repos/core` defines only `DomainController` (`/api/v1/domains`), so neither endpoint above exists in that checkout. `repos/admin` calls the same rotate path from `CoreAdminClient`.

## Side Effects

- Sends HTTP requests to core. Each successful `POST` rotates a live IP.
- Writes progress to stdout and failures to stderr.

## Verification

- Run: `go test ./...` from `repos/internal-tools/aeroctl`
- This package has no `*_test.go` files, so the command runs no tests here.
