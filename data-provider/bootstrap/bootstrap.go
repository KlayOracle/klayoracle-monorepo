package bootstrap

import "os"

// Nodes Bootstrap nodes for coordinating data providers
func Nodes() []string {
	env := os.Getenv("environment")

	if env == "docker" {
		return []string{"bootstrap_dp1:50002", "bootstrap_dp2:50002", "bootstrap_dp3:50002"}
	}

	return []string{"0.0.0.0:50002", "0.0.0.0:50003", "0.0.0.0:50004"}
}
