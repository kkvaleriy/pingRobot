package client

import "github.com/kkvaleriy/pingRobot/internal/domain/services"

func Set(serviceInfo Dto) {
	services.Set(serviceInfo.Name, serviceInfo.Status, serviceInfo.TimeOfCheck)
}
