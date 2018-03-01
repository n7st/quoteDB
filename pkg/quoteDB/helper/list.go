// The helper package provides functions for common tasks.
package helper

import "strings"

// LinesFromHistory() searches a history list for a start and end point and
// grabs all the lines in between.
func LinesFromHistory(input []map[string]string, options []string) (output []map[string]string) {
	matched := false

	for _, line := range input {
		message := strings.ToLower(line["message"])
		firstIdx := strings.ToLower(options[0])
		secondIdx := strings.ToLower(options[1])

		if strings.HasPrefix(message, firstIdx) {
			matched = true
		}

		// Keep appending lines until we don't match anymore
		if matched {
			output = append(output, line)
		}

		if strings.HasPrefix(message, secondIdx) {
			matched = false
			break
		}
	}

	return
}

// LastNLinesFromHistory() gets the last n lines from the input array. If the
// length of the input array is shorter than the requested length, it is set
// to be the length of the array.
func LastNLinesFromHistory(input []map[string]string, n int) (output []map[string]string) {
	inpLen := len(input)
	end := inpLen - n

	if n >= inpLen {
		end = 0
	}

	for i := inpLen - 1; i >= end; i-- {
		output = append(output, input[i])
	}

	for i := len(output)/2 - 1; i >= 0; i-- {
		opp := len(output) - 1 - i

		output[i], output[opp] = output[opp], output[i]
	}

	return
}
