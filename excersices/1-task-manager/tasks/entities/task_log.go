package entities

import (
	"time"

	"github.com/ilyasa1211/go-exercises/excersices/1-task-manager/tasks/interfaces"
)

type TaskLog struct {
	interfaces.Log
	StartTime time.Time
	EndTime   time.Time
	DeltaTime time.Duration
}
