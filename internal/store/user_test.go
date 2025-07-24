package store

import (
	"crypto/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/jsingh0402/notion-pomodoro/models"
)

func generateTestKeys(t *testing.T) []byte {
	t.Helper()

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}
	return key
}

func getTestUserStore(t *testing.T) *UserStore {
	t.Helper()

	tmpDir := t.TempDir()
	key := generateTestKeys(t)
	return NewUserStore(tmpDir, key)
}

func TestRegisterAndGetUser(t *testing.T) {
	store := getTestUserStore(t)

	user := models.User{
		Username:    "alice",
		NotionToken: "secret-token",
		DatabaseID:  "db-123",
	}

	err := store.RegisterUser(user)
	if err != nil {
		t.Fatalf("RegisterUser failed: %v", err)
	}

	fetched, err := store.GetUser("alice")
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	if fetched.NotionToken != "secret-token" || fetched.DatabaseID != "db-123" {
		t.Errorf("Fetched user does not match original. Got %v, expected %v", fetched, user)
	}
}

func TestOverwriteUser(t *testing.T) {
	store := getTestUserStore(t)

	user := models.User{
		Username:    "bob",
		NotionToken: "token-old",
		DatabaseID:  "db-old",
	}

	_ = store.RegisterUser(user)

	// Overwrite
	user.NotionToken = "token-new"
	user.DatabaseID = "db-new"

	err := store.RegisterUser(user)
	if err != nil {
		t.Fatalf("Overwrite RegisterUser failed: %v", err)
	}

	fetched, err := store.GetUser("bob")
	if err != nil {
		t.Fatalf("GetUser after overwrite failed: %v", err)
	}

	if fetched.NotionToken != "token-new" || fetched.DatabaseID != "db-new" {
		t.Errorf("Overwrite failed. Got %+v", fetched)
	}
}

func TestGetUser_NotFound(t *testing.T) {
	store := getTestUserStore(t)

	_, err := store.GetUser("ghost")
	if err == nil {
		t.Error("Expected error for non-existent user, got nil")
	}
}

func TestLoadUsers_EmptyFile(t *testing.T) {
	store := getTestUserStore(t)

	// Touch empty file
	os.WriteFile(filepath.Join(store.filePath), []byte("[]"), 0644)

	users, err := store.LoadUsers()
	if err != nil {
		t.Fatalf("LoadUsers failed: %v", err)
	}

	if len(users) != 0 {
		t.Errorf("Expected empty user list, got %d", len(users))
	}
}
