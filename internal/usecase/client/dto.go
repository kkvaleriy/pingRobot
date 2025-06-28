package client

import "time"

type ClientDto struct {
	name        string
	status      string
	timeOfCheck time.Time
}
