package status

import "github.com/kkvaleriy/pingRobot/internal/domain/services"

func Check() []Dto {
	return createDto(services.Status())
}
