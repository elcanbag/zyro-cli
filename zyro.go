// main.go
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zyro pass [length]")
		return
	}

	command := os.Args[1]

	if command == "pass" {
		length := 12 // default
		if len(os.Args) >= 3 {
			l, err := strconv.Atoi(os.Args[2])
			if err == nil && l > 0 {
				length = l
			}
		}
		fmt.Println(generatePassword(length))
	} else {
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
