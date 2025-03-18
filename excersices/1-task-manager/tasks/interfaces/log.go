package interfaces

import "time"

type Log struct {
	Type        string
	Status      string
	Description string
	Timestamp   time.Time
}
