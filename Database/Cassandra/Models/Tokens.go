package Models

import (
	"snap/Database/Cassandra"
)

type Token struct {
	Cassandra.TableMetaData
}

var Tokens = Token{
	Cassandra.TableMetaData{
		Table: "tokens",
		Columns: map[string]struct{}{
			"communication_token": {},
			"user_id":             {},
		},
		Pk:       map[string]struct{}{"communication_token": {}},
		Keyspace: "snap"},
}
