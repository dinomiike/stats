package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var meanCmd = &cobra.Command{
	Use:   "mean",
	Short: "The average of a set of numbers",
	Long: `The mean is the average of a set of numbers. It is calculated by
	adding all the numbers together and dividing by the number of numbers.
	
	For example, the mean of 1, 2, 3, 4, and 5 is (1 + 2 + 3 + 4 + 5) / 5 = 3.`,
	Run: mean,
}

func init() {
	rootCmd.AddCommand(meanCmd)
}

func mean(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		cmd.Println("Please provide a set of numbers.")
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
	cmd.Printf("Result: %f\n", mean)
}
