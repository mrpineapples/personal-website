package utils

import (
	"fmt"
	"math"
	"regexp"
)

// GetReadingTime calculates the amount of time it should take to read the
// given text and returns it in a formatted string (example: "5 min read")
func GetReadingTime(text string) string {
	wordsPerMin := 250
	// Match one or more word characters
	reg := regexp.MustCompile(`\w+`)
	// Find all matches in the text
	matches := reg.FindAllString(text, -1)
	count := len(matches)
	time := math.Ceil(float64((count / wordsPerMin)))

	s := "<1 min read"
	if time > 0 {
		s = fmt.Sprintf("%v min read", time)
	}

	return s
}
