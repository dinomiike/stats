package cmd

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var meanCmd = &cobra.Command{
	Use:   "mean",
	Short: "The average of a set of numbers",
	Long: `The mean is the average of a set of numbers. It is calculated by adding all the numbers together and dividing by the number of numbers.
	
	For example, the mean of 1, 2, 3, 4, 5 is (1 + 2 + 3 + 4 + 5) / 5 = 3.
	
	To use this command, provide a set of numbers as arguments. For example:
	$ stats mean 1 2 3 4 5`,
	Run: mean,
}

var (
	successCopy = color.New(color.FgGreen, color.Bold).SprintFunc()
	errorCopy   = color.New(color.FgRed, color.Bold).SprintFunc()
	warnCopy    = color.New(color.FgYellow).SprintFunc()
)

func init() {
	rootCmd.AddCommand(meanCmd)
}

func mean(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		cmd.Printf("%s\n", errorCopy("Please provide a set of numbers."))
		return
	}

	var numbers []float64
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			cmd.Printf("Invalid number: %s\n", arg)
			return
		}
		numbers = append(numbers, num)
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	mean := sum / float64(len(numbers))
	cmd.Printf("%s %f\n", successCopy("Result:"), mean)
}
