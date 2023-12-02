package helpers

import "strings"

func PrefixExistsInList(sub string, list []string) bool {
	for _, str := range list {
		if strings.HasPrefix(str, sub) {
			return true
		}
	}
	return false
}
