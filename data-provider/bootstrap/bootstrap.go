package bootstrap

import "os"

// Nodes Bootstrap nodes for coordinating data providers
func Nodes() []string {
	env := os.Getenv("environment")

	if env == "staging" {
		return []string{"bootstrap-dp1.digioracle.link:50002", "bootstrap-dp2.digioracle.link:50002", "bootstrap-dp3.digioracle.link:50002"}
	}

	if env == "docker" {
		return []string{"bootstrap_dp1:50002", "bootstrap_dp2:50002", "bootstrap_dp3:50002"}
	}

	return []string{"0.0.0.0:50002", "0.0.0.0:50003", "0.0.0.0:50004"}
}
