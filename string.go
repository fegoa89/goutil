// Package string is a collection of small reusable Go functions
package string

import (
	"regexp"
	"strings"
	"unicode"
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

// Convert string to UpperCamelCase
func UpperCamelCase(inputStr string) (outputString string) {
	inputStr = strings.ToUpper(string(inputStr[0])) + inputStr[1:]
	return toCamelCase(inputStr)
}

// Convert string to lowerCamelCase
func LowerCamelCase(inputStr string) (outputString string) {
	inputStr = strings.ToLower(string(inputStr[0])) + inputStr[1:]
	return toCamelCase(inputStr)
}

// Converts a string to Snake Case
func ToSnakeCase(s string) string {
	return toSeparatedLowerCase(s, "_")
}

// Converts a string to Snake Case
func ToKebabCase(s string) string {
	return toSeparatedLowerCase(s, "-")
}

// Converts a computerized string into a human-friendly one
func Humanize(inputStr string) string {
	inputStr = strings.TrimSpace(inputStr)
	inputStr = regexp.MustCompile("([A-Z])|[_\\s]+|[-\\s]+").ReplaceAllString(inputStr, " $1")

	return strings.ToLower(inputStr)
}

// Truncates a string up to a specified string length
func Truncate(inputStr string, length int) string {
	if len(inputStr) > length {
		if length > 3 {
			length -= 3
		}
		inputStr = inputStr[0:length] + "..."
	}

	return inputStr
}

// Convert a space/dash/dot/underscore separated string to CamelCase
func toCamelCase(inputStr string) (camelCase string) {
	capitalizeNextWord := false
	stringDelimiters := []string{" ", "-", "_", "."}

	for _, v := range inputStr {
		if capitalizeNextWord {
			camelCase += strings.ToUpper(string(v))
			capitalizeNextWord = false
		} else {
			if contains(stringDelimiters, string(v)) {
				capitalizeNextWord = true
			} else {
				camelCase += string(v)
			}
		}
	}
	return
}

// Convert string into separated lower case
func toSeparatedLowerCase(inputString string, delimiter string) (outputString string) {
	inputString = strings.Trim(inputString, " ")
	stringDelimiters := []string{" ", "-", "_", "."}

	for k, v := range inputString {
		if v == unicode.ToUpper(v) && k == 0 {
			outputString += strings.ToLower(string(v))
		} else if contains(stringDelimiters, string(v)) {
			continue
		} else if v == unicode.ToUpper(v) {
			outputString += delimiter
			outputString += strings.ToLower(string(v))
		} else {
			outputString += string(v)
		}
	}
	return
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
