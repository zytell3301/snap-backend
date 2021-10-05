package Models

import (
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
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

func (u User) NewUser(user User) error {
	usr := map[string]interface{}{}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	switch err != nil {
	case true:
		return err
	}

	usr["id"] = user.Id.String()
	usr["balance"] = 0
	usr["phone"] = user.Phone
	usr["password"] = string(password)
	usr["driver_details"] = user.Driver_details
	usr["created_at"] = time.Now()

	batch := Users.Connection.NewBatch(gocql.LoggedBatch)

	userPkPhone := map[string]interface{}{}
	userPkPhone["phone"] = usr["phone"]
	userPkPhone["id"] = usr["id"]
	userPkPhone["password"] = usr["password"]

	UserPkPhone.NewRecord(userPkPhone, batch)

	err = u.Connection.ExecuteBatch(batch)

	return err
}
