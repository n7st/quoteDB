package helper

import (
	"fmt"
	"testing"
	"time"
)

// Example mock input structured the same as in the bot's memory
var input = []map[string]string{
	{
		"nick":      "User 1",
		"timestamp": time.Now().String(),
		"message":   "First message",
	},
	{
		"nick":      "User 2",
		"timestamp": time.Now().String(),
		"message":   "Foo Test",
	},
	{
		"nick":      "User 1",
		"timestamp": time.Now().String(),
		"message":   "Bar Test",
	},
	{
		"nick":      "User 3",
		"timestamp": time.Now().String(),
		"message":   "Baz Test",
	},
	{
		"nick":      "User 3",
		"timestamp": time.Now().String(),
		"message":   "Last Message",
	},
}

// TestLinesFromHistory() runs the LinesFromHistory() helper with the example
// input and ensures its output is as expected
func TestLinesFromHistory(t *testing.T) {
	options := []string{"Foo", "baz"} // lower case should work too

	t.Run("History check", func(t *testing.T) {
		lines := LinesFromHistory(input, options)

		// input[1..3] have been requested from history
		if len(lines) != 3 {
			fmt.Println("Incorrect number of lines returned from search")
			t.Fail()
		}

		if lines[0]["message"] != "Foo Test" {
			fmt.Println("First message not as expected")
			t.Fail()
		}

		if lines[1]["message"] != "Bar Test" {
			fmt.Println("Second message not as expected")
			t.Fail()
		}

		if lines[2]["message"] != "Baz Test" {
			fmt.Println("Second message not as expected")
			t.Fail()
		}
	})
}

// TestLastNLinesFromHistory() runs two tests to check the history returned is
// as expected.
func TestLastNLinesFromHistory(t *testing.T) {
	t.Run("Last 2 lines from history", func(t *testing.T) {
		lenParam := 2
		lines := LastNLinesFromHistory(input, lenParam)

		if len(lines) != lenParam {
			fmt.Printf("Expected 2 lines, got %d\n", len(lines))
			t.Fail()
		}

		if lines[0]["message"] != "Baz Test" {
			fmt.Printf("The first message is incorrect (%s)\n",
				lines[0]["message"])
			t.Fail()
		}

		if lines[1]["message"] != "Last Message" {
			fmt.Printf("The last message is incorrect (%s)\n",
				lines[1]["message"])
			t.Fail()
		}
	})

	t.Run("Last 3 lines from history", func(t *testing.T) {
		lenParam := 3
		lines := LastNLinesFromHistory(input, lenParam)

		if len(lines) != lenParam {
			fmt.Printf("Expected 3 lines, got %d\n", len(lines))
			t.Fail()
		}
	})

	// The number passed to LastNLinesFromHistory() is larger than the length
	// of the input. Check the array hasn't overflowed.
	t.Run("Check for overflow", func(t *testing.T) {
		lines := LastNLinesFromHistory(input, 7)

		if len(lines) != len(input) {
			fmt.Println("Expected five lines")
			t.Fail()
		}
	})
}
