package notion

import (
	"log"
	"os"
)

// Config holds Notion API credentials and config values
type Config struct {
	NotionToken string
	DatabaseID  string
}

// LoadConfig reads credentials from environment variables
func LoadConfig() Config {
	notionToken := os.Getenv("NOTION_TOKEN")
	databaseID := os.Getenv("NOTION_DATABASE_ID")

	if notionToken == "" || databaseID == "" {
		log.Fatal("Environment variables NOTION_TOKEN and NOTION_DATABASE_ID must be set")
	}

	return Config{
		NotionToken: notionToken,
		DatabaseID:  databaseID,
	}
}
