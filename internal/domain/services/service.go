package services

import "time"

type Service struct {
	status    string
	lastCheck time.Time
}

func (s *Service) Status() string {
	return s.status
}

func (s *Service) TimeOfLastCheck() time.Time {
	return s.lastCheck
}
