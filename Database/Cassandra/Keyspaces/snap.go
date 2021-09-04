package Keyspaces

import (
	"github.com/gocql/gocql"
	"snap/Database/Cassandra"
	"time"
)

var connection = Cassandra.Connection{}

func initiateSnapKeyspace() {
	connection.Cluster = gocql.NewCluster(Cassandra.Configs.GetString("host"))
	connection.Cluster.Consistency = gocql.One
	connection.Cluster.Timeout = 1 * time.Second
	connection.Cluster.Keyspace = "snap"
	connection.Cluster.Port = Cassandra.Configs.GetInt("port")
	connection.Session, _ = connection.Cluster.CreateSession()
	Cassandra.ConnectionManager.AddSession("snap", connection)
}
