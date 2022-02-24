package helpers

import (
	"math/rand"
	"strconv"
	"time"
)

// RandEntry returns a random entry from a supplied list.
func RandEntry(slice []string) string {
	rand.Seed(time.Now().UnixNano())
	return slice[rand.Intn(len(slice))]
}

// IntegersFromString tapes numbers from a string and puts them together into a int64.
func IntegersFromString(text string) int64 {
	integer := ""
	for _, char := range text {
		switch char {
		case '1':
			integer += "1"
		case '2':
			integer += "2"
		case '3':
			integer += "3"
		case '4':
			integer += "4"
		case '5':
			integer += "5"
		case '6':
			integer += "6"
		case '7':
			integer += "7"
		case '8':
			integer += "8"
		case '9':
			integer += "9"
		case '0':
			integer += "0"
		default:
			continue
		}
	}
	transformed, _ := strconv.Atoi(integer)
	return int64(transformed)
}

// TimeSinceJesus returns how many days have passed since year 0.
func TimeSinceJesus() int {
	return (time.Now().Year() * 356) + time.Now().YearDay()
}
