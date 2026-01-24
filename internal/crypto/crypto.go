package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

const (
	pbkdf2Iterations = 100_000
	saltLen          = 16
	nonceLen         = 12
	keyLen           = 32
)

// deriveKeyFromPassword derives an encryption key from a password
func deriveKeyFromPassword(password string, salt []byte) [keyLen]byte {
	var key [keyLen]byte
	derivedKey := pbkdf2.Key(
		[]byte(password),
		salt,
		pbkdf2Iterations,
		keyLen,
		sha256.New,
	)
	copy(key[:], derivedKey)
	return key
}

// Encrypt encrypts plaintext with the given password
func Encrypt(password string, plaintext []byte) (string, error) {
	// Generate random salt
	salt := make([]byte, saltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	// Derive key from password
	key := deriveKeyFromPassword(password, salt)

	// Create cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Generate random nonce
	nonce := make([]byte, nonceLen)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	// Combine salt + nonce + ciphertext
	data := append(salt, nonce...)
	data = append(data, ciphertext...)

	// Encode to base64
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}

// Decrypt decrypts ciphertext with the given password
func Decrypt(password string, dataB64 string) ([]byte, error) {
	// Decode from base64
	data, err := base64.StdEncoding.DecodeString(dataB64)
	if err != nil {
		return nil, fmt.Errorf("base64 decode error: %w", err)
	}

	// Check minimum length
	if len(data) < saltLen+nonceLen {
		return nil, fmt.Errorf("data too short")
	}

	// Extract salt, nonce, ciphertext
	salt := data[:saltLen]
	nonceBytes := data[saltLen : saltLen+nonceLen]
	ciphertext := data[saltLen+nonceLen:]

	// Derive key
	key := deriveKeyFromPassword(password, salt)

	// Create cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// Decrypt
	plaintext, err := gcm.Open(nil, nonceBytes, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed. incorrect password or corrupted data")
	}

	return plaintext, nil
}
