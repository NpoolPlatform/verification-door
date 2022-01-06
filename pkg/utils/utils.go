package utils

import (
	"regexp"
	"strings"
)

func ValidPhoneNumber(phone string) bool {
	if len(phone) > 14 || len(phone) < 6 {
		return false
	}

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
