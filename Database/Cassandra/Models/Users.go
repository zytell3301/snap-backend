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
	Balance        int
	Driver_details Driver
	Created_at     time.Time
}

var Users = User{
	TableMetaData: Cassandra.TableMetaData{
		Table: "users",
		Columns: map[string]struct{}{
			"Id":             {},
			"Phone":          {},
			"Balance":        {},
			"Driver_details": {},
			"Created_at":     {},
		},
		Pk:       map[string]struct{}{"Id": {}},
		Keyspace: "snap",
	},
}

func (u User) GetDriverDetails(conditions map[string]interface{}) (user User,err error) {
	statement := u.GetSelectStatement(conditions, []string{"Id","Driver_details"})
	err = statement.Scan(&user.Id,&user.Driver_details)

	return
}
