package client

import "github.com/kkvaleriy/pingRobot/internal/domain/services"

func Set(serviceInfo ClientDto) {
	services.Set(serviceInfo.name, serviceInfo.status, serviceInfo.timeOfCheck)
}
