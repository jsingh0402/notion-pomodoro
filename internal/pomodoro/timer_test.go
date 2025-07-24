package pomodoro

import (
	"testing"
)

func TestStartPomodoro(t *testing.T) {
	duration := 1

	result := StartPomodoro(duration)

	if result.DurationMinutes != duration {
		t.Errorf("Expected duration %d, got %d", duration, result.DurationMinutes)
	}
}
