package Keyspaces

var initiators = []func(){initiateSnapKeyspace}

func InitiateKeyspaces() {
	for _, initiator := range initiators {
		initiator()
	}
}
