package Models

import (
	"github.com/gocql/gocql"
	"snap/Database/Cassandra"
)

type Token struct {
	Cassandra.TableMetaData
	communication_token string
	user_id             gocql.UUID
}

var Tokens = Token{
	TableMetaData: Cassandra.TableMetaData{
		Table: "tokens",
		Columns: map[string]struct{}{
			"communication_token": {},
			"user_id":             {},
		},
		Pk:       map[string]struct{}{"communication_token": {}},
		Keyspace: "snap"},
}
