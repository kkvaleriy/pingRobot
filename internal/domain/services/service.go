package services

import "time"

type Service struct {
	status    string
	lastCheck time.Time
}
