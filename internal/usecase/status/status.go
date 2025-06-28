package status

import "github.com/kkvaleriy/pingRobot/internal/domain/services"

func Check() []StatusDto {
	return createDto(services.Status())
}
