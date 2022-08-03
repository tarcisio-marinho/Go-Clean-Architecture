package shared

import "strings"

func IsNullOrEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
