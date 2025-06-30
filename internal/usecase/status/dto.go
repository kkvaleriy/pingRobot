package status

import (
	"time"

	"github.com/kkvaleriy/pingRobot/internal/domain/services"
)

type Dto struct {
	name      string
	status    string
	lastCheck time.Time
}

func createDto(s map[string]services.Service) []Dto {
	statuses := []Dto{}

	for name, service := range s {
		statuses = append(statuses, Dto{
			name:      name,
			status:    service.Status(),
			lastCheck: service.TimeOfLastCheck(),
		})
	}

	return statuses
}
