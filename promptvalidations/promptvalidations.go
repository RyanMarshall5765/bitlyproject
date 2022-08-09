package promptvalidations

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/RyanMarshall5765/bitlyproject/cfg"
)

func ValidateFile() func(string) error {
	return func(input string) error {
		f, err := os.Open(input)
		if err != nil {
			return fmt.Errorf("invalid file location: %s not found", input)
		}
		defer f.Close()
		return nil
	}
}

func ValidateRange(start, end int) func(string) error {
	return func(input string) error {
		intInput, err := strconv.Atoi(input)
		if intInput < start || intInput > end || err != nil {
			return fmt.Errorf("invalid input! expected number from: %d-%d", start, end)
		}
		return nil
	}
}

func ValidateStringPattern(pattern string) func(string) error {
	return func(input string) error {
		match, err := regexp.MatchString(pattern, input)
		if !match || err != nil {
			return fmt.Errorf("invalid format! expected format: %s", cfg.TimeStampFormat)
		}
		return nil
	}
}
