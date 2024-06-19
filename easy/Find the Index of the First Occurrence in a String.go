package easy

import "strings"

func StrStr(haystack string, needle string) int {
	//I guess the aim of this was to expose me to the existence of this function in golang.

	return strings.Index(haystack, needle)

}
