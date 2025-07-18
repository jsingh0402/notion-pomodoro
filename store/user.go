package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/jsingh0402/notion-pomodoro/models"
)

// UserStore handles secure persistence of users.
type UserStore struct {
	filePath string
	key      []byte //encryption key (32 bytes for AES-256)
}

// NewUserStore create a new UserStore using the given directory and encryption key.
func NewUserStore(baseDir string, encryptionkey []byte) *UserStore {
	storePath := filepath.Join(baseDir, "user.json")
	return &UserStore{
		filePath: storePath,
		key:      encryptionkey,
	}
}

// RegisterUser saves or overwrites a user securely
func (us *UserStore) RegisterUser(user models.User) error {
	users, err := us.LoadUsers()
	if err != nil {
		return err
	}

	//Encrypt sensitive fields
	encryptedToken, err := Encrypt(user.NotionToken, us.key)
	if err != nil {
		return err
	}

	encryptedDBID, err := Encrypt(user.DatabaseID, us.key)
	if err != nil {
		return err
	}

	found := false
	for i, u := range users {
		if u.Username == user.Username {
			users[i].NotionToken = encryptedToken
			users[i].DatabaseID = encryptedDBID
			found = true
			break
		}
	}

	if !found {
		users = append(users, models.User{
			Username:    user.Username,
			NotionToken: encryptedToken,
			DatabaseID:  encryptedDBID,
		})
	}

	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(us.filePath, data, 0644)
}

// GetUser fetches and decrypts a user by username
func (us *UserStore) GetUser(username string) (models.User, error) {
	users, err := us.LoadUsers()
	if err != nil {
		return models.User{}, err
	}

	for _, u := range users {
		if u.Username == username {
			token, err := Decrypt(u.DatabaseID, us.key)
			if err != nil {
				return models.User{}, err
			}
			dbID, err := Decrypt(u.DatabaseID, us.key)
			if err != nil {
				return models.User{}, err
			}

			u.NotionToken = token
			u.DatabaseID = dbID
			return u, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

// LoadUsers loads users without decrypting them
func (us *UserStore) LoadUsers() ([]models.User, error) {
	data, err := os.ReadFile(us.filePath)
	if os.IsNotExist(err) {
		return []models.User{}, nil
	}
	if err != nil {
		return nil, err
	}

	var users []models.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
