# dnscheck

## Overview

Backs `aeroctl dns check`. It confirms that a domain publishes SPF, DKIM and DMARC TXT records and prints one status line per record.

## Interface

- `Run(domain, selector string) int` (`dnscheck.go`): checks the three records below and returns a process exit code.

| Check | Host looked up | Required TXT prefix |
|---|---|---|
| SPF | `<domain>` | `v=spf1` |
| DKIM | `<selector>._domainkey.<domain>` | `v=DKIM1` |
| DMARC | `_dmarc.<domain>` | `v=DMARC1` |

## Invariants

- `Run` returns `1` if any check fails and `0` if all pass. It always runs all three checks.
- `hasTXT` tests presence only, with a case-sensitive `strings.HasPrefix`. It does not compare record contents with what core generated.
- A DNS lookup error counts as `MISSING`, so a resolver outage looks the same as an absent record.

## Side Effects

- Sends DNS TXT queries through `net.LookupTXT`.
- Writes one `OK` or `MISSING` line per check to stdout.

## Verification

- Run: `go test ./...` from `repos/internal-tools/aeroctl`
- This package has no `*_test.go` files, so the command runs no tests here.
