package main

import (
	"snap/Database/Cassandra/Keyspaces"
	"snap/Database/Redis"
	"snap/Database/Uuid"
)

func initiatePackages() {
	Redis.InitiateRedis()
	Uuid.InitSpace()
	Keyspaces.InitiateKeyspaces()
}
