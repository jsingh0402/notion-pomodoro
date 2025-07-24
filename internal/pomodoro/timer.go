package pomodoro

import (
	"fmt"
	"time"
)

// SessionResult holds the details of one Pomodoro session
type SessionResult struct {
	DurationMinutes int
	StartTime       time.Time
	EndTime         time.Time
}

// StartPomodoro starts a timer for given minutes and returns session result
func StartPomodoro(durationMinutes int) SessionResult {
	fmt.Printf("Starting Pomodoro: %d minutes \n", durationMinutes)
	start := time.Now()
	end := start.Add(time.Duration(durationMinutes) * time.Minute)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for now := range ticker.C {
		remaining := end.Sub(now)
		if remaining <= 0 {
			fmt.Println("\n Pomodoro complete!")
			break
		}
		fmt.Printf("\r⏳ Time remaining: %s", remaining.Truncate(time.Second))
	}

	return SessionResult{
		DurationMinutes: durationMinutes,
		StartTime:       start,
		EndTime:         time.Now(),
	}
}
