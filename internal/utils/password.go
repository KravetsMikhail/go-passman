package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	numbers   = "0123456789"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// GeneratePassword generates a random password
func GeneratePassword(length int, useNumbers bool, useSpecial bool) string {
	charset := lowercase + uppercase

	if useNumbers {
		charset += numbers
	}
	if useSpecial {
		charset += special
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

// ChoosePasswordOptions prompts user for password generation options
func ChoosePasswordOptions() (int, bool, bool) {
	var lengthStr string
	fmt.Print("Enter password length (default 16): ")
	fmt.Scanln(&lengthStr)

	length := 16
	if lengthStr != "" {
		fmt.Sscanf(lengthStr, "%d", &length)
	}

	var useNumbers string
	fmt.Print("Include numbers? (y/n, default y): ")
	fmt.Scanln(&useNumbers)
	useNumStr := strings.ToLower(useNumbers) != "n"

	var useSpecial string
	fmt.Print("Include special characters? (y/n, default y): ")
	fmt.Scanln(&useSpecial)
	useSpecStr := strings.ToLower(useSpecial) != "n"

	return length, useNumStr, useSpecStr
}
