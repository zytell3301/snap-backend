package Models

import (
	"errors"
	"github.com/gocql/gocql"
	"snap/Database/Cassandra"
)

type Token struct {
	Cassandra.TableMetaData
	Communication_token string
	User_id             gocql.UUID
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

func (t Token) GetToken(communicationToken string) (token *Token, err error) {
	statement := t.GetSelectStatement(map[string]interface{}{"communication_token": communicationToken}, []string{"*"})
	token = &Token{}

	switch statement == nil {
	case true:
		return token, errors.New("an error occurred while creating query statement. Probable wrong argument supplied")
	}

	statement.Scan(&token.communication_token, &token.user_id)

	return
}
