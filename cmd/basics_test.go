package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestMean(t *testing.T) {
	// create a buffer to capture the output
	output := &bytes.Buffer{}

	// create a mock command to simulate the mean command
	cmd := &cobra.Command{
		Use: "mean",
		Run: mean,
	}
	cmd.SetOut(output)

	cmd.Flags().BoolP("verbose", "v", false, "Disable verbose mode")

	// test cases:
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Happy path - valid input",
			args:     []string{"1", "2", "3", "4", "5"},
			expected: "Result: 3.00",
		},
		{
			name:     "Happy path - works with unsorted input",
			args:     []string{"5", "9", "3", "7", "2", "1", "15"},
			expected: "Result: 6.00",
		},
		{
			name:     "Happy path - works with decimals",
			args:     []string{"1.05", "2.15", "9.99", "3.6", "7.772"},
			expected: "Result: 4.91",
		},
		{
			name:     "No params given",
			args:     []string{},
			expected: "Please provide a set of numbers.",
		},
		{
			name:     "Unhappy path - invalid input",
			args:     []string{"1", "a", "3"},
			expected: "Invalid number: a",
		},
	}

	// execute mean tests:
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// reset the output buffer
			output.Reset()

			// set the arguments for the command
			cmd.SetArgs(tt.args)

			// execute the command
			cmd.Execute()

			// validate the output
			if !strings.Contains(output.String(), tt.expected) {
				t.Errorf("[error]: expected %q, but got %q", tt.expected, output.String())
			}
		})
	}
}
