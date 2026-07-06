package rotate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ip struct {
	ID          int64   `json:"id"`
	Address     string  `json:"address"`
	HealthScore float64 `json:"healthScore"`
}

// Run lists IPs in a pool via core's API and rotates the unhealthy ones.
func Run(coreURL, pool string, below float64, dryRun bool) int {
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/ip-pools/%s/ips", coreURL, pool))
	if err != nil {
		fmt.Fprintln(os.Stderr, "list ips:", err)
		return 1
	}
	defer resp.Body.Close()

	var ips []ip
	if err := json.NewDecoder(resp.Body).Decode(&ips); err != nil {
		fmt.Fprintln(os.Stderr, "decode:", err)
		return 1
	}

	for _, i := range ips {
		if i.HealthScore >= below {
			continue
		}
		if dryRun {
			fmt.Printf("would rotate %s (health %.1f)\n", i.Address, i.HealthScore)
			continue
		}
		r, err := http.Post(fmt.Sprintf("%s/api/v1/ips/%d/rotate", coreURL, i.ID), "application/json", nil)
		if err != nil || r.StatusCode >= 300 {
			fmt.Fprintf(os.Stderr, "rotate %s failed: %v\n", i.Address, err)
			return 1
		}
		fmt.Printf("rotated %s\n", i.Address)
	}
	return 0
}
