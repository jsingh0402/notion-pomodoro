package config

import (
	"errors"
	"os"
)

// Config holds Notion API credentials and config values
type NotionConfig struct {
	Token      string
	DatabaseID string
}

type PomodoroConfig struct {
	DefaultDuration int // in minutes
	EnableLogging   bool
}

type Config struct {
	Notion   NotionConfig
	Pomodoro PomodoroConfig
}

// LoadConfig reads credentials from environment variables
func LoadConfig() (*Config, error) {
	notionToken := os.Getenv("NOTION_TOKEN")
	databaseID := os.Getenv("NOTION_DATABASE_ID")

	if notionToken == "" || databaseID == "" {
		return nil, errors.New("environment variables NOTION_TOKEN and NOTION_DATABASE_ID must be set")
	}

	return &Config{
		NotionToken: notionToken,
		DatabaseID:  databaseID,
	}, nil
}
