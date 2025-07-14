package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jsingh0402/notion-pomodoro/notion"
	"github.com/jsingh0402/notion-pomodoro/pomodoro"
	"github.com/jsingh0402/notion-pomodoro/utils"
)

func main() {
	utils.Info("Welcome to Notion Pomodoro!")

	// Set session duration (default: 25 min)
	sessionDuration := 25
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter session duration in minutes (press Enter to use 25): ")
	durationInput, _ := reader.ReadString('\n')
	durationInput = strings.TrimSpace(durationInput)
	if durationInput != "" {
		if val, err := strconv.Atoi(durationInput); err == nil {
			sessionDuration = val
		}
	}

	utils.Info(fmt.Sprintf("Starting a %d-minute Pomodoro...", sessionDuration))

	pomodoro.StartPomodoro(sessionDuration)

	utils.Info("Pomodoro session completed!")

	// Ask for log info
	fmt.Print("Enter what you worked on this session: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	fmt.Print("Enter category (DSA, Project, SD, etc.): ")
	topic, _ := reader.ReadString('\n')
	topic = strings.TrimSpace(topic)

	// Log to Notion
	err := notion.AddEntryToNotion(task, topic, sessionDuration)
	if err != nil {
		utils.Error("Failed to log to Notion: " + err.Error())
	} else {
		utils.Info("Successfully logged session to Notion.")
	}
}
