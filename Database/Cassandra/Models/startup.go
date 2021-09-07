package Models

import "snap/Database/Cassandra"

func EvaluateCassandraModelsSession() {
	Tokens.Connection = Cassandra.ConnectionManager.GetSession(Tokens.Keyspace)
	Users.Connection = Cassandra.ConnectionManager.GetSession(Users.Keyspace)
}
