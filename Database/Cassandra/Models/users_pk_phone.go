package Models

import (
	"fmt"
	"github.com/gocql/gocql"
	"snap/Database/Cassandra"
)

type UsersPkPhone struct {
	Cassandra.TableMetaData
	Id       gocql.UUID
	Phone    string
	Password string
}

var UserPkPhone = UsersPkPhone{
	TableMetaData: Cassandra.TableMetaData{
		Table: "users_pk_phone",
		Columns: map[string]struct{}{
			"id":       {},
			"phone":    {},
			"password": {},
		},
		Keyspace: "snap",
		Pk:       map[string]struct{}{"phone": {}},
	},
}

func (p UsersPkPhone) GetUser(phone string) UsersPkPhone {
	statement := p.GetSelectStatement(map[string]interface{}{"phone": phone}, []string{"*"})
	user := map[string]interface{}{}
	err := statement.MapScan(user)

	switch err != nil {
	case true:
		fmt.Printf("An error occurred while fetching user. Error: %v", err)
		return UsersPkPhone{}
	}

	p.Id = user["id"].(gocql.UUID)
	p.Phone = user["phone"].(string)
	p.Password = user["password"].(string)

	return p
}
