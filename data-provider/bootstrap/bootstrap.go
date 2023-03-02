package bootstrap

// Nodes Bootstrap nodes for coordinating data providers
func Nodes() []string {
	return []string{"localhost:50001", "localhost:50003", "localhost:50004"}
}
