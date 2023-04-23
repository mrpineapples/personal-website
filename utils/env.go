package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv() {
	envFile, err := os.Open(".env")
	if err != nil {
		fmt.Println("Error opening .env file:", err)
		return
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '#' {
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				os.Setenv(parts[0], parts[1])
			}
		}
	}
}
