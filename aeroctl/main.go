package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aerosend/internal-tools/aeroctl/internal/dnscheck"
	"github.com/aerosend/internal-tools/aeroctl/internal/rotate"
)

func usage() {
	fmt.Fprintln(os.Stderr, `aeroctl - Aerosend ops CLI

Usage:
  aeroctl dns check <domain> [--selector aero1]
  aeroctl ip rotate --pool <name> [--below 70] [--dry-run]`)
	os.Exit(2)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	switch os.Args[1] + " " + os.Args[2] {
	case "dns check":
		fs := flag.NewFlagSet("dns check", flag.ExitOnError)
		selector := fs.String("selector", "aero1", "DKIM selector")
		fs.Parse(os.Args[3:])
		if fs.NArg() != 1 {
			usage()
		}
		os.Exit(dnscheck.Run(fs.Arg(0), *selector))
	case "ip rotate":
		fs := flag.NewFlagSet("ip rotate", flag.ExitOnError)
		pool := fs.String("pool", "", "IP pool name")
		below := fs.Float64("below", 70, "rotate IPs with health score below this")
		dry := fs.Bool("dry-run", false, "print what would be rotated")
		coreURL := fs.String("core", envOr("AEROSEND_CORE_URL", "http://localhost:8081"), "core API base URL")
		fs.Parse(os.Args[3:])
		if *pool == "" {
			usage()
		}
		os.Exit(rotate.Run(*coreURL, *pool, *below, *dry))
	default:
		usage()
	}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
