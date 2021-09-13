package Models

import (
	"github.com/gocql/gocql"
	"snap/Database/Cassandra"
	"time"
)

type Driver struct {
	Name        string `cql:"name"`
	Lastname    string `cql:"lastname"`
	Vehicle_no  string `cql:vehicle_no`
	Balance     int    `cql:"balance"`
	Profile_pic string `cql:"profile_pic`
}

type User struct {
	Cassandra.TableMetaData
	id             gocql.UUID
	phone          string
	balance        int
	driver_details Driver
	created_at     time.Time
}

var Users = User{
	TableMetaData: Cassandra.TableMetaData{
		Table: "users",
		Columns: map[string]struct{}{
			"id":             {},
			"phone":          {},
			"balance":        {},
			"driver_details": {},
			"created_at":     {},
		},
		Pk:       map[string]struct{}{"id": {}},
		Keyspace: "snap",
	},
}
