package adapters

import "strings"

func isNotFound(err error) bool {
	errorMessage := getMessage("not_found")
	return strings.Contains(err.Error(), errorMessage)
}

func getMessage(errorType string) string {
	messages := getErrors()
	return messages[errorType]
}

func getErrors() map[string]string {
	return map[string]string{
		"not_found": "gogm: data not found",
	}
}
