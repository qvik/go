package os

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// MustGetenv reads an environment variable value or panics if not found
func MustGetenv(varname string) string {
	value := os.Getenv(varname)
	if value == "" {
		panic(fmt.Errorf("Missing env variable %v", varname))
	}

	return value
}

// MustGetenvFile reads and returns contents of a file pointed to by given
// environment variable as a string, or panics if not found or unable to read
func MustGetenvFile(varname string) string {
	filePath := MustGetenv(varname)

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("Failed to read file: %v", err))
	}

	s := strings.Trim(string(bytes), "\n")

	return s
}

// MustGetenvInt reads an int environment variable value or panics if
// not found or not integer
func MustGetenvInt(varname string) int {
	s := MustGetenv(varname)

	if i, err := strconv.Atoi(s); err != nil {
		panic(fmt.Errorf("Expected int value for env var %v - got: %v",
			varname, s))
	} else {
		return i
	}
}
