package main

import (
	"snap/Database/Cassandra/Keyspaces"
	"snap/Database/Cassandra/Models"
	"snap/Database/Redis"
	"snap/Database/Uuid"
	"snap/DriversService"
)

func initiatePackages() {
	Redis.InitiateRedis()
	Uuid.InitSpace()
	Keyspaces.InitiateKeyspaces()
	go DriversService.InitiateGrpcServices()
	Models.EvaluateCassandraModelsSession()
}
