# internal

## Overview

Holds the Go packages behind the `aeroctl` subcommands. Go's `internal` rule keeps them importable only from within the `aeroctl` module.

## Subdirectories & Boundaries

- dnscheck: [./dnscheck/AGENTS.md](./dnscheck/AGENTS.md) - `Run` checks SPF, DKIM and DMARC TXT records through DNS.
- rotate: [./rotate/AGENTS.md](./rotate/AGENTS.md) - `Run` lists a pool's IPs through core's API and rotates the unhealthy ones.
