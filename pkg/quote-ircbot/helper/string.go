// The helper package provides functions for common tasks.
package helper

import "regexp"

var quoteRegex = regexp.MustCompile("([^\"]*)")

// OptionsFromString() transforms a string with quoted sections into a list of
// the contents of the quoted regions.
// input:  "Hello" "World"
// output: [ "Hello", "World" ]
func OptionsFromString(input string) (output []string) {
	matches := quoteRegex.FindAllStringSubmatch(input, -1)

	for i, v := range matches {
		// Skip non-options (i.e. strings between options)
		if i % 2 == 0 {
			continue
		}

		output = append(output, v[0])
	}

	return
}
