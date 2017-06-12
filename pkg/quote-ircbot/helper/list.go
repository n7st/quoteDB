// The helper package provides functions for common tasks.
package helper

import "strings"

// LinesFromHistory() searches a history list for a start and end point and
// grabs all the lines in between.
func LinesFromHistory(input []map[string]string, options []string) (output []map[string]string) {
	matched := false

	for _, line := range input {
		if strings.HasPrefix(line["message"], options[0]) {
			matched = true
		}

		// Keep appending lines until we don't match anymore
		if matched {
			output = append(output, line)
		}

		if strings.HasPrefix(line["message"], options[1]) {
			matched = false
			break
		}
	}

	return
}
