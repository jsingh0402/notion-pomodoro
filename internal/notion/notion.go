package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jsingh0402/notion-pomodoro/config"
)

type NotionRequestBody struct {
	Parent     map[string]string      `json:"parent"`
	Properties map[string]interface{} `json:"properties"`
}

func AddEntryToNotion(task string, topic string, duration int) error {
	url := "https://api.notion.com/v1/pages"

	// Load configuration
	c := config.LoadConfig()
	data := NotionRequestBody{
		Parent: map[string]string{
			"database_id": c.DatabaseID,
		},
		Properties: map[string]interface{}{
			"Session Name": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"text": map[string]string{
							"content": task,
						},
					},
				},
			},
			"Date": map[string]interface{}{
				"date": map[string]string{
					"start": time.Now().Format(time.RFC3339),
				},
			},
			"Topic/Area": map[string]interface{}{
				"multi_select": []map[string]string{
					{"name": topic},
				},
			},
			"Duration (Minutes)": map[string]interface{}{
				"number": duration,
			},
		},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.NotionToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return fmt.Errorf("notion API error: status code %d", res.StatusCode)
	}

	return nil
}
