package client

import "time"

type Dto struct {
	Name        string
	Status      string
	TimeOfCheck time.Time
}
