package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoadEnv() {
	isProd := gin.Mode() == "release"
	envFilename := ".env.dev"
	if isProd {
		envFilename = ".env.production"
	}

	envFile, err := os.Open(envFilename)
	if err != nil {
		fmt.Println("Could not open env file:", err)
		return
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '#' {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := parts[1]
				os.Setenv(key, value)
			}
		}
	}
}
