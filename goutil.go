// Package goutil is a collection of small reusable Go functions 
package goutil

import(
	"unicode"
	"strings"
)

// Capitalizes the first letter of a string.
func Capitalize(str string) string {
    for i, v := range str {
        return string(unicode.ToUpper(v)) + str[i+1:]
    }
    return ""
}

// Capitalizes string excluding words such as "and", "a", "with", "or".
// You can provide an optional string slice with words that will be excluded too
func CapitalizeTitle(input string, omitWords ...[]string) string {
	title := strings.Fields(input)
	lowercaseWords := []string{"and", "an", "a", "on", "with", "or"}

	if len(omitWords) > 0 {
		for _, word := range omitWords[0] {
			lowercaseWords = append(lowercaseWords, word)
		}
	}

	for index, word := range title {
		if contains(lowercaseWords, word) {
			title[index] = word
		} else {
			title[index] = strings.Title(word)
		}
	}
	return strings.Join(title, " ")
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}