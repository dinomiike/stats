package cmd

import (
	"math"
	"sort"
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

var medianCmd = &cobra.Command{
	Use:   "median",
	Short: "The median or middle number of a sorted set of input",
	Long: `The median is the middle number in a set of numbers. It is calculated by sorting the numbers and finding the middle number.
	If there is an even number of numbers, the median is the average of the two middle numbers.
	
	For example, the median of 1, 2, 3, 4, 5 is 3. The median of 1, 2, 3, 4 is (2 + 3) / 2 = 2.5.
	
	To use this command, provide a set of numbers as arguments. For example:
	$ stats median 1 2 3 4 5`,
	Run: median,
}

var modeCmd = &cobra.Command{
	Use:   "mode",
	Short: "The mode is the most common number in the set",
	Long: `The mode is the most common number in the set of numbers. If there are no repeat numbers, the result is "no mode."
	
	For example, the mode of 1, 2, 3, 2, 5, 6, 2, 8 is 2.
	The mode of 1, 2, 3, 4, 5 is no mode.
	
	To use this command, provide a set of numbers as arguments. For example:
	$ stats mode 1 2 3 4 5 3`,
	Run: mode,
}

var (
	successCopy = color.New(color.FgGreen, color.Bold).SprintFunc()
	errorCopy   = color.New(color.FgRed, color.Bold).SprintFunc()
	warnCopy    = color.New(color.FgYellow, color.Bold).SprintFunc()
	generalCopy = color.New(color.FgWhite).SprintFunc()
)

func init() {
	rootCmd.AddCommand(meanCmd)
	rootCmd.AddCommand(medianCmd)
	rootCmd.AddCommand(modeCmd)

	meanCmd.Flags().BoolP("verbose", "v", false, "Enable verbose mode")
	medianCmd.Flags().BoolP("verbose", "v", false, "Enable verbose mode")
	modeCmd.Flags().BoolP("verbose", "v", false, "Enable verbose mode")
}

// implementation of the mean command
func mean(cmd *cobra.Command, args []string) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		cmd.Printf("%s\n", errorCopy("Error parsing verbose flag."))
		return
	}

	if verbose {
		cmd.Printf("%s\n", warnCopy("Verbose mode is enabled. Calculating mean..."))
	}

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

	if verbose {
		cmd.Printf("%s %v\n", generalCopy("Calculating mean for numbers:"), numbers)
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	mean := sum / float64(len(numbers))

	if verbose {
		cmd.Printf("%s %f\n", generalCopy("Sum of numbers:"), sum)
		cmd.Printf("%s %d\n", generalCopy("Count of numbers:"), len(numbers))
		cmd.Printf("%s %f / %d\n", generalCopy("Mean is calculated as sum / count:"), sum, len(numbers))
	}

	cmd.Printf("%s %f (%.2f)\n", successCopy("Result:"), mean, math.Floor(mean*100)/100)
}

// implementation of the median command
func median(cmd *cobra.Command, args []string) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		cmd.Printf("%s\n", errorCopy("Error parsing verbose flag."))
		return
	}

	if verbose {
		cmd.Printf("%s\n", warnCopy("Verbose mode is enabled. Calculating median..."))
	}

	if len(args) == 0 {
		cmd.Help()
		cmd.Printf("%s\n", errorCopy("Please provide a set of numbers."))
		return
	}

	var numbers []float64
	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			cmd.Printf("%s %s\n", errorCopy("Invalid number:"), arg)
			return
		}
		numbers = append(numbers, num)
	}

	if verbose {
		cmd.Printf("%s %v\n", generalCopy("Calculating median for numbers:"), numbers)
	}

	// sort numbers in ascending order
	sort.Float64s(numbers)

	if verbose {
		cmd.Printf("%s %v\n", generalCopy("Sorted numbers:"), numbers)
	}

	var median float64
	if len(numbers)%2 == 0 {
		// middle is middle average of the 2 numbers in the middle
		firstMiddleIndex := (len(numbers) / 2) - 1
		secondMiddleIndex := len(numbers) / 2
		if verbose {
			cmd.Printf("%s %f + %f / 2\n", warnCopy("Because of an even number of values, the median is calculated as:"), numbers[firstMiddleIndex], numbers[secondMiddleIndex])
		}
		median = (numbers[firstMiddleIndex] + numbers[secondMiddleIndex]) / 2
	} else {
		if verbose {
			cmd.Printf("%s %f\n", generalCopy("Median is the middle number:"), numbers[len(numbers)/2])
		}
		median = numbers[len(numbers)/2]
	}

	cmd.Printf("%s %f (%.2f)\n", successCopy("Result:"), median, math.Floor(median*100)/100)
}

// implementation of the mode command
func mode(cmd *cobra.Command, args []string) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		cmd.Printf("%s\n", errorCopy("Error parsing verbose flag."))
		return
	}

	if verbose {
		cmd.Printf("%s\n", warnCopy("Verbose mode is enabled. Calculating median..."))
	}

	if len(args) == 0 {
		cmd.Help()
		cmd.Printf("%s\n", errorCopy("Please provide a set of numbers."))
		return
	}

	// store the frequency of each number
	frequency := make(map[string]int)

	// group numbers by frequency to test for bimodal and multimodal outcomes
	var frequencyGroup [][]string

	// count the occurrences of each number
	for _, arg := range args {
		if _, err := strconv.ParseFloat(arg, 64); err != nil {
			cmd.Printf("%s\n", errorCopy("Invalid number:", arg))
			return
		}
		// increment the frequency of this number by 1
		frequency[arg]++
		// get the frequency of this number
		currentFrequency := frequency[arg]
		// note that this implies a 0 occurrence can't happen
		// i prefer this because it makes counting the length of the frequencyGroup more accurate, without having 0 index as an issue

		// Ensure the slice at the current frequency index is initialized
		if currentFrequency >= len(frequencyGroup) {
			// Extend the slice to accommodate the current frequency index
			newGroup := make([][]string, currentFrequency+1)
			copy(newGroup, frequencyGroup)
			frequencyGroup = newGroup
		}
		frequencyGroup[currentFrequency] = append(frequencyGroup[currentFrequency], arg)
	}

	cmd.Printf("frequencyGroup: %v\n", frequencyGroup)

	// check for no mode
	// the 0 index represents numbers in the list that don't appear -- always empty
	// the 1 index represents numbers that appear only one time
	// therefore, if your length is 2, there are no repeated numbers
	if len(frequencyGroup) == 2 {
		cmd.Printf("%s\n", warnCopy("Result: no mode"))
		return
	}

	lastSet := frequencyGroup[len(frequencyGroup)-1]
	// check for bimodal by getting the number of scores in the last set of the frequencyGroup
	if len(lastSet) == 2 {
		cmd.Printf("%s %v\n", successCopy("Result: bimodal:"), lastSet)
		return
	}

	// check for multimodal in the same way as bimodal, except it's anything more than 2
	if len(lastSet) > 2 {
		cmd.Printf("%s %v\n", successCopy("Result: multimodal:"), lastSet)
		return
	}

	cmd.Printf("%s %s\n", successCopy("Result:"), lastSet[0])
}
