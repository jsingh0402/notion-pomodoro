package models

import "time"

// Session Represents a completed Pomodoro session to be logged
type Session struct {
	Username  string
	StartedAt time.Time
	EndedAt   time.Time
	TaskName  string
	Notes     string
}
