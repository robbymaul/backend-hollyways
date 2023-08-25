package utility

import (
	"errors"
	"html"
	"net/mail"
	"regexp"
)

func ValidateAndSanitazeEmail(inputEmail string) (string, error) {
	_, err := mail.ParseAddress(inputEmail)
	if err != nil {
		return "", err
	}

	sanitizedEmail := html.EscapeString(inputEmail)

	return sanitizedEmail, nil
}

func ValidateInput(input string) (string, error) {
	specialChars := "`~!@#$%^&*()_+-=[{|:;}]'',<.>/?"

	pattern := "[" + regexp.QuoteMeta(specialChars) + "]"

	re := regexp.MustCompile(pattern)
	if re.MatchString(input) {
		return "", errors.New("Input contains special characters")
	}

	return input, nil
}
