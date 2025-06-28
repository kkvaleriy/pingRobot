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
