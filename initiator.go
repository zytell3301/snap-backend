package main

import (
	"snap/Database/Redis"
	"snap/Database/Uuid"
)

func initiatePackages() {
	Redis.InitiateRedis()
	Uuid.InitSpace()
}
