package status

import (
	"time"

	"github.com/kkvaleriy/pingRobot/internal/domain/services"
)

type StatusDto struct {
	name      string
	status    string
	lastCheck time.Time
}

func createDto(s map[string]services.Service) []StatusDto {
	statuses := []StatusDto{}

	for name, service := range s {
		statuses = append(statuses, StatusDto{
			name:      name,
			status:    service.Status(),
			lastCheck: service.TimeOfLastCheck(),
		})
	}

	return statuses
}
