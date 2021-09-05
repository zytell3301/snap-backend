package Models

import "snap/Database/Cassandra"

var Tokens = Cassandra.TableMetaData{
	Table: "tokens",
	Columns: map[string]struct{}{
		"communication_token": {},
		"user_id":             {},
	},
	Pk:       map[string]struct{}{"communication_token": {}},
	Keyspace: "snap",
}
