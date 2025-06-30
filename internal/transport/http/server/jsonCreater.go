package server

import (
	"encoding/json"
	"time"

	"github.com/kkvaleriy/pingRobot/internal/usecase/status"
)

type serviceStatus struct {
	Name      string    `json:"service_name"`
	Status    string    `json:"service_status"`
	LastCheck time.Time `json:"time_of_last_check"`
}

type servicesStatus struct {
	Services []serviceStatus `json:"services"`
}

func CreateJson(servicesStatusDto []status.Dto) ([]byte, error) {
	var s servicesStatus

	for _, serviceStatusDto := range servicesStatusDto {
		ss := serviceStatus{
			Name:      serviceStatusDto.Name,
			Status:    serviceStatusDto.Status,
			LastCheck: serviceStatusDto.LastCheck,
		}
		s.Services = append(s.Services, ss)
	}

	return json.Marshal(s)
}
