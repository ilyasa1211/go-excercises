package entities

import "time"

type BaseTask struct {
	ID          string
	CreatedAt   time.Time
	Description string
}
