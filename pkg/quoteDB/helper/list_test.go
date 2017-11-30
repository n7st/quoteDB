package helper

import (
	"fmt"
	"testing"
	"time"

	"github.com/n7st/quoteDB/pkg/quoteDB/helper"
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
	options := []string{"Foo", "Baz"}

	t.Run("History check", func(t *testing.T) {
		lines := helper.LinesFromHistory(input, options)

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
