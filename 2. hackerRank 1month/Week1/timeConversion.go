package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here

	letter := rune(s[8])
	hour, _ := strconv.Atoi(string(s[0:2]))

	if letter == 'A' && hour != 12 {
		return string(s[0:8])
	} else if hour == 12 && letter == 'P' {
		return string(s[0:8])
	} else if hour == 12 && letter == 'A' {
		return "00" + string(s[2:8])
	}

	hour += 12
	hourString := strconv.Itoa(hour)
	return hourString + string(s[2:8])

}
