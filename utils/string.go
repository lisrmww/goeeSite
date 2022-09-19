package utils

import (
	"strings"
)

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	if s == "id" {
		return "ID"
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	if s == "ID" {
		return "id"
	}
	return strings.ToLower(s[:1]) + s[1:]
}
