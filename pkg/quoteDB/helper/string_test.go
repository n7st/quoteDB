package helper

import (
	"fmt"
	"testing"
)

// TestOptionsFromString() runs the OptionsFromString() helper twice and
// validates the output.
func TestOptionsFromString(t *testing.T) {
	input1 := `A string with "one" and "two" individual word options in it`
	input2 := `A string with "one multiword option" and a "second multiword option"`
	input3 := `A string with "a first match", "a second match" and "a third match"`
	output1 := OptionsFromString(input1)
	output2 := OptionsFromString(input2)
	output3 := OptionsFromString(input3)

	t.Run("Options check", func(t *testing.T) {
		if len(output1) != 2 {
			fmt.Println("[1] There should only be two options")
			t.Fail()
		}

		if output1[0] != "one" {
			fmt.Println("[1] The first option is incorrect")
			t.Fail()
		}

		if output1[1] != "two" {
			fmt.Println("[1] The second option is incorrect")
			t.Fail()
		}

		if len(output2) != 2 {
			fmt.Println("[2] There should only be two options")
			t.Fail()
		}

		if output2[0] != "one multiword option" {
			fmt.Println("[2] The first option is incorrect")
			t.Fail()
		}

		if output2[1] != "second multiword option" {
			fmt.Println("[2] The second option is incorrect")
			t.Fail()
		}

		if len(output3) != 3 {
			fmt.Println("[3] There should be three options")
			t.Fail()
		}

		if output3[2] != "a third match" {
			fmt.Println("[3] The third option is incorrect")
			t.Fail()
		}
	})
}
