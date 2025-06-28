package services

import "time"

type Service struct {
	status    string
	lastCheck time.Time
}

func (s *Service) Status() string {
	return s.status
}

