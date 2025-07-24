package store

import (
	"crypto/rand"
	"testing"
)

func generateRandomKey(t *testing.T) []byte {
	t.Helper()
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate encryption key: %v", err)
	}
	return key
}

func TestEncryptDecryptBasic(t *testing.T) {
	key := generateRandomKey(t)
	plaintext := "my-secret-data"

	encrypted, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if decrypted != plaintext {
		t.Errorf("Decrypted text mismatch, Expected %s, got %s", plaintext, decrypted)
	}
}

func TestEncryptDecryptEmptyString(t *testing.T) {
	key := generateRandomKey(t)
	plaintext := ""

	encrypted, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if decrypted != plaintext {
		t.Errorf("Decrypted text mismatch, Expected %s, got %s", plaintext, decrypted)
	}
}

func TestDecryptWithWrongKeyFails(t *testing.T) {
	key1 := generateRandomKey(t)
	key2 := generateRandomKey(t)
	plaintext := "sensitive-info"

	encrypted, err := Encrypt(plaintext, key1)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	_, err = Decrypt(encrypted, key2)
	if err == nil {
		t.Error("Expected decryption with wrong key to fail, but it succeeded")
	}
}

func TestDecryptWithCorruptedCiphertextFails(t *testing.T) {
	key := generateRandomKey(t)
	plaintext := "corrupt me"

	encrypted, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	// Corrupt the ciphertext
	raw := []byte(encrypted)
	if len(raw) > 5 {
		raw[5] ^= 0xFF // flip one byte
	}
	corrupted := string(raw)

	_, err = Decrypt(corrupted, key)
	if err == nil {
		t.Error("Expected error on corrupted ciphertext, got none")
	}
}
