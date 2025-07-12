package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const version = "ZYRO 1.0.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zyro [command]")
		return
	}

	command := os.Args[1]

	switch command {
	case "pass":
		length := 12 // default
		if len(os.Args) >= 3 {
			l, err := strconv.Atoi(os.Args[2])
			if err == nil && l > 0 {
				length = l
			}
		}
		fmt.Println(generatePassword(length))

	case "-v", "--version":
		fmt.Println(version)

	case "help":
		printHelp()

	case "hash":
		if len(os.Args) < 3 {
			fmt.Println("Usage: zyro hash [text]")
			return
		}
		text := os.Args[2]
		fmt.Println(sha256Hash(text))

	case "bcrypt":
		if len(os.Args) < 3 {
			fmt.Println("Usage: zyro bcrypt [text]")
			return
		}
		text := os.Args[2]
		hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(hash))

	default:
		fmt.Println("Unknown command:", command)
	}
}

func generatePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	rand.Seed(time.Now().UnixNano())
	pass := make([]byte, length)
	for i := range pass {
		pass[i] = charset[rand.Intn(len(charset))]
	}
	return string(pass)
}

func printHelp() {
	fmt.Println(`Zyro CLI - Password and Utility Tool
Usage:
  zyro pass [length]       Generate a random password (default: 12 chars)
  zyro hash [text]         Generate SHA256 hash of input
  zyro bcrypt [text]       Generate bcrypt hash of input
  zyro -v | --version      Show current version
  zyro help                Show this help message`)
}

func sha256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
