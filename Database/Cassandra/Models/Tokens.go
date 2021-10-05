package Models

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
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

	statement.Scan(&token.Communication_token, &token.User_id)

	return
}

func (t Token) GenerateToken(userId gocql.UUID) string {
	token, _ := bcrypt.GenerateFromPassword([]byte(userId.String()), bcrypt.DefaultCost)
	batch := t.Connection.NewBatch(gocql.LoggedBatch)
	t.NewRecord(map[string]interface{}{"user_id": userId.String(), "communication_token": string(token)}, batch)
	err := t.Connection.ExecuteBatch(batch)

	switch err != nil {
	case true:
		fmt.Printf("An error occurred while inserting new token. Error: %v", err)
		return ""
	}

	return string(token)
}
