// Package string is a collection of small reusable Go functions
package string

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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

// Converts a string to Kebab Case
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

// Returns true if string is empty or blank
func IsEmptyOrBlank(inputStr string) bool {
	return IsEmpty(inputStr) || IsBlank(inputStr)
}

// Returns true if string is blank
func IsBlank(inputStr string) bool {
	return len(strings.TrimSpace(inputStr)) == 0
}

// Returns true if string is empty
func IsEmpty(inputStr string) bool {
	return len(inputStr) == 0
}

// Deletes all whitespaces from a String
func DeleteWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsSpace(r) {
			return r
		}
		return -1
	}, str)
}

// Return true if both strings are case insensitive equal
func CaseInsensitiveEquals(firstString, secondString string) bool {
	return strings.ToLower(firstString) == strings.ToLower(secondString)
}

// Determines whether the string is a boolean representation
func IsBoolean(inputStr string) bool {
	if inputStr == `true` || inputStr == `false` {
		return true
	}

	return false
}

// Determines whether the string is a integer representation
func IsInteger(inputStr string) bool {
	if _, err := strconv.Atoi(inputStr); err == nil {
		return true
	}

	return false
}

// Determines whether the string is a float representation
func IsFloat(inputStr string) bool {
	if _, err := strconv.ParseFloat(inputStr, 64); err == nil {
		return true
	}

	return false
}

// Returns index of the specified character or substring in a particular String
func IndexOf(inputStr string, substring string) int {
	if IsEmpty(inputStr) || IsEmpty(substring) {
		return -1
	}

	return strings.Index(inputStr, substring)
}

// Removes the accents from a string, converting them to their non-accented corresponding characters.
func RemoveAccents(s string) string {
	byteArray := make([]byte, len(s))

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	dst, _, e := t.Transform(byteArray, []byte(s), true)
	if e != nil {
		panic(e)
	}

	return string(byteArray[:dst])
}

// isMn compares if Unicode characters are in category Mn (Mark, nonspacing)
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
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
