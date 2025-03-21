package parser

import (
	"fmt"
	"strings"

	"github.com/vphatfla/lets-db/db/io"
)

// Function to parse the raw string query, return the appropriate handler, key, value, and error
func ParseQuery(query string) (io.IOHandler, string, string,  error) {
	query = strings.TrimSpace(query)

	arr := strings.Split(query, " ")

	if len(arr) == 0 {
		return nil, "", "", fmt.Errorf("Empty query")
	}

	switch	cmd := strings.ToLower(arr[0]); cmd {
	case "select":
		if key, err := parseSelect(arr); err != nil {
			return nil, "", "", err
		} else {
			var h io.IOHandler = &io.IOSelectHandler{}
			return h, key, "", nil
		}
	case "delete":
		if key, err := parseDelete(arr); err != nil {
			return nil, "", "", err
		} else {
			var h io.IOHandler = &io.IODeleteHandler{}
			return h, key, "", nil
		}
	case "insert":
		if key, value, err := parseInsert(arr); err != nil {
			return nil, "", "", err
		} else {
			var h io.IOHandler = &io.IOInsertHandler{}
			return h, key, value, nil
		}
	default:
		return nil, "", "", fmt.Errorf("No existing command, command given in query %s ", query)
	}
}

// Parse the select query, syntax is select xxx where xxx is the key value
// return the key and error if exist
func parseSelect(arr []string) (string, error) {
	if len(arr) != 2 {
		return "", fmt.Errorf("Incorrect format for select query")
	}
	return arr[1], nil
}

// Parse the delete query, syntax is "delete xxx" where xxx is the key value
// return the key and error if exist
func parseDelete(arr []string) (string, error) {
	if len(arr) != 2 {
		return "", fmt.Errorf("Incorrect format for delete query")
	}
	return arr[1], nil
}

// Parse the insert query, syntax is "insert xxx value yyy" where xxx is the key value
// return the key, value  and error if exist
func parseInsert(arr []string) (string, string, error) {
	if len(arr) != 4 {
		return "", "", fmt.Errorf("Incorrect format for select query")
	}

	if temp := strings.ToLower(arr[2]); temp != "value" {
		return "", "", fmt.Errorf("Incorrect format for write query, must be INSERT <key> value <value")
	}

	return arr[1], arr[3], nil
}
