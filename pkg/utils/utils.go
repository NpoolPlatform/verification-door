package utils

import (
	"regexp"
	"strings"
)

func ValidPhoneNumber(phone string) bool {
	start := strings.HasPrefix(phone, "+")
	if !start {
		return false
	}

	b, err := regexp.MatchString("^[0-9]*$", phone[1:])
	if err != nil || !b {
		return false
	}

	return b
}
