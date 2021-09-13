package Models

import (
	"github.com/gocql/gocql"
	"snap/Database/Cassandra"
	"time"
)

type Driver struct {
	name        string
	lastname    string
	vehicle_no  string
	balance     int
	profile_pic string
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
