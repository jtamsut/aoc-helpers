package helpers

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// WrapError ... handles errors
func WrapError(e error) {
	log.Fatal(errors.Wrap(e, "Couldn't parse input"))
}

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
		WrapError(err)

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

// Sum ... adds up the elements in a list
func Sum(lst ...int64) int64 {
	var sum int64
	for _, val := range lst {
		sum += val
	}
	return sum
}

// LettersToNums ... returns map that maps letter to number
func LettersToNums() map[string]int {
	return map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
		"n": 13,
		"o": 14,
		"p": 15,
		"q": 16,
		"r": 17,
		"s": 18,
		"t": 19,
		"u": 20,
		"v": 21,
		"w": 22,
		"x": 23,
		"y": 24,
		"z": 25,
	}
}
