package stringutil

import "strings"

// InArrayStrings search string key on array
func InArrayStrings(s []string, c string) bool {
	return strings.Contains(strings.Join(s, "."), c)
}
