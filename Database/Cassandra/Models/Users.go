package Models

import (
	"github.com/gocql/gocql"
	"snap/Database/Cassandra"
	"time"
)

type Driver struct {
	Name        string `cql:"name"`
	Lastname    string `cql:"lastname"`
	Vehicle_no  string `cql:"vehicle_no"`
	Balance     int    `cql:"Balance"`
	Profile_pic string `cql:"profile_pic"`
}

type User struct {
	Cassandra.TableMetaData
	Id             gocql.UUID
	Phone          string
	Password       string
	Balance        int
	Driver_details Driver
	Created_at     time.Time
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

func (u User) GetDriverDetails(conditions map[string]interface{}) (user User, err error) {
	statement := u.GetSelectStatement(conditions, []string{"Id", "Driver_details"})
	err = statement.Scan(&user.Id, &user.Driver_details)

	return
}
