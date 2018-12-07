package helpers

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// BreakOnNewLines ... takes a file and returns a slice of strings
func BreakOnNewLines(fileName string) ([]string, error) {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		var emptySlice []string
		return emptySlice, err
	}

	return strings.Split(strings.TrimSpace(string(file)), "\n"), nil
}

// ParseIntList ... takes a filename and returns a list of integers
func ParseIntList(fileName string) []int64 {
	s, err := BreakOnNewLines(fileName)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Couldn't parse input"))

	}
	var d []int64
	for _, v := range s {
		parsed, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal(errors.Wrap(err, "Could not parse input"))
		}
		d = append(d, parsed)
	}
	return d
}
