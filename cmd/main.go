package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jsingh0402/notion-pomodoro/models"
	"github.com/jsingh0402/notion-pomodoro/store"
)

func main() {
	//CLI flags
	var (
		register    = flag.Bool("register", false, "Register a new user")
		get         = flag.String("get", "", "Get user info by username")
		username    = flag.String("username", "", "Username")
		notionToken = flag.String("token", "", "Notion API Token")
		notionDB    = flag.String("db", "", "Notion database ID")
	)

	flag.Parse()

	// Load the secret key
	secretKey := os.Getenv("POMODORO_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("Environment variable POMODORO_SECRET_KEY is not set")
	}

	key, err := base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		log.Fatalf("Invalid base64 in POMODORO_SECRET_KEY: %v", err)
	}

	// Setup Store Path
	home, _ := os.UserHomeDir()
	storeDir := filepath.Join(home, ".notion-pomodoro")
	err = os.Mkdir(storeDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create store directory: %v", err)
	}

	userStore := store.NewUserStore(storeDir, key)

	switch {
	case *register:
		if *username == "" || *notionToken == "" || *notionDB == "" {
			log.Fatal("--register requires --username, --token, --db")
		}

		user := models.User{
			Username:    *username,
			NotionToken: *notionToken,
			DatabaseID:  *notionDB,
		}

		err := userStore.RegisterUser(user)
		if err != nil {
			log.Fatalf("RegisterUser failed: %v", err)
		}
		fmt.Println("User registered successfully.")

	case *get != "":
		user, err := userStore.GetUser(*get)
		if err != nil {
			log.Fatalf("GetUser failed: %v", err)
		}

		fmt.Println("Username: ", user.Username)
		fmt.Println("NotionToken: ", user.NotionToken)
		fmt.Println("DatabaseID: ", user.DatabaseID)

	default:
		fmt.Println("Pomodoro CLI")
		flag.Usage()
	}

}
