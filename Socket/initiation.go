package Socket

import "github.com/nats-io/nats.go"

var Connection *nats.Conn

func InitateSocket() {
	connection, err := nats.Connect(nats.DefaultURL)

	switch err != nil {
	case true:
		panic(err)
	}

	Connection = connection
}
