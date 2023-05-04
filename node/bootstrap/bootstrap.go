package bootstrap

import "os"

type BT struct {
	OrgID  string
	Addr   string
	Domain string
}

// Nodes Bootstrap nodes for coordinating Node providers
func Nodes() []BT {
	env := os.Getenv("environment")

	if env == "docker" {
		return []BT{
			BT{
				"535797165684",
				"bootstrap_node1:50052",
				"node-1.origineum.com",
			},
			BT{
				"535797165685",
				"bootstrap_node2:50052",
				"node-2.origineum.com",
			},
			BT{
				"535797165686",
				"bootstrap_node3:50052",
				"node-3.origineum.com",
			},
		}
	}

	return []BT{
		BT{
			"535797165684",
			"0.0.0.0:50051",
			"node-1.origineum.com",
		},
		//BT{
		//	"535797165685",
		//	"0.0.0.0:50052",
		//	"node-2.origineum.com",
		//},
		//BT{
		//	"535797165686",
		//	"0.0.0.0:50053",
		//	"node-3.origineum.com",
		//},
	}
}
