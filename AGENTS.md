# internal-tools

## Overview

Tools the engineering and ops team uses day to day, none of them customer-facing. The two projects are independent: a Go CLI and a Vue dashboard, with separate manifests and no shared code. Both depend on core's REST API under `/api/v1`, which `repos/core` owns.

## Subdirectories & Boundaries

- aeroctl: [./aeroctl/AGENTS.md](./aeroctl/AGENTS.md) - Go CLI with `dns check` (SPF, DKIM and DMARC lookups) and `ip rotate` (rotates unhealthy IPs through core's API).
- ops-dashboard: [./ops-dashboard/AGENTS.md](./ops-dashboard/AGENTS.md) - Read-only Vue 3 app showing RabbitMQ queue health and IP pool health.

## Verification

- Run: `go test ./...` from `aeroctl`
- Run: `npm run build` from `ops-dashboard`, which has no test script. The build currently fails; see [its contract](./ops-dashboard/AGENTS.md).
