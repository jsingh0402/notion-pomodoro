package utils

import (
	"fmt"
	"time"
)

// Info prints an informational message with a timestamp
func Info(message string) {
	fmt.Printf("[INFO] %s — %s\n", time.Now().Format(time.RFC822), message)
}

// Error prints an error message with a timestamp
func Error(message string) {
	fmt.Printf("[ERROR] %s — %s\n", time.Now().Format(time.RFC822), message)
}
