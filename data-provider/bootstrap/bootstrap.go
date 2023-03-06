package bootstrap

// Nodes Bootstrap nodes for coordinating data providers
func Nodes() []string {
	return []string{"localhost:50002"}
	//return []string{"localhost:50002", "localhost:50003", "localhost:50004"}
}
