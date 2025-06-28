package status

import "time"

type StatusDto struct {
	name      string
	status    string
	lastCheck time.Time
}
