package Models

import "snap/Database/Cassandra"

var Users = Cassandra.TableMetaData{
	Table: "users",
	Columns: map[string]struct{}{
		"id":             {},
		"phone":          {},
		"balance":        {},
		"driver_details": {},
		"created_at":     {},
	},
	Pk:         map[string]struct{}{"id": {}},
	Keyspace:   "snap",
}
