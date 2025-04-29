package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

type testCase struct {
	name     string
	args     []string
	expected string
}

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
	tests := []testCase{
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
			name:     "No params provided",
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
	testRunner(t, tests, output, cmd)
}

func TestMedian(t *testing.T) {
	output := &bytes.Buffer{}

	cmd := &cobra.Command{
		Use: "median",
		Run: median,
	}
	cmd.SetOut(output)

	cmd.Flags().BoolP("verbose", "v", false, "Disable verbose mode")

	// test cases:
	tests := []testCase{
		{
			name:     "Happy path - valid input",
			args:     []string{"1", "2", "3", "4", "5"},
			expected: "Result: 3",
		},
		{
			name:     "Happy path - unsorted, valid input and even number of parameters",
			args:     []string{"10", "3", "5", "9", "2", "6"},
			expected: "Result: 5.5",
		},
		{
			name:     "Happy path - works with decimals",
			args:     []string{"5.3", "2.62", "9.81", "14.01", "11.5"},
			expected: "Result: 9.81",
		},
		{
			name:     "No params provided",
			args:     []string{},
			expected: "Please provide a set of numbers.",
		},
		{
			name:     "Unhappy path - invalid input",
			args:     []string{"9", "z", "3", "y", "5", "x"},
			expected: "Invalid number: z",
		},
	}

	// execute the tests:
	testRunner(t, tests, output, cmd)
}

func testRunner(t *testing.T, tests []testCase, output *bytes.Buffer, cmd *cobra.Command) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// reset the output buffer
			output.Reset()

			// set the arguments for the command
			cmd.SetArgs(tt.args)

			// execute the command
			cmd.Execute()

			// validate the input
			if !strings.Contains(output.String(), tt.expected) {
				t.Errorf("[error]: expected %q, but got %q", tt.expected, output)
			}
		})
	}
}
