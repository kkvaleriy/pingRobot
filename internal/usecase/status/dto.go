package status

import (
	"time"

	"github.com/kkvaleriy/pingRobot/internal/domain/services"
)

type Dto struct {
	Name      string
	Status    string
	LastCheck time.Time
}

func createDto(s map[string]services.Service) []Dto {
	statuses := []Dto{}

	for name, service := range s {
		statuses = append(statuses, Dto{
			Name:      name,
			Status:    service.Status(),
			LastCheck: service.TimeOfLastCheck(),
		})
	}

	return statuses
}
