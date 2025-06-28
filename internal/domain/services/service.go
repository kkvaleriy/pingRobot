package services

import "time"

type service struct {
	status    string
	lastCheck time.Time
}
