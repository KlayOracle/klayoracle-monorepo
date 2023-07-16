package bootstrap

import "os"

type BT struct {
	OrgID  string
	Addr   string
	Domain string
	Name   string
}

// Nodes Bootstrap nodes for coordinating Node providers
func Nodes() []BT {
	env := os.Getenv("environment")

	if env == "staging" {
		return []BT{
			BT{
				"535797165684",
				"bootstrap-node1.digioracle.link:50051",
				"bootstrap-node-1.digioracle.link",
				"DigiOracle Inc. bt-node-1",
			},
			BT{
				"535797165685",
				"bootstrap-node2.digioracle.link:50051",
				"bootstrap-node-2.digioracle.link",
				"DigiOracle Inc. bt-node-2",
			},
			BT{
				"535797165686",
				"bootstrap-node3.digioracle.link:50051",
				"bootstrap-node-3.digioracle.link",
				"DigiOracle Inc. bt-node-3",
			},
		}
	}

	if env == "docker" {
		return []BT{
			BT{

				"535797165684",
				"bootstrap_node1:50051",
				"node-1.origineum.com",
				"KlayOracle Inc. bt-node-1",
			},
			BT{
				"535797165685",
				"bootstrap_node2:50051",
				"node-2.origineum.com",
				"KlayOracle Inc. bt-node-2",
			},
			BT{
				"535797165686",
				"bootstrap_node3:50051",
				"node-3.origineum.com",
				"KlayOracle Inc. bt-node-3",
			},
		}
	}

	return []BT{
		BT{
			"535797165684",
			"0.0.0.0:50051",
			"node-1.origineum.com",
			"KlayOracle Inc. bt-node-1",
		},
		BT{
			"535797165685",
			"0.0.0.0:50052",
			"node-2.origineum.com",
			"KlayOracle Inc. bt-node-2",
		},
		BT{
			"535797165686",
			"0.0.0.0:50053",
			"node-3.origineum.com",
			"KlayOracle Inc. bt-node-3",
		},
	}
}
