package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// integersFromString tapes numbers from a string and puts them together into a int64.
func integersFromString(text string) int64 {
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

// timeSinceJesus returns how many days have passed since year 0.
func timeSinceJesus() int {
	return (time.Now().Year() * 356) + time.Now().YearDay()
}

// hashSelf finds the path of the bot executable, reads it and returns MD5 hash of it.
func hashSelf() (string, error) {
	hash := ""

	selfPath, err := os.Executable()
	if err != nil {
		fmt.Println(err.Error())
		return hash, err
	}

	contentRaw, err := ioutil.ReadFile(selfPath)
	if err != nil {
		fmt.Println(err.Error())
		return hash, err
	}

	md5Hash := md5.Sum(contentRaw)
	hash = hex.EncodeToString(md5Hash[:])

	return hash, nil
}
