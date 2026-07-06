package dnscheck

import (
	"fmt"
	"net"
	"strings"
)

type check struct {
	name   string
	host   string
	prefix string
}

// Run verifies SPF, DKIM and DMARC TXT records exist for the domain. Returns a process exit code.
func Run(domain, selector string) int {
	checks := []check{
		{"SPF", domain, "v=spf1"},
		{"DKIM", selector + "._domainkey." + domain, "v=DKIM1"},
		{"DMARC", "_dmarc." + domain, "v=DMARC1"},
	}
	failed := 0
	for _, c := range checks {
		ok := hasTXT(c.host, c.prefix)
		status := "OK"
		if !ok {
			status = "MISSING"
			failed++
		}
		fmt.Printf("%-6s %-45s %s\n", c.name, c.host, status)
	}
	if failed > 0 {
		return 1
	}
	return 0
}

func hasTXT(host, prefix string) bool {
	records, err := net.LookupTXT(host)
	if err != nil {
		return false
	}
	for _, r := range records {
		if strings.HasPrefix(r, prefix) {
			return true
		}
	}
	return false
}
